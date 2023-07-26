proto:
 protoc --go_out=. --go-grpc_out=. ./pkg/pb/cart.proto
server:
 go run cmd/main.go