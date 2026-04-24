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
	List           []string `json:"list,omitempty"`
	Query          string   `json:"query,omitempty"`
	Suspense       bool     `json:"suspense,omitempty"`
	ForceImmediate bool     `json:"forceImmediate"`
}

type DeleteResponseDataItem struct {
	// Determines if the deletion is already executed immediately.
	Executed bool `json:"executed,omitempty"`

	// Task id
	TaskId int64 `json:"taskId,omitempty"`

	// Rows affected upon the execution, has more than zero if anything has removed.
	RowsAffected int64 `json:"rowsAffected,omitempty"`
}

type DeleteResponseData struct {
	Item DeleteResponseDataItem `json:"item,omitempty"`
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
