package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jonathangunawan/go-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcServer string = "localhost:8080"
)

func main() {
	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())) // without TLS

	conn, err := grpc.NewClient(grpcServer, opts...)
	if err != nil {
		log.Fatal(err)
	}

	// init grpc client
	client := pb.NewProductSvcClient(conn)

	// call the server (usually in repository layer or at the same layer)
	res, err := client.GetAllProduct(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
