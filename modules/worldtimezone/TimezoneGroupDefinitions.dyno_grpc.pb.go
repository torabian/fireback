//
//
//
//
//

package worldtimezone

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"pixelplux.com/fireback/modules/workspaces"
)

const _ = grpc.SupportPackageIsVersion7

type TimezoneGroupsClient interface {
	TimezoneGroupActionCreate(ctx context.Context, in *TimezoneGroupEntity, opts ...grpc.CallOption) (*TimezoneGroupCreateReply, error)
	TimezoneGroupActionUpdate(ctx context.Context, in *TimezoneGroupEntity, opts ...grpc.CallOption) (*TimezoneGroupCreateReply, error)
	TimezoneGroupActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*TimezoneGroupQueryReply, error)
	TimezoneGroupActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*TimezoneGroupReply, error)
	TimezoneGroupActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error)
}

type timezoneGroupsClient struct {
	cc grpc.ClientConnInterface
}

func NewTimezoneGroupsClient(cc grpc.ClientConnInterface) TimezoneGroupsClient {
	return &timezoneGroupsClient{cc}
}

func (c *timezoneGroupsClient) TimezoneGroupActionCreate(ctx context.Context, in *TimezoneGroupEntity, opts ...grpc.CallOption) (*TimezoneGroupCreateReply, error) {
	out := new(TimezoneGroupCreateReply)
	err := c.cc.Invoke(ctx, "/TimezoneGroups/TimezoneGroupActionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timezoneGroupsClient) TimezoneGroupActionUpdate(ctx context.Context, in *TimezoneGroupEntity, opts ...grpc.CallOption) (*TimezoneGroupCreateReply, error) {
	out := new(TimezoneGroupCreateReply)
	err := c.cc.Invoke(ctx, "/TimezoneGroups/TimezoneGroupActionUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timezoneGroupsClient) TimezoneGroupActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*TimezoneGroupQueryReply, error) {
	out := new(TimezoneGroupQueryReply)
	err := c.cc.Invoke(ctx, "/TimezoneGroups/TimezoneGroupActionQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timezoneGroupsClient) TimezoneGroupActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*TimezoneGroupReply, error) {
	out := new(TimezoneGroupReply)
	err := c.cc.Invoke(ctx, "/TimezoneGroups/TimezoneGroupActionGetOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timezoneGroupsClient) TimezoneGroupActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest, opts ...grpc.CallOption) (*workspaces.RemoveReply, error) {
	out := new(workspaces.RemoveReply)
	err := c.cc.Invoke(ctx, "/TimezoneGroups/TimezoneGroupActionRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type TimezoneGroupsServer interface {
	TimezoneGroupActionCreate(context.Context, *TimezoneGroupEntity) (*TimezoneGroupCreateReply, error)
	TimezoneGroupActionUpdate(context.Context, *TimezoneGroupEntity) (*TimezoneGroupCreateReply, error)
	TimezoneGroupActionQuery(context.Context, *workspaces.QueryFilterRequest) (*TimezoneGroupQueryReply, error)
	TimezoneGroupActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*TimezoneGroupReply, error)
	TimezoneGroupActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error)
	mustEmbedUnimplementedTimezoneGroupsServer()
}

type UnimplementedTimezoneGroupsServer struct {
}

func (UnimplementedTimezoneGroupsServer) TimezoneGroupActionCreate(context.Context, *TimezoneGroupEntity) (*TimezoneGroupCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimezoneGroupActionCreate not implemented")
}
func (UnimplementedTimezoneGroupsServer) TimezoneGroupActionUpdate(context.Context, *TimezoneGroupEntity) (*TimezoneGroupCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimezoneGroupActionUpdate not implemented")
}
func (UnimplementedTimezoneGroupsServer) TimezoneGroupActionQuery(context.Context, *workspaces.QueryFilterRequest) (*TimezoneGroupQueryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimezoneGroupActionQuery not implemented")
}
func (UnimplementedTimezoneGroupsServer) TimezoneGroupActionGetOne(context.Context, *workspaces.QueryFilterRequest) (*TimezoneGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimezoneGroupActionGetOne not implemented")
}
func (UnimplementedTimezoneGroupsServer) TimezoneGroupActionRemove(context.Context, *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimezoneGroupActionRemove not implemented")
}
func (UnimplementedTimezoneGroupsServer) mustEmbedUnimplementedTimezoneGroupsServer() {}

type UnsafeTimezoneGroupsServer interface {
	mustEmbedUnimplementedTimezoneGroupsServer()
}

func RegisterTimezoneGroupsServer(s grpc.ServiceRegistrar, srv TimezoneGroupsServer) {
	s.RegisterService(&TimezoneGroups_ServiceDesc, srv)
}

func _TimezoneGroups_TimezoneGroupActionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimezoneGroupEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimezoneGroupsServer).TimezoneGroupActionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TimezoneGroups/TimezoneGroupActionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimezoneGroupsServer).TimezoneGroupActionCreate(ctx, req.(*TimezoneGroupEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimezoneGroups_TimezoneGroupActionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimezoneGroupEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimezoneGroupsServer).TimezoneGroupActionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TimezoneGroups/TimezoneGroupActionUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimezoneGroupsServer).TimezoneGroupActionUpdate(ctx, req.(*TimezoneGroupEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimezoneGroups_TimezoneGroupActionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimezoneGroupsServer).TimezoneGroupActionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TimezoneGroups/TimezoneGroupActionQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimezoneGroupsServer).TimezoneGroupActionQuery(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimezoneGroups_TimezoneGroupActionGetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimezoneGroupsServer).TimezoneGroupActionGetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TimezoneGroups/TimezoneGroupActionGetOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimezoneGroupsServer).TimezoneGroupActionGetOne(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimezoneGroups_TimezoneGroupActionRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(workspaces.QueryFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimezoneGroupsServer).TimezoneGroupActionRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TimezoneGroups/TimezoneGroupActionRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimezoneGroupsServer).TimezoneGroupActionRemove(ctx, req.(*workspaces.QueryFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var TimezoneGroups_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TimezoneGroups",
	HandlerType: (*TimezoneGroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TimezoneGroupActionCreate",
			Handler:    _TimezoneGroups_TimezoneGroupActionCreate_Handler,
		},
		{
			MethodName: "TimezoneGroupActionUpdate",
			Handler:    _TimezoneGroups_TimezoneGroupActionUpdate_Handler,
		},
		{
			MethodName: "TimezoneGroupActionQuery",
			Handler:    _TimezoneGroups_TimezoneGroupActionQuery_Handler,
		},
		{
			MethodName: "TimezoneGroupActionGetOne",
			Handler:    _TimezoneGroups_TimezoneGroupActionGetOne_Handler,
		},
		{
			MethodName: "TimezoneGroupActionRemove",
			Handler:    _TimezoneGroups_TimezoneGroupActionRemove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/worldtimezone/TimezoneGroupDefinitions.dyno.proto",
}
