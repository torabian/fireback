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

type GeoCountrysClient interface {
	GeoCountryActionCreate(ctx context.Context, in *GeoCountryEntity, opts ...grpc.CallOption) (*GeoCountryCreateReply, error)
	GeoCountryActionUpdate(ctx context.Context, in *GeoCountryEntity, opts ...grpc.CallOption) (*GeoCountryCreateReply, error)
	GeoCountryActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoCountryQueryReply, error)
	GeoCountryActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoCountryReply, error)
	GeoCountryActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error)
}

type geoCountrysClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoCountrysClient(cc grpc.ClientConnInterface) GeoCountrysClient {
	return &geoCountrysClient{cc}
}

func (c *geoCountrysClient) GeoCountryActionCreate(ctx context.Context, in *GeoCountryEntity, opts ...grpc.CallOption) (*GeoCountryCreateReply, error) {
	out := new(GeoCountryCreateReply)
	err := c.cc.Invoke(ctx, "/GeoCountrys/GeoCountryActionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCountrysClient) GeoCountryActionUpdate(ctx context.Context, in *GeoCountryEntity, opts ...grpc.CallOption) (*GeoCountryCreateReply, error) {
	out := new(GeoCountryCreateReply)
	err := c.cc.Invoke(ctx, "/GeoCountrys/GeoCountryActionUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCountrysClient) GeoCountryActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoCountryQueryReply, error) {
	out := new(GeoCountryQueryReply)
	err := c.cc.Invoke(ctx, "/GeoCountrys/GeoCountryActionQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCountrysClient) GeoCountryActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoCountryReply, error) {
	out := new(GeoCountryReply)
	err := c.cc.Invoke(ctx, "/GeoCountrys/GeoCountryActionGetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCountrysClient) GeoCountryActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error) {
	out := new(workspaces.RemoveReply)
	err := c.cc.Invoke(ctx, "/GeoCountrys/GeoCountryActionRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type GeoCountrysServer interface {
	GeoCountryActionCreate(context.Context, *GeoCountryEntity) (*GeoCountryCreateReply, error)
	GeoCountryActionUpdate(context.Context, *GeoCountryEntity) (*GeoCountryCreateReply, error)
	GeoCountryActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoCountryQueryReply, error)
	GeoCountryActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoCountryReply, error)
	GeoCountryActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error)
	mustEmbedUnimplementedGeoCountrysServer()
}

type UnimplementedGeoCountrysServer struct {
}

func (UnimplementedGeoCountrysServer) GeoCountryActionCreate(context.Context, *GeoCountryEntity) (*GeoCountryCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCountryActionCreate not implemented")
}
func (UnimplementedGeoCountrysServer) GeoCountryActionUpdate(context.Context, *GeoCountryEntity) (*GeoCountryCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCountryActionUpdate not implemented")
}
func (UnimplementedGeoCountrysServer) GeoCountryActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoCountryQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCountryActionQuery not implemented")
}
func (UnimplementedGeoCountrysServer) GeoCountryActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoCountryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCountryActionGetOne not implemented")
}
func (UnimplementedGeoCountrysServer) GeoCountryActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCountryActionRemove not implemented")
}
func (UnimplementedGeoCountrysServer) mustEmbedUnimplementedGeoCountrysServer() {}

type UnsafeGeoCountrysServer interface {
	mustEmbedUnimplementedGeoCountrysServer()
}

func RegisterGeoCountrysServer(s grpc.ServiceRegistrar, srv GeoCountrysServer) {
	s.RegisterService(&GeoCountrys_ServiceDesc, srv)
}

func _GeoCountrys_GeoCountryActionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoCountryEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCountrysServer).GeoCountryActionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoCountrys/GeoCountryActionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCountrysServer).GeoCountryActionCreate(ctx, req.(*GeoCountryEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCountrys_GeoCountryActionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoCountryEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCountrysServer).GeoCountryActionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoCountrys/GeoCountryActionUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCountrysServer).GeoCountryActionUpdate(ctx, req.(*GeoCountryEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCountrys_GeoCountryActionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCountrysServer).GeoCountryActionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoCountrys/GeoCountryActionQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCountrysServer).GeoCountryActionQuery(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCountrys_GeoCountryActionGetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCountrysServer).GeoCountryActionGetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoCountrys/GeoCountryActionGetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCountrysServer).GeoCountryActionGetOne(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCountrys_GeoCountryActionRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCountrysServer).GeoCountryActionRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoCountrys/GeoCountryActionRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCountrysServer).GeoCountryActionRemove(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var GeoCountrys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GeoCountrys",
	HandlerType: (*GeoCountrysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeoCountryActionCreate",
			Handler:    _GeoCountrys_GeoCountryActionCreate_Handler,
		},
		{
			MethodName: "GeoCountryActionUpdate",
			Handler:    _GeoCountrys_GeoCountryActionUpdate_Handler,
		},
		{
			MethodName: "GeoCountryActionQuery",
			Handler:    _GeoCountrys_GeoCountryActionQuery_Handler,
		},
		{
			MethodName: "GeoCountryActionGetOne",
			Handler:    _GeoCountrys_GeoCountryActionGetOne_Handler,
		},
		{
			MethodName: "GeoCountryActionRemove",
			Handler:    _GeoCountrys_GeoCountryActionRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/geo/GeoCountryDefinitions.dyno.proto",
}
