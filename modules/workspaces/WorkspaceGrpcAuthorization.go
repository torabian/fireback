package workspaces

import (
	context "context"
	"strconv"
	"strings"

	"google.golang.org/grpc/metadata"
)

func GetKey(ctx *context.Context, field string) string {
	data, _ := metadata.FromIncomingContext(*ctx)

	valueList := []string{}
	valueList = data.Get(field)
	queryString := ""
	if len(valueList) > 0 {
		queryString = valueList[0]
	}

	return queryString
}

func ExtractQueryDslFromGrpcContext(ctx *context.Context, authResult *AuthResultDto) QueryDSL {

	workspaceId := authResult.WorkspaceId
	internal_sql := authResult.InternalSql

	queryString := GetKey(ctx, "query")

	o := GetKey(ctx, "startIndex")
	startIndex, _ := strconv.Atoi(o)

	l := GetKey(ctx, "itemsPerPage")

	itemsPerPage, _ := strconv.Atoi(l)

	if startIndex < 0 {
		startIndex = 0
	}

	switch {
	case itemsPerPage > 100:
		itemsPerPage = 100
	case itemsPerPage <= 0:
		itemsPerPage = 20
	}

	f := QueryDSL{
		Query:         queryString,
		StartIndex:    startIndex,
		ItemsPerPage:  itemsPerPage,
		InternalQuery: *internal_sql,
		UserId:        *authResult.UserId,
		WorkspaceId:   *workspaceId,
		Language:      "en",
		Region:        "us",
		UniqueId:      GetKey(ctx, "uniqueId"),
	}

	acceptLang := GetKey(ctx, "accept-language")
	if acceptLang != "" && len(acceptLang) == 2 {
		f.Language = strings.ToLower(acceptLang)
	}

	return f
}
