package fireback

import (
	"encoding/json"
)

// The base class definition for deleteResponseDto
type DeleteResponseDto struct {
	Data DeleteResponseDtoData `json:"data" yaml:"data"`
}

// The base class definition for data
type DeleteResponseDtoData struct {
	Item DeleteResponseDtoDataItem `json:"item" yaml:"item"`
}

// The base class definition for item
type DeleteResponseDtoDataItem struct {
	// If the deletion executed immediately.
	Executed bool `json:"executed" yaml:"executed"`
	// The query selector which would be used to delete the content.
	RowsAffected int64 `json:"rowsAffected" yaml:"rowsAffected"`
}

func (x *DeleteResponseDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
