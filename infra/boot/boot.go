package boot

import (
	"log"
	"os"

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
	}

	return err
}

func parseFlags() error {
	return nil
}
