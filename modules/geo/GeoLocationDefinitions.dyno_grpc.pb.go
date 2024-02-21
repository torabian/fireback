//
//
//
//
//

package geo

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"pixelplux.com/fireback/modules/workspaces"
)

const _ = grpc.SupportPackageIsVersion7

type GeoLocationsClient interface {
	GeoLocationActionCreate(ctx context.Context, in *GeoLocationEntity, opts ...grpc.CallOption) (*GeoLocationCreateReply, error)
	GeoLocationActionUpdate(ctx context.Context, in *GeoLocationEntity, opts ...grpc.CallOption) (*GeoLocationCreateReply, error)
	GeoLocationActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoLocationQueryReply, error)
	GeoLocationActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoLocationReply, error)
	GeoLocationActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error)
}

type geoLocationsClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoLocationsClient(cc grpc.ClientConnInterface) GeoLocationsClient {
	return &geoLocationsClient{cc}
}

func (c *geoLocationsClient) GeoLocationActionCreate(ctx context.Context, in *GeoLocationEntity, opts ...grpc.CallOption) (*GeoLocationCreateReply, error) {
	out := new(GeoLocationCreateReply)
	err := c.cc.Invoke(ctx, "/GeoLocations/GeoLocationActionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoLocationsClient) GeoLocationActionUpdate(ctx context.Context, in *GeoLocationEntity, opts ...grpc.CallOption) (*GeoLocationCreateReply, error) {
	out := new(GeoLocationCreateReply)
	err := c.cc.Invoke(ctx, "/GeoLocations/GeoLocationActionUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoLocationsClient) GeoLocationActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoLocationQueryReply, error) {
	out := new(GeoLocationQueryReply)
	err := c.cc.Invoke(ctx, "/GeoLocations/GeoLocationActionQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoLocationsClient) GeoLocationActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoLocationReply, error) {
	out := new(GeoLocationReply)
	err := c.cc.Invoke(ctx, "/GeoLocations/GeoLocationActionGetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoLocationsClient) GeoLocationActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error) {
	out := new(workspaces.RemoveReply)
	err := c.cc.Invoke(ctx, "/GeoLocations/GeoLocationActionRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type GeoLocationsServer interface {
	GeoLocationActionCreate(context.Context, *GeoLocationEntity) (*GeoLocationCreateReply, error)
	GeoLocationActionUpdate(context.Context, *GeoLocationEntity) (*GeoLocationCreateReply, error)
	GeoLocationActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoLocationQueryReply, error)
	GeoLocationActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoLocationReply, error)
	GeoLocationActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error)
	mustEmbedUnimplementedGeoLocationsServer()
}

type UnimplementedGeoLocationsServer struct {
}

func (UnimplementedGeoLocationsServer) GeoLocationActionCreate(context.Context, *GeoLocationEntity) (*GeoLocationCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocationActionCreate not implemented")
}
func (UnimplementedGeoLocationsServer) GeoLocationActionUpdate(context.Context, *GeoLocationEntity) (*GeoLocationCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocationActionUpdate not implemented")
}
func (UnimplementedGeoLocationsServer) GeoLocationActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoLocationQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocationActionQuery not implemented")
}
func (UnimplementedGeoLocationsServer) GeoLocationActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoLocationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocationActionGetOne not implemented")
}
func (UnimplementedGeoLocationsServer) GeoLocationActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocationActionRemove not implemented")
}
func (UnimplementedGeoLocationsServer) mustEmbedUnimplementedGeoLocationsServer() {}

type UnsafeGeoLocationsServer interface {
	mustEmbedUnimplementedGeoLocationsServer()
}

func RegisterGeoLocationsServer(s grpc.ServiceRegistrar, srv GeoLocationsServer) {
	s.RegisterService(&GeoLocations_ServiceDesc, srv)
}

func _GeoLocations_GeoLocationActionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoLocationEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoLocationsServer).GeoLocationActionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoLocations/GeoLocationActionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoLocationsServer).GeoLocationActionCreate(ctx, req.(*GeoLocationEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoLocations_GeoLocationActionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoLocationEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoLocationsServer).GeoLocationActionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoLocations/GeoLocationActionUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoLocationsServer).GeoLocationActionUpdate(ctx, req.(*GeoLocationEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoLocations_GeoLocationActionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoLocationsServer).GeoLocationActionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoLocations/GeoLocationActionQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoLocationsServer).GeoLocationActionQuery(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoLocations_GeoLocationActionGetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoLocationsServer).GeoLocationActionGetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoLocations/GeoLocationActionGetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoLocationsServer).GeoLocationActionGetOne(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoLocations_GeoLocationActionRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoLocationsServer).GeoLocationActionRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoLocations/GeoLocationActionRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoLocationsServer).GeoLocationActionRemove(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var GeoLocations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GeoLocations",
	HandlerType: (*GeoLocationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeoLocationActionCreate",
			Handler:    _GeoLocations_GeoLocationActionCreate_Handler,
		},
		{
			MethodName: "GeoLocationActionUpdate",
			Handler:    _GeoLocations_GeoLocationActionUpdate_Handler,
		},
		{
			MethodName: "GeoLocationActionQuery",
			Handler:    _GeoLocations_GeoLocationActionQuery_Handler,
		},
		{
			MethodName: "GeoLocationActionGetOne",
			Handler:    _GeoLocations_GeoLocationActionGetOne_Handler,
		},
		{
			MethodName: "GeoLocationActionRemove",
			Handler:    _GeoLocations_GeoLocationActionRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/geo/GeoLocationDefinitions.dyno.proto",
}
