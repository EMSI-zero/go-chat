package colrepo

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type ColConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

var Username = "COL_DB_USER_NAME"
var Password = "COL_DB_PASS"
var Host = "COL_DB_HOST"
var Port = "COL_DB_PORT"
var Database = "COL_DB_DATABASE"

type ColDBConn interface {
	GetConn() (*mongo.Client, error)
}

func (cfg *ColConfig) buildConfigFromEnv() error {
	cfg.Username = os.Getenv(Username)
	if cfg.Username == "" {
		return fmt.Errorf("%v env must be specified: username", Username)
	}

	cfg.Password = os.Getenv(Password)
	if cfg.Password == "" {
		return fmt.Errorf("%v env must be specified: password", Password)
	}

	cfg.Host = os.Getenv(Host)
	if cfg.Host == "" {
		return fmt.Errorf("%v env must be specified: host", Host)
	}

	cfg.Port = os.Getenv(Port)
	if cfg.Port == "" {
		return fmt.Errorf("%v env must be specified: port", Port)
	}

	cfg.Database = os.Getenv(Database)
	if cfg.Database == "" {
		return fmt.Errorf("%v env must be specified: database", Database)
	}

	return nil
}

func NewColDBConn() (ColDBConn, error) {
	return NewMongoConn()
}
