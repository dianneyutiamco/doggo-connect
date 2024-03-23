package dog_profile

import (
	"context"
	"testing"
)

func TestGetDogProfile(t *testing.T) {
	// Create an instance of the service implementation
	dogProfileService := &DogProfileServiceImpl{}

	tests := []struct {
		expectedErr error
		in          *DogProfile
	}{
		{
			expectedErr: nil,
			in: &DogProfile{
				Name: "Koogla",
			},
		},
	}

	for _, tt := range tests {
		req := &DogProfile{
			Name: tt.in.Name,
		}

		_, err := dogProfileService.CreateDogProfile(context.Background(), req)

		if err != tt.expectedErr {
			t.Errorf("Expected Error is not the same!")
		}
	}

}
