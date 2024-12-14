package main

import (
	"fmt"
	"log"
	"net/http"

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

	log.Println("Starting mail server")
	mailServer := mailer.NewMailServer()

	m := monolith.NewMonolith(config, db, mux, mailServer, modules, tokenProvider)

	log.Println("Initializing modules")
	err = m.InitModules()
	if err != nil {
		return fmt.Errorf("error while initializing modules: %v", err)
	}

	server := &http.Server{
		Addr:    m.Config().RestPort,
		Handler: authorizer,
	}

	log.Println("Starting server")
	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error while starting server: %v", err)
	}

	return nil
}

func registerHealthCheck(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
