//go:build !omit_grpc

package worldtimezone

import (
	context "context"

	"pixelplux.com/fireback/modules/workspaces"
)

type TimezoneGroupGrpcServer struct {
	UnimplementedTimezoneGroupsServer
}

func (s *TimezoneGroupGrpcServer) TimezoneGroupActionCreate(ctx context.Context, in *TimezoneGroupEntity) (*TimezoneGroupCreateReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_TIMEZONEGROUP_CREATE}); authError == nil {
		data, err := TimezoneGroupActionCreate(in, query)
		return &TimezoneGroupCreateReply{Data: data, Error: err}, nil
	} else {
		return &TimezoneGroupCreateReply{Data: nil, Error: authError}, nil
	}
}

func (s *TimezoneGroupGrpcServer) TimezoneGroupActionUpdate(ctx context.Context, in *TimezoneGroupEntity) (*TimezoneGroupCreateReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_TIMEZONEGROUP_UPDATE}); authError == nil {

		entity, err := TimezoneGroupActionUpdate(query, in)
		return &TimezoneGroupCreateReply{Data: entity, Error: err}, nil
	} else {
		return &TimezoneGroupCreateReply{Error: authError}, nil
	}
}

func (s *TimezoneGroupGrpcServer) TimezoneGroupActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_TIMEZONEGROUP_DELETE}); authError == nil {
		affectedRows, err := TimezoneGroupActionRemove(query)
		return &workspaces.RemoveReply{RowsAffected: affectedRows, Error: err}, nil
	} else {
		return &workspaces.RemoveReply{Error: authError}, nil
	}
}

func (s *TimezoneGroupGrpcServer) TimezoneGroupActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest) (*TimezoneGroupQueryReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_TIMEZONEGROUP_CREATE}); authError == nil {
		items, meta, err := TimezoneGroupActionQuery(query)
		return &TimezoneGroupQueryReply{Items: items, TotalItems: meta.TotalItems, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &TimezoneGroupQueryReply{Error: authError}, nil
	}
}

func (s *TimezoneGroupGrpcServer) TimezoneGroupActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest) (*TimezoneGroupReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_TIMEZONEGROUP_QUERY}); authError == nil {
		entity, err := TimezoneGroupActionGetOne(query)
		return &TimezoneGroupReply{Data: entity, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &TimezoneGroupReply{Error: authError}, nil
	}
}
