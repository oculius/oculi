package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	fmt.Println("Starting server...")
}
