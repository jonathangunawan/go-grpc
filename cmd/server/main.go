package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	grpchandler "github.com/jonathangunawan/go-grpc/handler/grpc"
	"github.com/jonathangunawan/go-grpc/infra"
	"github.com/jonathangunawan/go-grpc/interceptor"
	"github.com/jonathangunawan/go-grpc/pb"
	"github.com/jonathangunawan/go-grpc/repository"
	"github.com/jonathangunawan/go-grpc/usecase"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := infra.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// init depedency
	pr := repository.NewProductRepo()
	puc := usecase.NewProductUsecase(pr)
	h := grpchandler.NewProductHandler(puc)
	intr := interceptor.Interceptor{}
	srv := grpc.NewServer(grpc.UnaryInterceptor(intr.UnaryInterceptor))

	// register grpc handler
	pb.RegisterProductSvcServer(srv, h)

	// start tcp listener
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// setup graceful for grpc server
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-sigCh
		srv.GracefulStop()
		wg.Done()
	}()

	// assigning tcp listener to the grpc server
	srv.Serve(lis)

	// waiting for graceful process until done
	wg.Wait()
}
