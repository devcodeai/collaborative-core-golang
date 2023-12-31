package main

import (
	"fmt"
	"log"

	"github.com/devcodeai/collaborative-core-golang/Database"
	"github.com/devcodeai/collaborative-core-golang/Routes"
	"github.com/devcodeai/collaborative-core-golang/Utils"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}

	MYSQL_USER := Utils.EnvConfig("MYSQL_USER", "root")
	MYSQL_PASSWORD := Utils.EnvConfig("MYSQL_PASSWORD", "password")
	MYSQL_HOST := Utils.EnvConfig("MYSQL_HOST", "localhost")
	MYSQL_PORT := Utils.EnvConfig("MYSQL_PORT", "3306")
	MYSQL_DBNAME := Utils.EnvConfig("MYSQL_DBNAME", "devcode")
	DB_URL := Database.GetDatabaseURL(MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DBNAME)

	Database.ConnectDatabase(DB_URL)

	HOST := Utils.EnvConfig("HOST", "0.0.0.0")
	PORT := Utils.EnvConfig("PORT", "3030")
	SERVER_URL := fmt.Sprintf("%v:%v", HOST, PORT)

	gin.SetMode(gin.ReleaseMode)
	router := Routes.SetupRouter()
	fmt.Printf("[SERVER] Server to be run on http://%v/api \n", SERVER_URL)
	serverErr := router.Run(SERVER_URL)
	if serverErr != nil {
		fmt.Printf("[SERVER] Failed to start the server!\n")
		fmt.Printf("[SERVER] Error: %v\n", serverErr)
	}
}
