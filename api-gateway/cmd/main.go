package main

import (
	"log"
	"os"
	"time"

	"github.com/billzayy/social-media/api-gateway/internal/handlers"
	"github.com/billzayy/social-media/api-gateway/internal/routes"
	"github.com/billzayy/social-media/api-gateway/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/billzayy/social-media/api-gateway/docs"
)

// @title						Swagger Social Media CRUD API
// @version					1.0
// @description				This is an API Gateway Server.
// @BasePath					/
// @schemes					http https
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load("./internal/.env")

	if err != nil {
		if err := godotenv.Load("../internal/.env"); err != nil {
			log.Fatalf("Error loading .env on Postgres")
		}
	}

	docs.SwaggerInfo.Host = "localhost:" + os.Getenv("REST_PORT")

	r := gin.New()

	h := handlers.NewHandlers(&handlers.PortList{
		AuthPort:         os.Getenv("AUTH_GRPC"),
		PostPort:         os.Getenv("POST_GRPC"),
		UserPort:         os.Getenv("USER_GRPC"),
		NotificationPort: os.Getenv("NOTIFICATION_GRPC"),
		ChatPort:         os.Getenv("CHAT_GRPC"),
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173","https://poised-janice-billzay-21089a93.koyeb.app"},                   // Allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Content-Type", "Authorization"},           // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                          // Exposed headers
		AllowCredentials: true,                                                // Allow credentials (cookies)
		MaxAge:           12 * time.Hour,                                      // Cache duration for preflight requests
	}))

	r.Use(gin.Logger())
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})

	socketServer := utils.NewWebSocketServer()

	go socketServer.Run()

	routes.SetupRoutes(r, h, socketServer)

	log.Printf("REST API server started on :%v\n", os.Getenv("REST_PORT"))
	r.Run(":" + os.Getenv("REST_PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
