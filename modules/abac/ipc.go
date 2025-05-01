package abac

import (
	"net/url"
	"strings"

	"github.com/torabian/fireback/modules/fireback"
)

func WithIPCAuthorization(q *fireback.QueryDSL, sourceQuery string, actionPath string, ipcSecurity string) {
	if q.ItemsPerPage == 0 {
		q.ItemsPerPage = 10
	}

	var getBookPath = fireback.NewUrl(actionPath)
	match, ok := getBookPath.Match("/" + sourceQuery)
	if ok && match.Params["uniqueId"] != "" {
		q.UniqueId = match.Params["uniqueId"]
	}

	myUrl, _ := url.Parse(sourceQuery)

	params, _ := url.ParseQuery(myUrl.RawQuery)
	token := params.Get("token")
	// workspaceId := params.Get("workspaceId")
	deep := params.Get("deep")
	withPreloads := params.Get("withPreloads")
	query := params.Get("query")
	q.Authorization = token

	if deep == "true" {
		q.Deep = true
	}

	if query != "" && q.Query == "" {
		q.Query = query
	}

	if len(withPreloads) > 0 {
		q.WithPreloads = strings.Split(strings.Trim(withPreloads, " "), ",")
	}

	if ipcSecurity == "user" {
		// handle error
		user, _ := GetUserFromToken(token)
		q.UserId = user.UniqueId
	}

	if ipcSecurity == "workspace" {

		// context := &AuthContext{
		// 	WorkspaceId:  workspaceId,
		// 	Token:        token,
		// 	Capabilities: []string{},
		// }
		// result, err := WithAuthorizationPure(context)

		// data, _ := json.MarshalIndent(result, "", "  ")
		// dataer, _ := json.MarshalIndent(err, "", "  ")
		// conte, _ := json.MarshalIndent(context, "", "  ")

	}

	// if err != nil {
	// 	c.AbortWithStatusJSON(int(err.HttpCode), gin.H{"error": err})
	// 	return
	// }

}

func ActionArgumentFormatQuery(query string, affix string, ipcSecurity string) fireback.QueryDSL {

	queryParsed := fireback.DtoFromString[fireback.QueryDSL](query)
	WithIPCAuthorization(&queryParsed, query, affix, ipcSecurity)
	return queryParsed
}

func ActionArgumentsFormatPostOne[T any](dto string, query string, affix string, ipcSecurity string) (*T, fireback.QueryDSL) {
	dtoParsed := fireback.DtoFromString[T](dto)
	queryParsed := fireback.DtoFromString[fireback.QueryDSL](query)
	WithIPCAuthorization(&queryParsed, query, affix, ipcSecurity)
	return &dtoParsed, queryParsed
}

func ActionArgumentsFormatUpdateOne[T any](dto string, query string, affix string, ipcSecurity string) (fireback.QueryDSL, *T) {
	dtoParsed := fireback.DtoFromString[T](dto)
	queryParsed := fireback.DtoFromString[fireback.QueryDSL](query)
	WithIPCAuthorization(&queryParsed, query, affix, ipcSecurity)
	return queryParsed, &dtoParsed
}

func ActionArgumentsFormatDeleteDSL(query string, affix string, ipcSecurity string) fireback.QueryDSL {
	queryParsed := fireback.DtoFromString[fireback.QueryDSL](query)
	WithIPCAuthorization(&queryParsed, query, affix, ipcSecurity)
	return queryParsed
}
