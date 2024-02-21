//go:build !omit_grpc

package geo

import (
	context "context"

	"github.com/torabian/fireback/modules/workspaces"
)

type GeoStateGrpcServer struct {
	UnimplementedGeoStatesServer
}

func (s *GeoStateGrpcServer) GeoStateActionCreate(ctx context.Context, in *GeoStateEntity) (*GeoStateCreateReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOSTATE_CREATE}); authError == nil {
		data, err := GeoStateActionCreate(in, query)
		return &GeoStateCreateReply{Data: data, Error: err}, nil
	} else {
		return &GeoStateCreateReply{Data: nil, Error: authError}, nil
	}
}

func (s *GeoStateGrpcServer) GeoStateActionUpdate(ctx context.Context, in *GeoStateEntity) (*GeoStateCreateReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOSTATE_UPDATE}); authError == nil {

		entity, err := GeoStateActionUpdate(query, in)
		return &GeoStateCreateReply{Data: entity, Error: err}, nil
	} else {
		return &GeoStateCreateReply{Error: authError}, nil
	}
}

func (s *GeoStateGrpcServer) GeoStateActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOSTATE_DELETE}); authError == nil {
		affectedRows, err := GeoStateActionRemove(query)
		return &workspaces.RemoveReply{RowsAffected: affectedRows, Error: err}, nil
	} else {
		return &workspaces.RemoveReply{Error: authError}, nil
	}
}

func (s *GeoStateGrpcServer) GeoStateActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoStateQueryReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOSTATE_CREATE}); authError == nil {
		items, meta, err := GeoStateActionQuery(query)
		return &GeoStateQueryReply{Items: items, TotalItems: meta.TotalItems, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoStateQueryReply{Error: authError}, nil
	}
}

func (s *GeoStateGrpcServer) GeoStateActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoStateReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOSTATE_QUERY}); authError == nil {
		entity, err := GeoStateActionGetOne(query)
		return &GeoStateReply{Data: entity, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoStateReply{Error: authError}, nil
	}
}
