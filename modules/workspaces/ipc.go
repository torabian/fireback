package workspaces

import (
	"net/url"
	"strings"
)

func WithIPCAuthorization(q *QueryDSL, sourceQuery string, actionPath string, ipcSecurity string) {
	if q.ItemsPerPage == 0 {
		q.ItemsPerPage = 10
	}

	var getBookPath = NewUrl(actionPath)
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

func ActionArgumentFormatQuery(query string, affix string, ipcSecurity string) QueryDSL {

	queryParsed := DtoFromString[QueryDSL](query)
	WithIPCAuthorization(&queryParsed, query, affix, ipcSecurity)
	return queryParsed
}

func ActionArgumentsFormatPostOne[T any](dto string, query string, affix string, ipcSecurity string) (*T, QueryDSL) {
	dtoParsed := DtoFromString[T](dto)
	queryParsed := DtoFromString[QueryDSL](query)
	WithIPCAuthorization(&queryParsed, query, affix, ipcSecurity)
	return &dtoParsed, queryParsed
}

func ActionArgumentsFormatUpdateOne[T any](dto string, query string, affix string, ipcSecurity string) (QueryDSL, *T) {
	dtoParsed := DtoFromString[T](dto)
	queryParsed := DtoFromString[QueryDSL](query)
	WithIPCAuthorization(&queryParsed, query, affix, ipcSecurity)
	return queryParsed, &dtoParsed
}

func ActionArgumentsFormatDeleteDSL(query string, affix string, ipcSecurity string) QueryDSL {
	queryParsed := DtoFromString[QueryDSL](query)
	WithIPCAuthorization(&queryParsed, query, affix, ipcSecurity)
	return queryParsed
}

/*
	IPC has never been used, but we can create them like:
	items := []Module2Action{}
	items = append(items, academy.GetAcSectionModule2Actions()...)
	items = append(items, academy.GetExamSessionModule2Actions()...)
	items = append(items, academy.GetExamModule2Actions()...)
	items = append(items, academy.GetExamSessionReviewModule2Actions()...)
	items = append(items, academy.GetAcTaskModule2Actions()...)
	items = append(items, workspaces.GetPassportModule2Actions()...)
	items = append(items, GetWorkspaceModule2Actions()...)

	xapp.CliActions = func() []cli.Command {
		return []cli.Command{
			CLIInit,
			CLIAboutCommand,
			CLIDoctor,
			CLIServiceCommand,
			GenerateBindings(items),
		**/
