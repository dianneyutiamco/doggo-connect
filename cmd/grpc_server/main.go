package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/dianneyutiamco/doggo-connect/internal/api/dog_profile"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "Server port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening to %v.\n", lis.Addr())
	s := grpc.NewServer()
	dogProfileService := &dog_profile.DogProfileServiceImpl{}
	dog_profile.RegisterDogProfileServiceServer(s, dogProfileService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Println("Hello Puppies!")
}
