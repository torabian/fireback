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

type GeoProvincesClient interface {
	GeoProvinceActionCreate(ctx context.Context, in *GeoProvinceEntity, opts ...grpc.CallOption) (*GeoProvinceCreateReply, error)
	GeoProvinceActionUpdate(ctx context.Context, in *GeoProvinceEntity, opts ...grpc.CallOption) (*GeoProvinceCreateReply, error)
	GeoProvinceActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoProvinceQueryReply, error)
	GeoProvinceActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoProvinceReply, error)
	GeoProvinceActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error)
}

type geoProvincesClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoProvincesClient(cc grpc.ClientConnInterface) GeoProvincesClient {
	return &geoProvincesClient{cc}
}

func (c *geoProvincesClient) GeoProvinceActionCreate(ctx context.Context, in *GeoProvinceEntity, opts ...grpc.CallOption) (*GeoProvinceCreateReply, error) {
	out := new(GeoProvinceCreateReply)
	err := c.cc.Invoke(ctx, "/GeoProvinces/GeoProvinceActionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoProvincesClient) GeoProvinceActionUpdate(ctx context.Context, in *GeoProvinceEntity, opts ...grpc.CallOption) (*GeoProvinceCreateReply, error) {
	out := new(GeoProvinceCreateReply)
	err := c.cc.Invoke(ctx, "/GeoProvinces/GeoProvinceActionUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoProvincesClient) GeoProvinceActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoProvinceQueryReply, error) {
	out := new(GeoProvinceQueryReply)
	err := c.cc.Invoke(ctx, "/GeoProvinces/GeoProvinceActionQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoProvincesClient) GeoProvinceActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoProvinceReply, error) {
	out := new(GeoProvinceReply)
	err := c.cc.Invoke(ctx, "/GeoProvinces/GeoProvinceActionGetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoProvincesClient) GeoProvinceActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error) {
	out := new(workspaces.RemoveReply)
	err := c.cc.Invoke(ctx, "/GeoProvinces/GeoProvinceActionRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type GeoProvincesServer interface {
	GeoProvinceActionCreate(context.Context, *GeoProvinceEntity) (*GeoProvinceCreateReply, error)
	GeoProvinceActionUpdate(context.Context, *GeoProvinceEntity) (*GeoProvinceCreateReply, error)
	GeoProvinceActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoProvinceQueryReply, error)
	GeoProvinceActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoProvinceReply, error)
	GeoProvinceActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error)
	mustEmbedUnimplementedGeoProvincesServer()
}

type UnimplementedGeoProvincesServer struct {
}

func (UnimplementedGeoProvincesServer) GeoProvinceActionCreate(context.Context, *GeoProvinceEntity) (*GeoProvinceCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoProvinceActionCreate not implemented")
}
func (UnimplementedGeoProvincesServer) GeoProvinceActionUpdate(context.Context, *GeoProvinceEntity) (*GeoProvinceCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoProvinceActionUpdate not implemented")
}
func (UnimplementedGeoProvincesServer) GeoProvinceActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoProvinceQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoProvinceActionQuery not implemented")
}
func (UnimplementedGeoProvincesServer) GeoProvinceActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoProvinceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoProvinceActionGetOne not implemented")
}
func (UnimplementedGeoProvincesServer) GeoProvinceActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoProvinceActionRemove not implemented")
}
func (UnimplementedGeoProvincesServer) mustEmbedUnimplementedGeoProvincesServer() {}

type UnsafeGeoProvincesServer interface {
	mustEmbedUnimplementedGeoProvincesServer()
}

func RegisterGeoProvincesServer(s grpc.ServiceRegistrar, srv GeoProvincesServer) {
	s.RegisterService(&GeoProvinces_ServiceDesc, srv)
}

func _GeoProvinces_GeoProvinceActionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoProvinceEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoProvincesServer).GeoProvinceActionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoProvinces/GeoProvinceActionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoProvincesServer).GeoProvinceActionCreate(ctx, req.(*GeoProvinceEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoProvinces_GeoProvinceActionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoProvinceEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoProvincesServer).GeoProvinceActionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoProvinces/GeoProvinceActionUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoProvincesServer).GeoProvinceActionUpdate(ctx, req.(*GeoProvinceEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoProvinces_GeoProvinceActionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoProvincesServer).GeoProvinceActionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoProvinces/GeoProvinceActionQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoProvincesServer).GeoProvinceActionQuery(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoProvinces_GeoProvinceActionGetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoProvincesServer).GeoProvinceActionGetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoProvinces/GeoProvinceActionGetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoProvincesServer).GeoProvinceActionGetOne(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoProvinces_GeoProvinceActionRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoProvincesServer).GeoProvinceActionRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoProvinces/GeoProvinceActionRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoProvincesServer).GeoProvinceActionRemove(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var GeoProvinces_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GeoProvinces",
	HandlerType: (*GeoProvincesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeoProvinceActionCreate",
			Handler:    _GeoProvinces_GeoProvinceActionCreate_Handler,
		},
		{
			MethodName: "GeoProvinceActionUpdate",
			Handler:    _GeoProvinces_GeoProvinceActionUpdate_Handler,
		},
		{
			MethodName: "GeoProvinceActionQuery",
			Handler:    _GeoProvinces_GeoProvinceActionQuery_Handler,
		},
		{
			MethodName: "GeoProvinceActionGetOne",
			Handler:    _GeoProvinces_GeoProvinceActionGetOne_Handler,
		},
		{
			MethodName: "GeoProvinceActionRemove",
			Handler:    _GeoProvinces_GeoProvinceActionRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/geo/GeoProvinceDefinitions.dyno.proto",
}
