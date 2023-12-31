// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: ProtocolB/server.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// Response_RMQClient is the client API for Response_RMQ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Response_RMQClient interface {
	V_Cupos(ctx context.Context, in *Cupos, opts ...grpc.CallOption) (*Interesados, error)
}

type response_RMQClient struct {
	cc grpc.ClientConnInterface
}

func NewResponse_RMQClient(cc grpc.ClientConnInterface) Response_RMQClient {
	return &response_RMQClient{cc}
}

func (c *response_RMQClient) V_Cupos(ctx context.Context, in *Cupos, opts ...grpc.CallOption) (*Interesados, error) {
	out := new(Interesados)
	err := c.cc.Invoke(ctx, "/ProtocolB.response_RMQ/V_Cupos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Response_RMQServer is the server API for Response_RMQ service.
// All implementations must embed UnimplementedResponse_RMQServer
// for forward compatibility
type Response_RMQServer interface {
	V_Cupos(context.Context, *Cupos) (*Interesados, error)
	mustEmbedUnimplementedResponse_RMQServer()
}

// UnimplementedResponse_RMQServer must be embedded to have forward compatible implementations.
type UnimplementedResponse_RMQServer struct {
}

func (UnimplementedResponse_RMQServer) V_Cupos(context.Context, *Cupos) (*Interesados, error) {
	return nil, status.Errorf(codes.Unimplemented, "method V_Cupos not implemented")
}
func (UnimplementedResponse_RMQServer) mustEmbedUnimplementedResponse_RMQServer() {}

// UnsafeResponse_RMQServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Response_RMQServer will
// result in compilation errors.
type UnsafeResponse_RMQServer interface {
	mustEmbedUnimplementedResponse_RMQServer()
}

func RegisterResponse_RMQServer(s grpc.ServiceRegistrar, srv Response_RMQServer) {
	s.RegisterService(&Response_RMQ_ServiceDesc, srv)
}

func _Response_RMQ_V_Cupos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cupos)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Response_RMQServer).V_Cupos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProtocolB.response_RMQ/V_Cupos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Response_RMQServer).V_Cupos(ctx, req.(*Cupos))
	}
	return interceptor(ctx, in, info, handler)
}

// Response_RMQ_ServiceDesc is the grpc.ServiceDesc for Response_RMQ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Response_RMQ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProtocolB.response_RMQ",
	HandlerType: (*Response_RMQServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "V_Cupos",
			Handler:    _Response_RMQ_V_Cupos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ProtocolB/server.proto",
}
