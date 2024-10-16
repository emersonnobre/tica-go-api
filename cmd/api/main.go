package main

import (
	"os"

	fiberconfig "github.com/emersonnobre/tica-api-go/internal/delivery/fiber"
	"github.com/joho/godotenv"
)

func main() {
	var env string
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	envFile := pickEnvironmentFile(env)

	godotenv.Load(envFile)
	setup := fiberconfig.NewFiberSetup()
	setup.Execute()
}

func pickEnvironmentFile(env string) string {
	switch env {
	case "development":
		return ".env.development"
	case "production":
		return ".env.production"
	default:
		return ".env.production"
	}
}
