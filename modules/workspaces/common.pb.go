package workspaces

type Timestamp struct {
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

type QueryFilter struct {
	Query          string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	StartIndex     int64  `protobuf:"varint,2,opt,name=startIndex,proto3" json:"startIndex,omitempty"`
	ItemsPerPage   int64  `protobuf:"varint,3,opt,name=itemsPerPage,proto3" json:"itemsPerPage,omitempty"`
	Id             string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	AcceptLanguage string `protobuf:"bytes,5,opt,name=acceptLanguage,proto3" json:"acceptLanguage,omitempty"`
	UniqueId       string `protobuf:"bytes,6,opt,name=uniqueId,proto3" json:"uniqueId,omitempty"`
}

type RemoveRequestData struct {
	RowsAffected int64 `protobuf:"varint,1,opt,name=rowsAffected,proto3" json:"rowsAffected,omitempty"`
}

type QueryFilterRequest struct {
	Query *QueryFilter `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

type DeleteRequest struct {
	List     []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Query    string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Suspense bool     `protobuf:"varint,3,opt,name=suspense,proto3" json:"suspense,omitempty"`
}

type DeleteResponseData struct {
	RowsAffected int64 `json:"rowsAffected,omitempty"`
}

type DeleteResponse struct {
	Data *DeleteResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

type IPublicErrorItem struct {
	Location          string `json:"location,omitempty"`
	Message           string `json:"message,omitempty"`
	MessageTranslated string `json:"messageTranslated,omitempty"`
	ErrorParam        string `json:"errorParam,omitempty"`
	Type              string `json:"type,omitempty"`
}

type IErrorItem struct {
	Location   string     `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Message    *ErrorItem `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorParam string     `protobuf:"bytes,5,opt,name=errorParam,proto3" json:"errorParam,omitempty"`
	Type       string     `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

type EmptyRequest struct {
}

type OkayResponseData struct {
}

type OkayResponse struct {
	Data *OkayResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

// This is what we show to the public
type IPublicError struct {
	Message           string              `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	MessageTranslated string              `protobuf:"bytes,4,opt,name=messageTranslated,proto3" json:"messageTranslated,omitempty"`
	Errors            []*IPublicErrorItem `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	HttpCode          int32               `protobuf:"varint,3,opt,name=httpCode,proto3" json:"httpCode,omitempty"`
}

type IError struct {
	Message           ErrorItem     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	MessageTranslated string        `protobuf:"bytes,4,opt,name=messageTranslated,proto3" json:"messageTranslated,omitempty"`
	Errors            []*IErrorItem `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	HttpCode          int32         `protobuf:"varint,3,opt,name=httpCode,proto3" json:"httpCode,omitempty"`
}

type RemoveReply struct {
	RowsAffected int64   `protobuf:"varint,1,opt,name=rowsAffected,proto3" json:"rowsAffected,omitempty"`
	Error        *IError `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}
