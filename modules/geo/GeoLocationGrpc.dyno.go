//go:build !omit_grpc

package geo

import (
	context "context"

	"github.com/torabian/fireback/modules/workspaces"
)

type GeoLocationGrpcServer struct {
	UnimplementedGeoLocationsServer
}

func (s *GeoLocationGrpcServer) GeoLocationActionCreate(ctx context.Context, in *GeoLocationEntity) (*GeoLocationCreateReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOLOCATION_CREATE}); authError == nil {
		data, err := GeoLocationActionCreate(in, query)
		return &GeoLocationCreateReply{Data: data, Error: err}, nil
	} else {
		return &GeoLocationCreateReply{Data: nil, Error: authError}, nil
	}
}

func (s *GeoLocationGrpcServer) GeoLocationActionUpdate(ctx context.Context, in *GeoLocationEntity) (*GeoLocationCreateReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOLOCATION_UPDATE}); authError == nil {

		entity, err := GeoLocationActionUpdate(query, in)
		return &GeoLocationCreateReply{Data: entity, Error: err}, nil
	} else {
		return &GeoLocationCreateReply{Error: authError}, nil
	}
}

func (s *GeoLocationGrpcServer) GeoLocationActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOLOCATION_DELETE}); authError == nil {
		affectedRows, err := GeoLocationActionRemove(query)
		return &workspaces.RemoveReply{RowsAffected: affectedRows, Error: err}, nil
	} else {
		return &workspaces.RemoveReply{Error: authError}, nil
	}
}

func (s *GeoLocationGrpcServer) GeoLocationActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoLocationQueryReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOLOCATION_CREATE}); authError == nil {
		items, meta, err := GeoLocationActionQuery(query)
		return &GeoLocationQueryReply{Items: items, TotalItems: meta.TotalItems, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoLocationQueryReply{Error: authError}, nil
	}
}

func (s *GeoLocationGrpcServer) GeoLocationActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoLocationReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOLOCATION_QUERY}); authError == nil {
		entity, err := GeoLocationActionGetOne(query)
		return &GeoLocationReply{Data: entity, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoLocationReply{Error: authError}, nil
	}
}
