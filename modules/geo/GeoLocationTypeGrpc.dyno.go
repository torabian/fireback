//go:build !omit_grpc

package geo

import (
	context "context"

	"pixelplux.com/fireback/modules/workspaces"
)

type GeoLocationTypeGrpcServer struct {
	UnimplementedGeoLocationTypesServer
}

func (s *GeoLocationTypeGrpcServer) GeoLocationTypeActionCreate(ctx context.Context, in *GeoLocationTypeEntity) (*GeoLocationTypeCreateReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOLOCATIONTYPE_CREATE}); authError == nil {
		data, err := GeoLocationTypeActionCreate(in, query)
		return &GeoLocationTypeCreateReply{Data: data, Error: err}, nil
	} else {
		return &GeoLocationTypeCreateReply{Data: nil, Error: authError}, nil
	}
}

func (s *GeoLocationTypeGrpcServer) GeoLocationTypeActionUpdate(ctx context.Context, in *GeoLocationTypeEntity) (*GeoLocationTypeCreateReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOLOCATIONTYPE_UPDATE}); authError == nil {

		entity, err := GeoLocationTypeActionUpdate(query, in)
		return &GeoLocationTypeCreateReply{Data: entity, Error: err}, nil
	} else {
		return &GeoLocationTypeCreateReply{Error: authError}, nil
	}
}

func (s *GeoLocationTypeGrpcServer) GeoLocationTypeActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOLOCATIONTYPE_DELETE}); authError == nil {
		affectedRows, err := GeoLocationTypeActionRemove(query)
		return &workspaces.RemoveReply{RowsAffected: affectedRows, Error: err}, nil
	} else {
		return &workspaces.RemoveReply{Error: authError}, nil
	}
}

func (s *GeoLocationTypeGrpcServer) GeoLocationTypeActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoLocationTypeQueryReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOLOCATIONTYPE_CREATE}); authError == nil {
		items, meta, err := GeoLocationTypeActionQuery(query)
		return &GeoLocationTypeQueryReply{Items: items, TotalItems: meta.TotalItems, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoLocationTypeQueryReply{Error: authError}, nil
	}
}

func (s *GeoLocationTypeGrpcServer) GeoLocationTypeActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoLocationTypeReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOLOCATIONTYPE_QUERY}); authError == nil {
		entity, err := GeoLocationTypeActionGetOne(query)
		return &GeoLocationTypeReply{Data: entity, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoLocationTypeReply{Error: authError}, nil
	}
}
