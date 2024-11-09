# Go gRPC
This is an example of how to implement gRPC in Go as server and client
- [x] Unary RCP
- [ ] Server Streaming RPC
- [ ] Client Streaming RCP
- [ ] Bidirectional RPC

## Project Specification
- go: 1.22.4
- protobuf: proto3
- protoc-gen-go: v1.35.1
- protoc-gen-go-grpc: v1.5.1

## Notes
1. At least for now, repository isn't used that much. I just create the layer as an act of consistency :)
2. Since this project is focused for the gRPC implementation in go. Some of the code didn't consider about concurrency safe
3. We can split the proto file into separated git repo, so it will not to dependent with one services