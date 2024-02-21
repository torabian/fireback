//go:build !omit_grpc

package geo

import (
	context "context"

	"pixelplux.com/fireback/modules/workspaces"
)

type LocationDataGrpcServer struct {
	UnimplementedLocationDatasServer
}

func (s *LocationDataGrpcServer) LocationDataActionCreate(ctx context.Context, in *LocationDataEntity) (*LocationDataCreateReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_LOCATIONDATA_CREATE}); authError == nil {
		data, err := LocationDataActionCreate(in, query)
		return &LocationDataCreateReply{Data: data, Error: err}, nil
	} else {
		return &LocationDataCreateReply{Data: nil, Error: authError}, nil
	}
}

func (s *LocationDataGrpcServer) LocationDataActionUpdate(ctx context.Context, in *LocationDataEntity) (*LocationDataCreateReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_LOCATIONDATA_UPDATE}); authError == nil {

		entity, err := LocationDataActionUpdate(query, in)
		return &LocationDataCreateReply{Data: entity, Error: err}, nil
	} else {
		return &LocationDataCreateReply{Error: authError}, nil
	}
}

func (s *LocationDataGrpcServer) LocationDataActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_LOCATIONDATA_DELETE}); authError == nil {
		affectedRows, err := LocationDataActionRemove(query)
		return &workspaces.RemoveReply{RowsAffected: affectedRows, Error: err}, nil
	} else {
		return &workspaces.RemoveReply{Error: authError}, nil
	}
}

func (s *LocationDataGrpcServer) LocationDataActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest) (*LocationDataQueryReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_LOCATIONDATA_CREATE}); authError == nil {
		items, meta, err := LocationDataActionQuery(query)
		return &LocationDataQueryReply{Items: items, TotalItems: meta.TotalItems, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &LocationDataQueryReply{Error: authError}, nil
	}
}

func (s *LocationDataGrpcServer) LocationDataActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest) (*LocationDataReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_LOCATIONDATA_QUERY}); authError == nil {
		entity, err := LocationDataActionGetOne(query)
		return &LocationDataReply{Data: entity, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &LocationDataReply{Error: authError}, nil
	}
}
