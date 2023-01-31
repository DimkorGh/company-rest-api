package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Environment struct {
}

func NewEnvironment() *Environment {
	return &Environment{}
}

func (environment *Environment) Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while initializing .env file: %s", err.Error())
	}
}

func (environment *Environment) InitializeTestEnv() {
	path, _ := os.Getwd()

	for {
		err := godotenv.Load(string(path) + `/.env`)
		if err != nil {
			lastSlash := strings.LastIndex(path, "/")
			if lastSlash == -1 {
				log.Fatal()
				break
			}
			path = path[:lastSlash]

			continue
		}

		break
	}
}
