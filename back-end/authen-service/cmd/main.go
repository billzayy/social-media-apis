package main

import (
	"fmt"
	"log"
	"net"
	"time"

	auth "github.com/billzayy/social-media/back-end/authen-service/api"
	"github.com/billzayy/social-media/back-end/authen-service/internal/db"
	"github.com/billzayy/social-media/back-end/authen-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/authen-service/internal/handlers"
	"github.com/billzayy/social-media/back-end/authen-service/internal/routes"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	postgres, err := db.ConnectDB()
	defer postgres.Close()

	if err != nil {
		fmt.Println(err)
	}

	redis, err := db.ConnectRedis()

	defer redis.Close()

	if err != nil {
		fmt.Println(err)
	}

	h := handlers.NewHandlers(repositories.NewRepositories(postgres, redis).AuthRepository)

	go func() {
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

		routes.SetupRoutes(r, h)

		log.Println("REST API server started on :3000")
		r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}()

	go func() {
		grpcServer := grpc.NewServer()
		auth.RegisterAuthServiceServer(grpcServer, handlers.NewAuthGrpcServer(repositories.NewAuthRepository(postgres, redis)))

		lis, err := net.Listen("tcp", ":50051")

		if err != nil {
			log.Fatalf("failed to listen : %v", err)
		}

		log.Println("gRPC server started on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to server : %v", err)
		}
	}()

	// Block forever
	select {}
}
