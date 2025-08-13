package main

import (
	"context"
	"log"
	"net"

	pb "github.com/SudarshanHV/perl2go_grpc/go_server"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPrimeServiceServer
}

func (s *server) GetPrimes(ctx context.Context, req *pb.PrimeRequest) (*pb.PrimeResponse, error) {
	limit := req.Limit
	var primes []int32
	for num := int32(2); num <= limit; num++ {
		isPrime := true
		for div := int32(2); div*div <= num; div++ {
			if num%div == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, num)
		}
	}
	return &pb.PrimeResponse{Primes: primes}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen, error with server: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPrimeServiceServer(grpcServer, &server{})
	log.Println("gRPC Server up and running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
