// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: token_position.proto

package api

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

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionServiceClient interface {
	GetTokenPositions(ctx context.Context, in *GetTokenPositionsRequest, opts ...grpc.CallOption) (*GetTokenPositionsResponse, error)
	GetTokenPosition(ctx context.Context, in *GetTokenPositionRequest, opts ...grpc.CallOption) (*GetTokenPositionResponse, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) GetTokenPositions(ctx context.Context, in *GetTokenPositionsRequest, opts ...grpc.CallOption) (*GetTokenPositionsResponse, error) {
	out := new(GetTokenPositionsResponse)
	err := c.cc.Invoke(ctx, "/gmx.neutral.query.api.PositionService/GetTokenPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) GetTokenPosition(ctx context.Context, in *GetTokenPositionRequest, opts ...grpc.CallOption) (*GetTokenPositionResponse, error) {
	out := new(GetTokenPositionResponse)
	err := c.cc.Invoke(ctx, "/gmx.neutral.query.api.PositionService/GetTokenPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServiceServer is the server API for PositionService service.
// All implementations should embed UnimplementedPositionServiceServer
// for forward compatibility
type PositionServiceServer interface {
	GetTokenPositions(context.Context, *GetTokenPositionsRequest) (*GetTokenPositionsResponse, error)
	GetTokenPosition(context.Context, *GetTokenPositionRequest) (*GetTokenPositionResponse, error)
}

// UnimplementedPositionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (UnimplementedPositionServiceServer) GetTokenPositions(context.Context, *GetTokenPositionsRequest) (*GetTokenPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenPositions not implemented")
}
func (UnimplementedPositionServiceServer) GetTokenPosition(context.Context, *GetTokenPositionRequest) (*GetTokenPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenPosition not implemented")
}

// UnsafePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServiceServer will
// result in compilation errors.
type UnsafePositionServiceServer interface {
	mustEmbedUnimplementedPositionServiceServer()
}

func RegisterPositionServiceServer(s grpc.ServiceRegistrar, srv PositionServiceServer) {
	s.RegisterService(&PositionService_ServiceDesc, srv)
}

func _PositionService_GetTokenPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetTokenPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmx.neutral.query.api.PositionService/GetTokenPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetTokenPositions(ctx, req.(*GetTokenPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_GetTokenPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetTokenPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gmx.neutral.query.api.PositionService/GetTokenPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetTokenPosition(ctx, req.(*GetTokenPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PositionService_ServiceDesc is the grpc.ServiceDesc for PositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gmx.neutral.query.api.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokenPositions",
			Handler:    _PositionService_GetTokenPositions_Handler,
		},
		{
			MethodName: "GetTokenPosition",
			Handler:    _PositionService_GetTokenPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token_position.proto",
}
