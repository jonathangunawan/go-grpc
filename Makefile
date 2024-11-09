gen-protoc:
	protoc \
	--proto_path=./proto \
	--go_out=. \
	--go-grpc_out=. \
	proto/*.proto

run-srv:
	go run cmd/server/main.go