package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	posts "github.com/karaMuha/go-social/contents"
	authtoken "github.com/karaMuha/go-social/internal/auth_token"
	"github.com/karaMuha/go-social/internal/config"
	"github.com/karaMuha/go-social/internal/database/postgres"
	"github.com/karaMuha/go-social/internal/mailer"
	"github.com/karaMuha/go-social/internal/middleware"
	"github.com/karaMuha/go-social/internal/monolith"
	"github.com/karaMuha/go-social/users"

	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Println("Reading config")
	config, err := config.InitConfig()
	if err != nil {
		return fmt.Errorf("error while reading config: %v", err)
	}

	log.Println("Connecting to database")
	db, err := postgres.ConnectToDb(config.DbDriver, config.DbConnectionString)
	if err != nil {
		return fmt.Errorf("error while connecting to database: %v", err)
	}
	defer db.Close()

	log.Println("Reading private key")
	tokenProvider := authtoken.NewTokenProvider(config.PrivateKeyPath)

	mux := http.NewServeMux()
	registerHealthCheck(mux)
	rateLimiter := middleware.RateLimiter(mux)
	authorizer := middleware.Authorizer(rateLimiter, tokenProvider)

	modules := []monolith.IModule{
		&users.Module{},
		&posts.Module{},
	}

	log.Println("Connecting to mail server")
	mailServer := mailer.NewMailServer()

	m := monolith.NewMonolith(config, db, mux, mailServer, modules, tokenProvider)

	log.Println("Initializing modules")
	err = m.InitModules()
	if err != nil {
		return fmt.Errorf("error while initializing modules: %v", err)
	}

	server := &http.Server{
		Addr:        m.Config().RestPort,
		Handler:     authorizer,
		ReadTimeout: 5 * time.Second,
	}

	go startServer(server)

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	<-shutdown.Done()

	log.Println("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = server.Shutdown(ctx); err != nil {
		log.Printf("Shutdown with error: %v", err)
	}
	log.Println("Shutdown complete")

	return nil
}

func startServer(srv *http.Server) {
	log.Println("Start listening")
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Stopped listening: %v\n", err)
	}
}

func registerHealthCheck(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
