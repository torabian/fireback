//go:build !omit_grpc

package geo

import (
	context "context"

	"github.com/torabian/fireback/modules/workspaces"
)

type GeoCountryGrpcServer struct {
	UnimplementedGeoCountrysServer
}

func (s *GeoCountryGrpcServer) GeoCountryActionCreate(ctx context.Context, in *GeoCountryEntity) (*GeoCountryCreateReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOCOUNTRY_CREATE}); authError == nil {
		data, err := GeoCountryActionCreate(in, query)
		return &GeoCountryCreateReply{Data: data, Error: err}, nil
	} else {
		return &GeoCountryCreateReply{Data: nil, Error: authError}, nil
	}
}

func (s *GeoCountryGrpcServer) GeoCountryActionUpdate(ctx context.Context, in *GeoCountryEntity) (*GeoCountryCreateReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOCOUNTRY_UPDATE}); authError == nil {

		entity, err := GeoCountryActionUpdate(query, in)
		return &GeoCountryCreateReply{Data: entity, Error: err}, nil
	} else {
		return &GeoCountryCreateReply{Error: authError}, nil
	}
}

func (s *GeoCountryGrpcServer) GeoCountryActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOCOUNTRY_DELETE}); authError == nil {
		affectedRows, err := GeoCountryActionRemove(query)
		return &workspaces.RemoveReply{RowsAffected: affectedRows, Error: err}, nil
	} else {
		return &workspaces.RemoveReply{Error: authError}, nil
	}
}

func (s *GeoCountryGrpcServer) GeoCountryActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoCountryQueryReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOCOUNTRY_CREATE}); authError == nil {
		items, meta, err := GeoCountryActionQuery(query)
		return &GeoCountryQueryReply{Items: items, TotalItems: meta.TotalItems, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoCountryQueryReply{Error: authError}, nil
	}
}

func (s *GeoCountryGrpcServer) GeoCountryActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoCountryReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOCOUNTRY_QUERY}); authError == nil {
		entity, err := GeoCountryActionGetOne(query)
		return &GeoCountryReply{Data: entity, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoCountryReply{Error: authError}, nil
	}
}
