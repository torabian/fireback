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

type GeoStatesClient interface {
	GeoStateActionCreate(ctx context.Context, in *GeoStateEntity, opts ...grpc.CallOption) (*GeoStateCreateReply, error)
	GeoStateActionUpdate(ctx context.Context, in *GeoStateEntity, opts ...grpc.CallOption) (*GeoStateCreateReply, error)
	GeoStateActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoStateQueryReply, error)
	GeoStateActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoStateReply, error)
	GeoStateActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error)
}

type geoStatesClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoStatesClient(cc grpc.ClientConnInterface) GeoStatesClient {
	return &geoStatesClient{cc}
}

func (c *geoStatesClient) GeoStateActionCreate(ctx context.Context, in *GeoStateEntity, opts ...grpc.CallOption) (*GeoStateCreateReply, error) {
	out := new(GeoStateCreateReply)
	err := c.cc.Invoke(ctx, "/GeoStates/GeoStateActionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoStatesClient) GeoStateActionUpdate(ctx context.Context, in *GeoStateEntity, opts ...grpc.CallOption) (*GeoStateCreateReply, error) {
	out := new(GeoStateCreateReply)
	err := c.cc.Invoke(ctx, "/GeoStates/GeoStateActionUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoStatesClient) GeoStateActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoStateQueryReply, error) {
	out := new(GeoStateQueryReply)
	err := c.cc.Invoke(ctx, "/GeoStates/GeoStateActionQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoStatesClient) GeoStateActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoStateReply, error) {
	out := new(GeoStateReply)
	err := c.cc.Invoke(ctx, "/GeoStates/GeoStateActionGetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoStatesClient) GeoStateActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error) {
	out := new(workspaces.RemoveReply)
	err := c.cc.Invoke(ctx, "/GeoStates/GeoStateActionRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type GeoStatesServer interface {
	GeoStateActionCreate(context.Context, *GeoStateEntity) (*GeoStateCreateReply, error)
	GeoStateActionUpdate(context.Context, *GeoStateEntity) (*GeoStateCreateReply, error)
	GeoStateActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoStateQueryReply, error)
	GeoStateActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoStateReply, error)
	GeoStateActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error)
	mustEmbedUnimplementedGeoStatesServer()
}

type UnimplementedGeoStatesServer struct {
}

func (UnimplementedGeoStatesServer) GeoStateActionCreate(context.Context, *GeoStateEntity) (*GeoStateCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoStateActionCreate not implemented")
}
func (UnimplementedGeoStatesServer) GeoStateActionUpdate(context.Context, *GeoStateEntity) (*GeoStateCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoStateActionUpdate not implemented")
}
func (UnimplementedGeoStatesServer) GeoStateActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoStateQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoStateActionQuery not implemented")
}
func (UnimplementedGeoStatesServer) GeoStateActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoStateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoStateActionGetOne not implemented")
}
func (UnimplementedGeoStatesServer) GeoStateActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoStateActionRemove not implemented")
}
func (UnimplementedGeoStatesServer) mustEmbedUnimplementedGeoStatesServer() {}

type UnsafeGeoStatesServer interface {
	mustEmbedUnimplementedGeoStatesServer()
}

func RegisterGeoStatesServer(s grpc.ServiceRegistrar, srv GeoStatesServer) {
	s.RegisterService(&GeoStates_ServiceDesc, srv)
}

func _GeoStates_GeoStateActionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoStateEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoStatesServer).GeoStateActionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoStates/GeoStateActionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoStatesServer).GeoStateActionCreate(ctx, req.(*GeoStateEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoStates_GeoStateActionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoStateEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoStatesServer).GeoStateActionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoStates/GeoStateActionUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoStatesServer).GeoStateActionUpdate(ctx, req.(*GeoStateEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoStates_GeoStateActionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoStatesServer).GeoStateActionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoStates/GeoStateActionQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoStatesServer).GeoStateActionQuery(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoStates_GeoStateActionGetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoStatesServer).GeoStateActionGetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoStates/GeoStateActionGetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoStatesServer).GeoStateActionGetOne(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoStates_GeoStateActionRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoStatesServer).GeoStateActionRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoStates/GeoStateActionRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoStatesServer).GeoStateActionRemove(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var GeoStates_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GeoStates",
	HandlerType: (*GeoStatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeoStateActionCreate",
			Handler:    _GeoStates_GeoStateActionCreate_Handler,
		},
		{
			MethodName: "GeoStateActionUpdate",
			Handler:    _GeoStates_GeoStateActionUpdate_Handler,
		},
		{
			MethodName: "GeoStateActionQuery",
			Handler:    _GeoStates_GeoStateActionQuery_Handler,
		},
		{
			MethodName: "GeoStateActionGetOne",
			Handler:    _GeoStates_GeoStateActionGetOne_Handler,
		},
		{
			MethodName: "GeoStateActionRemove",
			Handler:    _GeoStates_GeoStateActionRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/geo/GeoStateDefinitions.dyno.proto",
}
