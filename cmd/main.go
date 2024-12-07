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

//	@title		Go Social API
//	@version	1.0

// @license.name	MIT
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

	log.Println("Creating middleware stack")
	middlewareStack := middleware.CreateStack(
		middleware.AuthMiddleware,
	)

	router := http.NewServeMux()

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
		Handler: middlewareStack(m.Mux(), tokenProvider),
	}

	log.Println("Starting server")
	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
