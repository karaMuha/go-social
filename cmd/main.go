package main

import (
	"log"
	"net/http"

	"github.com/karaMuha/go-social/internal/config"
	"github.com/karaMuha/go-social/internal/database/postgres"
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

	router := http.NewServeMux()

	modules := []monolith.Module{
		&users.Module{},
		&posts.Module{},
	}

	m := monolith.NewMonolith(*config, db, router, modules)

	log.Println("Initializing modules")
	err = m.InitModules()
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    m.Config().RestPort,
		Handler: m.Mux(),
	}

	log.Println("Starting server")
	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
