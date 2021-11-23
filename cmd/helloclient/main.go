package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/pgt502/go-starter/internal/app/hello/pb"
	"github.com/pgt502/go-starter/internal/pkg/constants"
	"google.golang.org/grpc"
)

func main() {
	port, ok := os.LookupEnv(constants.HelloServicePortEnvVar)
	if !ok {
		log.Fatalf("environment variable %s not set - exiting...", constants.HelloServicePortEnvVar)
	}
	url := fmt.Sprintf(":%s", port)
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to GRPC service at %s", url)
	}
	defer conn.Close()
	helloClient := pb.NewHelloServiceClient(conn)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	i := 0

	for {
		select {
		case <-c:
			log.Print("shutting down the client...")
			os.Exit(0)
		case <-time.Tick(time.Second):
			log.Printf("sending a request to the server...")
			req := &pb.HelloRequest{Name: fmt.Sprintf("%d", i)}
			resp, err := helloClient.Hello(context.TODO(), req)
			if err != nil {
				log.Printf("***error sending request to the server: %s", err)
				break
			}
			i++
			log.Printf("response from server: %+v", resp)
		}
	}
}
