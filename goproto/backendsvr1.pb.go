// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backendsvr1.proto

package gwproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("backendsvr1.proto", fileDescriptor_9b877adbc81cefbe) }

var fileDescriptor_9b877adbc81cefbe = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4a, 0x4c, 0xce,
	0x4e, 0xcd, 0x4b, 0x29, 0x2e, 0x2b, 0x32, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f,
	0x2f, 0x07, 0x33, 0xa4, 0x38, 0x73, 0x8b, 0xd3, 0x21, 0x62, 0x46, 0x5e, 0x5c, 0xdc, 0x4e, 0x08,
	0x85, 0x42, 0xd6, 0x5c, 0xec, 0xbe, 0xa9, 0x25, 0x19, 0xf9, 0x29, 0x86, 0x42, 0xe2, 0x7a, 0x50,
	0xe5, 0x7a, 0x50, 0x91, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29, 0x51, 0x4c, 0x89, 0x82,
	0x9c, 0x4a, 0x25, 0x06, 0x27, 0x0d, 0x2e, 0xc1, 0xe4, 0xfc, 0x5c, 0xbd, 0xaa, 0xc2, 0xbc, 0xd4,
	0x12, 0x3d, 0xf7, 0x70, 0xb0, 0x1a, 0x27, 0x76, 0x28, 0x23, 0x80, 0x71, 0x11, 0x13, 0x8c, 0x9d,
	0xc4, 0x06, 0xa6, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xf9, 0x31, 0xfd, 0xa5, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Backendsvr1Client is the client API for Backendsvr1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Backendsvr1Client interface {
	Method1(ctx context.Context, in *Method1Request, opts ...grpc.CallOption) (*Method1Reply, error)
}

type backendsvr1Client struct {
	cc *grpc.ClientConn
}

func NewBackendsvr1Client(cc *grpc.ClientConn) Backendsvr1Client {
	return &backendsvr1Client{cc}
}

func (c *backendsvr1Client) Method1(ctx context.Context, in *Method1Request, opts ...grpc.CallOption) (*Method1Reply, error) {
	out := new(Method1Reply)
	err := c.cc.Invoke(ctx, "/gwproto.Backendsvr1/Method1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Backendsvr1Server is the server API for Backendsvr1 service.
type Backendsvr1Server interface {
	Method1(context.Context, *Method1Request) (*Method1Reply, error)
}

// UnimplementedBackendsvr1Server can be embedded to have forward compatible implementations.
type UnimplementedBackendsvr1Server struct {
}

func (*UnimplementedBackendsvr1Server) Method1(ctx context.Context, req *Method1Request) (*Method1Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Method1 not implemented")
}

func RegisterBackendsvr1Server(s *grpc.Server, srv Backendsvr1Server) {
	s.RegisterService(&_Backendsvr1_serviceDesc, srv)
}

func _Backendsvr1_Method1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Method1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Backendsvr1Server).Method1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwproto.Backendsvr1/Method1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Backendsvr1Server).Method1(ctx, req.(*Method1Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backendsvr1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gwproto.Backendsvr1",
	HandlerType: (*Backendsvr1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Method1",
			Handler:    _Backendsvr1_Method1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backendsvr1.proto",
}
