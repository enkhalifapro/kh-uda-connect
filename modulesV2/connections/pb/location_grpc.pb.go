// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: pb/location.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LocationsService_GetLocations_FullMethodName = "/example.LocationsService/GetLocations"
)

// LocationsServiceClient is the client API for LocationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationsServiceClient interface {
	GetLocations(ctx context.Context, in *LocationsRequest, opts ...grpc.CallOption) (*LocationsList, error)
}

type locationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationsServiceClient(cc grpc.ClientConnInterface) LocationsServiceClient {
	return &locationsServiceClient{cc}
}

func (c *locationsServiceClient) GetLocations(ctx context.Context, in *LocationsRequest, opts ...grpc.CallOption) (*LocationsList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LocationsList)
	err := c.cc.Invoke(ctx, LocationsService_GetLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationsServiceServer is the server API for LocationsService service.
// All implementations must embed UnimplementedLocationsServiceServer
// for forward compatibility.
type LocationsServiceServer interface {
	GetLocations(context.Context, *LocationsRequest) (*LocationsList, error)
	mustEmbedUnimplementedLocationsServiceServer()
}

// UnimplementedLocationsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLocationsServiceServer struct{}

func (UnimplementedLocationsServiceServer) GetLocations(context.Context, *LocationsRequest) (*LocationsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocations not implemented")
}
func (UnimplementedLocationsServiceServer) mustEmbedUnimplementedLocationsServiceServer() {}
func (UnimplementedLocationsServiceServer) testEmbeddedByValue()                          {}

// UnsafeLocationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationsServiceServer will
// result in compilation errors.
type UnsafeLocationsServiceServer interface {
	mustEmbedUnimplementedLocationsServiceServer()
}

func RegisterLocationsServiceServer(s grpc.ServiceRegistrar, srv LocationsServiceServer) {
	// If the following call pancis, it indicates UnimplementedLocationsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LocationsService_ServiceDesc, srv)
}

func _LocationsService_GetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServiceServer).GetLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationsService_GetLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServiceServer).GetLocations(ctx, req.(*LocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationsService_ServiceDesc is the grpc.ServiceDesc for LocationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.LocationsService",
	HandlerType: (*LocationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocations",
			Handler:    _LocationsService_GetLocations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/location.proto",
}