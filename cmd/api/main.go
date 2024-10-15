package main

import (
	fiberconfig "github.com/emersonnobre/tica-api-go/internal/delivery/fiber"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	setup := fiberconfig.NewFiberSetup()
	setup.Execute()
}
