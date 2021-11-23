package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/pgt502/go-starter/internal/app/hello/pb"
	"github.com/pgt502/go-starter/internal/pkg/constants"
	"google.golang.org/grpc"
)

func main() {
	port, ok := os.LookupEnv(constants.HelloServicePortEnvVar)
	if !ok {
		log.Fatalf("environment variable %s not set - exiting...", constants.HelloServicePortEnvVar)
	}
	grpcService := pb.NewHelloServiceServer()
	listn, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to start a GRPC service on port: %s", port)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, grpcService)
	err = grpcServer.Serve(listn)
	if err != nil {
		log.Fatal("error serving GRPC service")
	}
}
