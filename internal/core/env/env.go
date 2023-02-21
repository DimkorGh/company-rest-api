package env

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Environment struct{}

func NewEnvironment() *Environment {
	return &Environment{}
}

func (env *Environment) Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while initializing .env file: %s", err.Error())
	}
}

func (env *Environment) InitializeTestEnv() {
	path, _ := os.Getwd()

	for {
		err := godotenv.Load(string(path) + `/.env`)
		if err != nil {
			lastSlash := strings.LastIndex(path, "/")
			if lastSlash == -1 {
				log.Fatal()
			}
			path = path[:lastSlash]

			continue
		}

		break
	}
}
