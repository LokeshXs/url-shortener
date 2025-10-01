package main

import (
	"time"

	"github.com/Lokeshxs/url-shortener/db"
	"github.com/Lokeshxs/url-shortener/routes"
	"github.com/gin-contrib/cors"
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

	// Configuring CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // your Next.js app
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Calling a function to handle incoming requests
	routes.RoutingHandler(server)

	// Start Server
	server.Run(":8080")

}
