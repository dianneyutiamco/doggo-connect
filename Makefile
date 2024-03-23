dog-profile:
	protoc -Iproto --go_opt=module=github.com/dianneyutiamco/doggo-connect/internal/api/dog_profile --go_out=internal/api/dog_profile --go-grpc_opt=module=github.com/dianneyutiamco/doggo-connect/internal/api/dog_profile --go-grpc_out=internal/api/dog_profile dog_profile.proto
	go build -o bin/calculator/server ./cmd/grpc_server