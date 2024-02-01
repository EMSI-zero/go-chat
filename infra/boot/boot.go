package boot

import (
	"log"
	"os"

	"github.com/EMSI-zero/go-chat/gateway"
	"github.com/EMSI-zero/go-chat/infra/colrepo"
	"github.com/EMSI-zero/go-chat/infra/dbrepo"
	"github.com/EMSI-zero/go-chat/infra/httputils"
	"github.com/EMSI-zero/go-chat/registry"
	"github.com/EMSI-zero/go-chat/services"
	"github.com/joho/godotenv"
)

var bootHasFinished bool

func IsAfterBoot() bool {
	return bootHasFinished
}

func Boot() (err error) {
	if err = parseFlags(); err != nil {
		return err
	}

	if err = LoadEnv(); err != nil {
		return err
	}

	log.Print("connecting to database")
	db, err := dbrepo.NewDBConn()
	if err != nil {
		return err
	}
	log.Print("database connected")

	log.Print("connecting to collection database")
	colDb, err := colrepo.NewColDBConn()
	if err != nil {
		return err
	}
	log.Print("collection database connected")

	//Create a new service registry
	sr := registry.NewServiceRegistry(db, colDb)

	services.InitServices(sr)

	httpServer, err := gateway.CreateHTTPServer(sr)

	if err := httputils.StartServer(httpServer); err != nil {
		return err
	}

	if err == nil {
		bootHasFinished = true
	}
	return nil
}

func LoadEnv() (err error) {
	env := os.Getenv("ENVIRONMENT")
	if env != "" {
		return
	}

	//check for test environment
	if env != "test" {
		log.Print("loading local environments")
		err = godotenv.Load(".env.local")
		if err != nil {
			return err
		}
	}

	return nil
}

func parseFlags() error {
	return nil
}
