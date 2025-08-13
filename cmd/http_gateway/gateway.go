package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/SudarshanHV/perl2go_grpc/go_server"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	err := pb.RegisterPrimeServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", []grpc.DialOption{grpc.WithInsecure()})

	if err != nil {
		log.Fatalf("Failed to start gateway: %v", err)
	}

	log.Println("HTTP Gateway is listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to serve :%v", err)
	}

}
