//
//
//
//
//

package geo

import (
	context "context"

	"github.com/torabian/fireback/modules/workspaces"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion7

type LocationDatasClient interface {
	LocationDataActionCreate(ctx context.Context, in *LocationDataEntity, opts ...grpc.CallOption) (*LocationDataCreateReply, error)
	LocationDataActionUpdate(ctx context.Context, in *LocationDataEntity, opts ...grpc.CallOption) (*LocationDataCreateReply, error)
	LocationDataActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*LocationDataQueryReply, error)
	LocationDataActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*LocationDataReply, error)
	LocationDataActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error)
}

type locationDatasClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationDatasClient(cc grpc.ClientConnInterface) LocationDatasClient {
	return &locationDatasClient{cc}
}

func (c *locationDatasClient) LocationDataActionCreate(ctx context.Context, in *LocationDataEntity, opts ...grpc.CallOption) (*LocationDataCreateReply, error) {
	out := new(LocationDataCreateReply)
	err := c.cc.Invoke(ctx, "/LocationDatas/LocationDataActionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDatasClient) LocationDataActionUpdate(ctx context.Context, in *LocationDataEntity, opts ...grpc.CallOption) (*LocationDataCreateReply, error) {
	out := new(LocationDataCreateReply)
	err := c.cc.Invoke(ctx, "/LocationDatas/LocationDataActionUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDatasClient) LocationDataActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*LocationDataQueryReply, error) {
	out := new(LocationDataQueryReply)
	err := c.cc.Invoke(ctx, "/LocationDatas/LocationDataActionQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDatasClient) LocationDataActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*LocationDataReply, error) {
	out := new(LocationDataReply)
	err := c.cc.Invoke(ctx, "/LocationDatas/LocationDataActionGetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationDatasClient) LocationDataActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error) {
	out := new(workspaces.RemoveReply)
	err := c.cc.Invoke(ctx, "/LocationDatas/LocationDataActionRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type LocationDatasServer interface {
	LocationDataActionCreate(context.Context, *LocationDataEntity) (*LocationDataCreateReply, error)
	LocationDataActionUpdate(context.Context, *LocationDataEntity) (*LocationDataCreateReply, error)
	LocationDataActionQuery(context.Context, *workspaces.QueryFilterRequest) (*LocationDataQueryReply, error)
	LocationDataActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*LocationDataReply, error)
	LocationDataActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error)
	mustEmbedUnimplementedLocationDatasServer()
}

type UnimplementedLocationDatasServer struct {
}

func (UnimplementedLocationDatasServer) LocationDataActionCreate(context.Context, *LocationDataEntity) (*LocationDataCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocationDataActionCreate not implemented")
}
func (UnimplementedLocationDatasServer) LocationDataActionUpdate(context.Context, *LocationDataEntity) (*LocationDataCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocationDataActionUpdate not implemented")
}
func (UnimplementedLocationDatasServer) LocationDataActionQuery(context.Context, *workspaces.QueryFilterRequest) (*LocationDataQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocationDataActionQuery not implemented")
}
func (UnimplementedLocationDatasServer) LocationDataActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*LocationDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocationDataActionGetOne not implemented")
}
func (UnimplementedLocationDatasServer) LocationDataActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocationDataActionRemove not implemented")
}
func (UnimplementedLocationDatasServer) mustEmbedUnimplementedLocationDatasServer() {}

type UnsafeLocationDatasServer interface {
	mustEmbedUnimplementedLocationDatasServer()
}

func RegisterLocationDatasServer(s grpc.ServiceRegistrar, srv LocationDatasServer) {
	s.RegisterService(&LocationDatas_ServiceDesc, srv)
}

func _LocationDatas_LocationDataActionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationDataEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationDatasServer).LocationDataActionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LocationDatas/LocationDataActionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationDatasServer).LocationDataActionCreate(ctx, req.(*LocationDataEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationDatas_LocationDataActionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationDataEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationDatasServer).LocationDataActionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LocationDatas/LocationDataActionUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationDatasServer).LocationDataActionUpdate(ctx, req.(*LocationDataEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationDatas_LocationDataActionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationDatasServer).LocationDataActionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LocationDatas/LocationDataActionQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationDatasServer).LocationDataActionQuery(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationDatas_LocationDataActionGetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationDatasServer).LocationDataActionGetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LocationDatas/LocationDataActionGetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationDatasServer).LocationDataActionGetOne(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationDatas_LocationDataActionRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationDatasServer).LocationDataActionRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LocationDatas/LocationDataActionRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationDatasServer).LocationDataActionRemove(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var LocationDatas_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LocationDatas",
	HandlerType: (*LocationDatasServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LocationDataActionCreate",
			Handler:    _LocationDatas_LocationDataActionCreate_Handler,
		},
		{
			MethodName: "LocationDataActionUpdate",
			Handler:    _LocationDatas_LocationDataActionUpdate_Handler,
		},
		{
			MethodName: "LocationDataActionQuery",
			Handler:    _LocationDatas_LocationDataActionQuery_Handler,
		},
		{
			MethodName: "LocationDataActionGetOne",
			Handler:    _LocationDatas_LocationDataActionGetOne_Handler,
		},
		{
			MethodName: "LocationDataActionRemove",
			Handler:    _LocationDatas_LocationDataActionRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/geo/LocationDataDefinitions.dyno.proto",
}
