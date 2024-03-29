// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x31, 0x0e, 0xc2, 0x30,
	0x0c, 0x45, 0xa7, 0x32, 0x14, 0x58, 0x0a, 0x53, 0xef, 0x40, 0x8c, 0x02, 0x5c, 0xa0, 0x13, 0x33,
	0x6c, 0x6c, 0x49, 0x65, 0xa5, 0x91, 0x42, 0x6c, 0x62, 0x17, 0x89, 0xdb, 0x23, 0xd2, 0x89, 0xf1,
	0x3d, 0x3d, 0x7d, 0xfd, 0x76, 0x2b, 0x58, 0xde, 0x71, 0x44, 0xc3, 0x85, 0x94, 0xba, 0x66, 0xc2,
	0x94, 0xa8, 0x5f, 0xeb, 0x87, 0x51, 0x16, 0x67, 0x87, 0x76, 0x73, 0xfd, 0xd9, 0xfb, 0x52, 0x76,
	0xb6, 0x6d, 0x2a, 0x77, 0x3b, 0x53, 0x6b, 0x53, 0xe9, 0x86, 0xaf, 0x19, 0x45, 0xfb, 0xfd, 0xbf,
	0x14, 0xa6, 0x2c, 0x38, 0x9c, 0x1f, 0x36, 0x44, 0x9d, 0x66, 0x6f, 0x46, 0x7a, 0x02, 0x07, 0xbd,
	0x1c, 0x2d, 0x04, 0x3a, 0x88, 0xba, 0xa2, 0x58, 0x20, 0x66, 0xc5, 0x92, 0x5d, 0x02, 0xc7, 0x0c,
	0x75, 0x00, 0xd8, 0xfb, 0x55, 0x3d, 0x70, 0xfa, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xdf, 0xbf,
	0x3e, 0xa5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/hello.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
