package main

import (
	"github.com/Lokeshxs/url-shortener/db"
	"github.com/Lokeshxs/url-shortener/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Loading the ENV variabels

	err := godotenv.Load()

	if err != nil {
		panic("Could not load the ENV variables!")
	}

	// Intializing the server
	server := gin.Default()

	// Connecting to Postgres DB
	db.InitDB()

	// Calling a function to handle incoming requests
	routes.RoutingHandler(server)

	// Start Server
	server.Run(":8080")

}
