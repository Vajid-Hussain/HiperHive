run: 
	go run cmd/main.go

protoServer:
	protoc ./pkg/pb/server.proto --go_out=./ --go-grpc_out=./

autoserver:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go



protoAuthService:
	protoc ./pkg/pb/auth.proto --go_out=./ --go-grpc_out=./