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

type GeoCitysClient interface {
	GeoCityActionCreate(ctx context.Context, in *GeoCityEntity, opts ...grpc.CallOption) (*GeoCityCreateReply, error)
	GeoCityActionUpdate(ctx context.Context, in *GeoCityEntity, opts ...grpc.CallOption) (*GeoCityCreateReply, error)
	GeoCityActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoCityQueryReply, error)
	GeoCityActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoCityReply, error)
	GeoCityActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error)
}

type geoCitysClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoCitysClient(cc grpc.ClientConnInterface) GeoCitysClient {
	return &geoCitysClient{cc}
}

func (c *geoCitysClient) GeoCityActionCreate(ctx context.Context, in *GeoCityEntity, opts ...grpc.CallOption) (*GeoCityCreateReply, error) {
	out := new(GeoCityCreateReply)
	err := c.cc.Invoke(ctx, "/GeoCitys/GeoCityActionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCitysClient) GeoCityActionUpdate(ctx context.Context, in *GeoCityEntity, opts ...grpc.CallOption) (*GeoCityCreateReply, error) {
	out := new(GeoCityCreateReply)
	err := c.cc.Invoke(ctx, "/GeoCitys/GeoCityActionUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCitysClient) GeoCityActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoCityQueryReply, error) {
	out := new(GeoCityQueryReply)
	err := c.cc.Invoke(ctx, "/GeoCitys/GeoCityActionQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCitysClient) GeoCityActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*GeoCityReply, error) {
	out := new(GeoCityReply)
	err := c.cc.Invoke(ctx, "/GeoCitys/GeoCityActionGetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoCitysClient) GeoCityActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error) {
	out := new(workspaces.RemoveReply)
	err := c.cc.Invoke(ctx, "/GeoCitys/GeoCityActionRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type GeoCitysServer interface {
	GeoCityActionCreate(context.Context, *GeoCityEntity) (*GeoCityCreateReply, error)
	GeoCityActionUpdate(context.Context, *GeoCityEntity) (*GeoCityCreateReply, error)
	GeoCityActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoCityQueryReply, error)
	GeoCityActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoCityReply, error)
	GeoCityActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error)
	mustEmbedUnimplementedGeoCitysServer()
}

type UnimplementedGeoCitysServer struct {
}

func (UnimplementedGeoCitysServer) GeoCityActionCreate(context.Context, *GeoCityEntity) (*GeoCityCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCityActionCreate not implemented")
}
func (UnimplementedGeoCitysServer) GeoCityActionUpdate(context.Context, *GeoCityEntity) (*GeoCityCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCityActionUpdate not implemented")
}
func (UnimplementedGeoCitysServer) GeoCityActionQuery(context.Context, *workspaces.QueryFilterRequest) (*GeoCityQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCityActionQuery not implemented")
}
func (UnimplementedGeoCitysServer) GeoCityActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*GeoCityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCityActionGetOne not implemented")
}
func (UnimplementedGeoCitysServer) GeoCityActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeoCityActionRemove not implemented")
}
func (UnimplementedGeoCitysServer) mustEmbedUnimplementedGeoCitysServer() {}

type UnsafeGeoCitysServer interface {
	mustEmbedUnimplementedGeoCitysServer()
}

func RegisterGeoCitysServer(s grpc.ServiceRegistrar, srv GeoCitysServer) {
	s.RegisterService(&GeoCitys_ServiceDesc, srv)
}

func _GeoCitys_GeoCityActionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoCityEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCitysServer).GeoCityActionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoCitys/GeoCityActionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCitysServer).GeoCityActionCreate(ctx, req.(*GeoCityEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCitys_GeoCityActionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoCityEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCitysServer).GeoCityActionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoCitys/GeoCityActionUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCitysServer).GeoCityActionUpdate(ctx, req.(*GeoCityEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCitys_GeoCityActionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCitysServer).GeoCityActionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoCitys/GeoCityActionQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCitysServer).GeoCityActionQuery(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCitys_GeoCityActionGetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCitysServer).GeoCityActionGetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoCitys/GeoCityActionGetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCitysServer).GeoCityActionGetOne(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoCitys_GeoCityActionRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoCitysServer).GeoCityActionRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GeoCitys/GeoCityActionRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoCitysServer).GeoCityActionRemove(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var GeoCitys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GeoCitys",
	HandlerType: (*GeoCitysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeoCityActionCreate",
			Handler:    _GeoCitys_GeoCityActionCreate_Handler,
		},
		{
			MethodName: "GeoCityActionUpdate",
			Handler:    _GeoCitys_GeoCityActionUpdate_Handler,
		},
		{
			MethodName: "GeoCityActionQuery",
			Handler:    _GeoCitys_GeoCityActionQuery_Handler,
		},
		{
			MethodName: "GeoCityActionGetOne",
			Handler:    _GeoCitys_GeoCityActionGetOne_Handler,
		},
		{
			MethodName: "GeoCityActionRemove",
			Handler:    _GeoCitys_GeoCityActionRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/geo/GeoCityDefinitions.dyno.proto",
}
