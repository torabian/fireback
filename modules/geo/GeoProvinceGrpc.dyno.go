//go:build !omit_grpc

package geo

import (
	context "context"

	"pixelplux.com/fireback/modules/workspaces"
)

type GeoProvinceGrpcServer struct {
	UnimplementedGeoProvincesServer
}

func (s *GeoProvinceGrpcServer) GeoProvinceActionCreate(ctx context.Context, in *GeoProvinceEntity) (*GeoProvinceCreateReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOPROVINCE_CREATE}); authError == nil {
		data, err := GeoProvinceActionCreate(in, query)
		return &GeoProvinceCreateReply{Data: data, Error: err}, nil
	} else {
		return &GeoProvinceCreateReply{Data: nil, Error: authError}, nil
	}
}

func (s *GeoProvinceGrpcServer) GeoProvinceActionUpdate(ctx context.Context, in *GeoProvinceEntity) (*GeoProvinceCreateReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOPROVINCE_UPDATE}); authError == nil {

		entity, err := GeoProvinceActionUpdate(query, in)
		return &GeoProvinceCreateReply{Data: entity, Error: err}, nil
	} else {
		return &GeoProvinceCreateReply{Error: authError}, nil
	}
}

func (s *GeoProvinceGrpcServer) GeoProvinceActionRemove(ctx context.Context, in *workspaces.QueryFilterRequest) (*workspaces.RemoveReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOPROVINCE_DELETE}); authError == nil {
		affectedRows, err := GeoProvinceActionRemove(query)
		return &workspaces.RemoveReply{RowsAffected: affectedRows, Error: err}, nil
	} else {
		return &workspaces.RemoveReply{Error: authError}, nil
	}
}

func (s *GeoProvinceGrpcServer) GeoProvinceActionQuery(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoProvinceQueryReply, error) {
	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOPROVINCE_CREATE}); authError == nil {
		items, meta, err := GeoProvinceActionQuery(query)
		return &GeoProvinceQueryReply{Items: items, TotalItems: meta.TotalItems, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoProvinceQueryReply{Error: authError}, nil
	}
}

func (s *GeoProvinceGrpcServer) GeoProvinceActionGetOne(ctx context.Context, in *workspaces.QueryFilterRequest) (*GeoProvinceReply, error) {

	if query, authError := workspaces.GrpcWithAuthorization(&ctx, []string{PERM_ROOT_GEOPROVINCE_QUERY}); authError == nil {
		entity, err := GeoProvinceActionGetOne(query)
		return &GeoProvinceReply{Data: entity, Error: workspaces.GormErrorToIError(err)}, nil
	} else {
		return &GeoProvinceReply{Error: authError}, nil
	}
}
