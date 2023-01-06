package main

import (
	"log"
	"os"

	"github.com/NonNtp/gin-gorm/db"
	"github.com/NonNtp/gin-gorm/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

func main() {
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db.ConnectDB()
	db.Migrate()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r := gin.Default()
	r.Use(cors.New(corsConfig))
	routes.ServeRoutes(r)
	
	r.Run(":" + os.Getenv("PORT"))
}
