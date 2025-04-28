package repository

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectService(host string, port string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(host+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to service: %v", err)
		return conn, err
	}

	return conn, nil
}
