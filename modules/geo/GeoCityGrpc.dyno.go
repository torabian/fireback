//go:build !omit_grpc

package geo

import (
	context "context"

	"github.com/torabian/fireback/modules/workspaces"
)

type GeoCityGrpcServer struct {
	UnimplementedGeoCitysServer
}

func (s *GeoCityGrpcServer) GeoCityActionCreate(ctx context.Context, in *GeoCityEntity) (*GeoCityCreateReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOCITY_CREATE}); authError == nil {
		data, err := GeoCityActionCreate(in, query)
		return &GeoCityCreateReply{Data: data, Error: err}, nil
	} else {
		return &GeoCityCreateReply{Data: nil, Error: authError}, nil
	}
}

func (s *GeoCityGrpcServer) GeoCityActionUpdate(ctx context.Context, in *GeoCityEntity) (*GeoCityCreateReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOCITY_UPDATE}); authError == nil {

		entity, err := GeoCityActionUpdate(query, in)
		return &GeoCityCreateReply{Data: entity, Error: err}, nil
	} else {
		return &GeoCityCreateReply{Error: authError}, nil
	}
}

func (s *GeoCityGrpcServer) GeoCityActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOCITY_DELETE}); authError == nil {
		affectedRows, err := GeoCityActionRemove(query)
		return &workspaces.RemoveReply{RowsAffected: affectedRows, Error: err}, nil
	} else {
		return &workspaces.RemoveReply{Error: authError}, nil
	}
}

func (s *GeoCityGrpcServer) GeoCityActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoCityQueryReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOCITY_CREATE}); authError == nil {
		items, meta, err := GeoCityActionQuery(query)
		return &GeoCityQueryReply{Items: items, TotalItems: meta.TotalItems, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoCityQueryReply{Error: authError}, nil
	}
}

func (s *GeoCityGrpcServer) GeoCityActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoCityReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOCITY_QUERY}); authError == nil {
		entity, err := GeoCityActionGetOne(query)
		return &GeoCityReply{Data: entity, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoCityReply{Error: authError}, nil
	}
}
