package main

import (
	"log"
	"net"
	"os"

	pb "github.com/Morning1139Angel/web-hw1/auth/grpc"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

var redisHost = os.Getenv("REDIS_HOST")
var redisPort = os.Getenv("REDIS_PORT")
var redisPassword = os.Getenv("REDIS_PASSWORD")

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()

	//initialize redis client
	rdb := initRedicClient()

	// Register implementation of the service
	authService := &authServer{rdb: rdb}
	pb.RegisterAuthServiceServer(server, authService)

	// Create a TCP listener
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start serving requests
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func initRedicClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0, // use default DB
	})
}