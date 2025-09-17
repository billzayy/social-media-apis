package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	chat "github.com/billzayy/social-media/back-end/chat-service/api"
	"github.com/billzayy/social-media/back-end/chat-service/internal/db"
	"github.com/billzayy/social-media/back-end/chat-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/chat-service/internal/handlers"
	"github.com/billzayy/social-media/back-end/chat-service/internal/routes"
	"github.com/billzayy/social-media/back-end/chat-service/internal/services"
	"github.com/billzayy/social-media/back-end/chat-service/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("./internal/.env")

	if err != nil {
		if err := godotenv.Load("../internal/.env"); err != nil {
			log.Fatal(err)
			log.Fatalf("Error loading .env")
		}
	}

	postgres, err := db.ConnectDB()

	if err != nil {
		fmt.Printf("Postgres: %v\n", err)
		return
	}
	defer postgres.Close()

	redis, err := db.ConnectRedis()

	if err != nil {
		fmt.Printf("Redis: %v\n", err)
		return
	}
	defer redis.Close()

	rabbitConn, rabbitCh, err := utils.RabbitMQ()

	if err != nil {
		fmt.Printf("RabbitMQ: %v\n", err)
		// return
	}

	defer rabbitConn.Close()

	socketServer := utils.NewWebSocketServer()

	go socketServer.Run()

	h := handlers.NewHandlers(services.NewServices(
		repositories.NewRepositories(postgres, redis).NotifyRepository,
		&services.RabbitService{
			Channel: rabbitCh,
		},
		socketServer,
	))

	add := flag.String("mode", "", "Chat Service Mode")

	flag.Parse()
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Content-Type", "Authorization"},           // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                          // Exposed headers
		AllowCredentials: true,                                                // Allow credentials (cookies)
		MaxAge:           12 * time.Hour,                                      // Cache duration for preflight requests
	}))

	r.Use(gin.Logger())
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})

	switch {
	case *add == "test":

		routes.SetupRoutes(r, h, socketServer)

		log.Printf("REST API server started on :%v\n", os.Getenv("REST_PORT"))

		r.Run(":" + os.Getenv("REST_PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	case *add == "deploy":
		grpcServer := grpc.NewServer()

		chat.RegisterChatServiceServer(grpcServer, handlers.NewChatGrpcServer(
			services.NewServices(
				repositories.NewChatRepository(postgres, redis),
				&services.RabbitService{
					Channel: rabbitCh,
				},
				socketServer,
			),
		))

		r.GET("/ws/chat", socketServer.HandleWebSocket)
		go r.Run(":" + os.Getenv("WS_PORT"))

		lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))

		if err != nil {
			log.Fatalf("failed to listen : %v", err)
		}

		log.Printf("gRPC server started on :%v\n", os.Getenv("GRPC_PORT"))
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to server : %v", err)
		}
	default:
		fmt.Println("Error command")
		os.Exit(1)
	}
}
