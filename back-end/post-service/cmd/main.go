package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	post "github.com/billzayy/social-media/back-end/post-service/api"
	"github.com/billzayy/social-media/back-end/post-service/internal/db"
	"github.com/billzayy/social-media/back-end/post-service/internal/db/repositories"
	"github.com/billzayy/social-media/back-end/post-service/internal/handlers"
	"github.com/billzayy/social-media/back-end/post-service/internal/routes"
	"github.com/billzayy/social-media/back-end/post-service/internal/services"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load("./internal/.env")

	if err != nil {
		if err := godotenv.Load("../internal/.env"); err != nil {
			log.Fatalf("Error loading .env")
		}
	}

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

	h := handlers.NewHandlers(services.NewServices(
		repositories.NewRepositories(postgres, redis).PostRepository,
		repositories.NewRepositories(postgres, redis).InteractRepository,
	))

	add := flag.String("mode", "", "Post Service Mode")

	flag.Parse()

	switch {
	case *add == "test":
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

		log.Printf("REST API server started on :%v\n", os.Getenv("REST_PORT"))
		r.Run(":" + os.Getenv("REST_PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	case *add == "deploy":
		grpcServer := grpc.NewServer()

		post.RegisterPostServiceServer(grpcServer, handlers.NewPostGrpcServer(
			services.NewServices(
				repositories.NewPostRepository(postgres, redis),
				repositories.NewInteractRepository(postgres, redis),
			),
		))

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
