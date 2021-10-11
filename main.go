package main

import (
	"fmt"
	"go-cart-api/route"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}
}

func main() {
	fmt.Println("Starting main application")

	loadEnv()
	serve := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Fatal(route.RunAPI(serve))
}
