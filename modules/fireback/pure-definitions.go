package fireback

type Timestamp struct {
	Seconds int64 `json:"seconds,omitempty"`
	Nanos   int32 `json:"nanos,omitempty"`
}

type QueryFilter struct {
	Query          string `json:"query,omitempty"`
	StartIndex     int64  `json:"startIndex,omitempty"`
	ItemsPerPage   int64  `json:"itemsPerPage,omitempty"`
	Id             string `json:"id,omitempty"`
	AcceptLanguage string `json:"acceptLanguage,omitempty"`
	UniqueId       string `json:"uniqueId,omitempty"`
}

type RemoveRequestData struct {
	RowsAffected int64 `json:"rowsAffected,omitempty"`
}

type QueryFilterRequest struct {
	Query *QueryFilter `json:"query,omitempty"`
}

type DeleteRequest struct {
	List     []string `json:"list,omitempty"`
	Query    string   `json:"query,omitempty"`
	Suspense bool     `json:"suspense,omitempty"`
}

type DeleteResponseData struct {
	RowsAffected int64 `json:"rowsAffected,omitempty"`
}

type DeleteResponse struct {
	Data *DeleteResponseData `json:"data,omitempty"`
}

type EmptyRequest struct {
}

type OkayResponseData struct {
}

type OkayResponse struct {
	Data *OkayResponseData `json:"data,omitempty"`
}

type RemoveReply struct {
	RowsAffected int64   `json:"rowsAffected,omitempty"`
	Error        *IError `json:"error,omitempty"`
}
