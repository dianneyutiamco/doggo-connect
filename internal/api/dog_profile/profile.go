package dog_profile

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

type DogProfileServiceImpl struct {
	DogProfileServiceServer
}

// Implementations for Dog Profile Service
func CreateDogProfile() {
	log.Printf("This creates a new dog profile.")
}

func (s *DogProfileServiceImpl) CreateDogProfile(ctx context.Context, req *DogProfile) (*emptypb.Empty, error) {
	return nil, nil
}
