package main

import (
	"log"
	"net/http"

	authtoken "github.com/karaMuha/go-social/internal/auth_token"
	"github.com/karaMuha/go-social/internal/config"
	"github.com/karaMuha/go-social/internal/database/postgres"
	"github.com/karaMuha/go-social/internal/mailer"
	"github.com/karaMuha/go-social/internal/middleware"
	"github.com/karaMuha/go-social/internal/monolith"
	"github.com/karaMuha/go-social/posts"
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
		return err
	}

	log.Println("Connecting to database")
	db, err := postgres.ConnectToDb(config.DbDriver, config.DbConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("Reading private key")
	tokenProvider := authtoken.NewTokenProvider(config.PrivateKeyPath)

	router := http.NewServeMux()
	registerHealthCheck(router)
	rateLimiter := middleware.RateLimiter(router)
	authorizer := middleware.Authorizer(rateLimiter, tokenProvider)

	modules := []monolith.Module{
		&users.Module{},
		&posts.Module{},
	}

	log.Println("Starting mail server")
	mailServer := mailer.NewMailServer()

	m := monolith.NewMonolith(*config, db, router, mailServer, modules, tokenProvider)

	log.Println("Initializing modules")
	err = m.InitModules()
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    m.Config().RestPort,
		Handler: authorizer,
	}

	log.Println("Starting server")
	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func registerHealthCheck(router *http.ServeMux) {
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
