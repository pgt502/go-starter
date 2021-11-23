package pb

import (
	context "context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type helloServiceServer struct {
}

func NewHelloServiceServer() HelloServiceServer {
	return &helloServiceServer{}
}

func (s *helloServiceServer) Hello(ctx context.Context, in *HelloRequest) (out *HelloResponse, err error) {
	if in == nil {
		err = status.Error(codes.InvalidArgument, "argument not provided")
		return
	}
	if in.Name == "" {
		err = status.Error(codes.InvalidArgument, "name not provided")
		return
	}
	out = &HelloResponse{
		Message: fmt.Sprintf("Hello %s", in.Name),
	}
	return
}
