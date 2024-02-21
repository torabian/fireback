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

type GeoLocationTypesClient interface {
	GeoLocationTypeActionCreate(ctx context.Context, in *GeoLocationTypeEntity, opts ...grpc.CallOption) (*GeoLocationTypeCreateReply, error)
	GeoLocationTypeActionUpdate(ctx context.Context, in *GeoLocationTypeEntity, opts ...grpc.CallOption) (*GeoLocationTypeCreateReply, error)
	GeoLocationTypeActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoLocationTypeQueryReply, error)
	GeoLocationTypeActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoLocationTypeReply, error)
	GeoLocationTypeActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error)
}

type geoLocationTypesClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoLocationTypesClient(cc grpc.ClientConnInterface) GeoLocationTypesClient {
	return &geoLocationTypesClient{cc}
}

func (c *geoLocationTypesClient) GeoLocationTypeActionCreate(ctx context.Context, in *GeoLocationTypeEntity, opts ...grpc.CallOption) (*GeoLocationTypeCreateReply, error) {
	out := new(GeoLocationTypeCreateReply)
	err := c.cc.Invoke(ctx, "/GeoLocationTypes/GeoLocationTypeActionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoLocationTypesClient) GeoLocationTypeActionUpdate(ctx context.Context, in *GeoLocationTypeEntity, opts ...grpc.CallOption) (*GeoLocationTypeCreateReply, error) {
	out := new(GeoLocationTypeCreateReply)
	err := c.cc.Invoke(ctx, "/GeoLocationTypes/GeoLocationTypeActionUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoLocationTypesClient) GeoLocationTypeActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoLocationTypeQueryReply, error) {
	out := new(GeoLocationTypeQueryReply)
	err := c.cc.Invoke(ctx, "/GeoLocationTypes/GeoLocationTypeActionQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoLocationTypesClient) GeoLocationTypeActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoLocationTypeReply, error) {
	out := new(GeoLocationTypeReply)
	err := c.cc.Invoke(ctx, "/GeoLocationTypes/GeoLocationTypeActionGetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoLocationTypesClient) GeoLocationTypeActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error) {
	out := new(workspaces.RemoveReply)
	err := c.cc.Invoke(ctx, "/GeoLocationTypes/GeoLocationTypeActionRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type GeoLocationTypesServer interface {
	GeoLocationTypeActionCreate(context.Context, *GeoLocationTypeEntity) (*GeoLocationTypeCreateReply, error)
	GeoLocationTypeActionUpdate(context.Context, *GeoLocationTypeEntity) (*GeoLocationTypeCreateReply, error)
	GeoLocationTypeActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoLocationTypeQueryReply, error)
	GeoLocationTypeActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoLocationTypeReply, error)
	GeoLocationTypeActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error)
	mustEmbedUnimplementedGeoLocationTypesServer()
}

type UnimplementedGeoLocationTypesServer struct {
}

func (UnimplementedGeoLocationTypesServer) GeoLocationTypeActionCreate(context.Context, *GeoLocationTypeEntity) (*GeoLocationTypeCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocationTypeActionCreate not implemented")
}
func (UnimplementedGeoLocationTypesServer) GeoLocationTypeActionUpdate(context.Context, *GeoLocationTypeEntity) (*GeoLocationTypeCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocationTypeActionUpdate not implemented")
}
func (UnimplementedGeoLocationTypesServer) GeoLocationTypeActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoLocationTypeQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocationTypeActionQuery not implemented")
}
func (UnimplementedGeoLocationTypesServer) GeoLocationTypeActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoLocationTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocationTypeActionGetOne not implemented")
}
func (UnimplementedGeoLocationTypesServer) GeoLocationTypeActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoLocationTypeActionRemove not implemented")
}
func (UnimplementedGeoLocationTypesServer) mustEmbedUnimplementedGeoLocationTypesServer() {}

type UnsafeGeoLocationTypesServer interface {
	mustEmbedUnimplementedGeoLocationTypesServer()
}

func RegisterGeoLocationTypesServer(s grpc.ServiceRegistrar, srv GeoLocationTypesServer) {
	s.RegisterService(&GeoLocationTypes_ServiceDesc, srv)
}

func _GeoLocationTypes_GeoLocationTypeActionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoLocationTypeEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoLocationTypesServer).GeoLocationTypeActionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoLocationTypes/GeoLocationTypeActionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoLocationTypesServer).GeoLocationTypeActionCreate(ctx, req.(*GeoLocationTypeEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoLocationTypes_GeoLocationTypeActionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoLocationTypeEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoLocationTypesServer).GeoLocationTypeActionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoLocationTypes/GeoLocationTypeActionUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoLocationTypesServer).GeoLocationTypeActionUpdate(ctx, req.(*GeoLocationTypeEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoLocationTypes_GeoLocationTypeActionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoLocationTypesServer).GeoLocationTypeActionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoLocationTypes/GeoLocationTypeActionQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoLocationTypesServer).GeoLocationTypeActionQuery(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoLocationTypes_GeoLocationTypeActionGetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoLocationTypesServer).GeoLocationTypeActionGetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoLocationTypes/GeoLocationTypeActionGetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoLocationTypesServer).GeoLocationTypeActionGetOne(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoLocationTypes_GeoLocationTypeActionRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoLocationTypesServer).GeoLocationTypeActionRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoLocationTypes/GeoLocationTypeActionRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoLocationTypesServer).GeoLocationTypeActionRemove(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var GeoLocationTypes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GeoLocationTypes",
	HandlerType: (*GeoLocationTypesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeoLocationTypeActionCreate",
			Handler:    _GeoLocationTypes_GeoLocationTypeActionCreate_Handler,
		},
		{
			MethodName: "GeoLocationTypeActionUpdate",
			Handler:    _GeoLocationTypes_GeoLocationTypeActionUpdate_Handler,
		},
		{
			MethodName: "GeoLocationTypeActionQuery",
			Handler:    _GeoLocationTypes_GeoLocationTypeActionQuery_Handler,
		},
		{
			MethodName: "GeoLocationTypeActionGetOne",
			Handler:    _GeoLocationTypes_GeoLocationTypeActionGetOne_Handler,
		},
		{
			MethodName: "GeoLocationTypeActionRemove",
			Handler:    _GeoLocationTypes_GeoLocationTypeActionRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/geo/GeoLocationTypeDefinitions.dyno.proto",
}
