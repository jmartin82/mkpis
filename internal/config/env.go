package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type configuration struct {
	DevelopBranch string `env:"DEVELOP_BRANCH_NAME" envDefault:"devel"`
	MasterBranch  string `env:"MASTER_BRANCH_NAME" envDefault:"master"`
	GitHubToken   string `env:"GITHUB_TOKEN"`
}

func loadConfig() *configuration {
	godotenv.Load() //load .env

	cfg := &configuration{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	return cfg
}

var Env = loadConfig()
