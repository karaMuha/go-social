package config

import (
	"errors"
	"os"
)

type Config struct {
	RestPort           string
	DbConnectionString string
	DbDriver           string
	PrivateKeyPath     string
}

func InitConfig() (*Config, error) {
	restPort, ok := os.LookupEnv("REST_PORT")
	if !ok {
		return nil, errors.New("could not load rest port environment variable")
	}

	dbConnectionString, ok := os.LookupEnv("DB_CONNECTION")
	if !ok {
		return nil, errors.New("could not load db connection environment variable")
	}

	dbDriver, ok := os.LookupEnv("DB_DRIVER")
	if !ok {
		return nil, errors.New("could not load db driver environment variable")
	}

	privateKeyPath, ok := os.LookupEnv("PRIVATE_KEY_PATH")
	if !ok {
		return nil, errors.New("could not load private key path environment variable")
	}

	return &Config{
		RestPort:           restPort,
		DbConnectionString: dbConnectionString,
		DbDriver:           dbDriver,
		PrivateKeyPath:     privateKeyPath,
	}, nil
}
