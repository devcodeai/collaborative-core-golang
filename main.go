package main

import (
	"log"
	"fmt"

	"github.com/devcodeai/collaborative-core-golang/Routes"
	"github.com/devcodeai/collaborative-core-golang/Utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	HOST := Utils.EnvConfig("HOST", "0.0.0.0")
	PORT := Utils.EnvConfig("PORT", "3030")
	SERVER_URL := fmt.Sprintf("%v:%v", HOST, PORT)

	router := Routes.SetupRouter()
	router.Run(SERVER_URL)
}
