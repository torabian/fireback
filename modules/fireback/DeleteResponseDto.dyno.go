package fireback

import "encoding/json"
import emigo "github.com/torabian/emi/emigo"

func GetDeleteResponseDtoCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "data",
			Type:     "object",
			Children: GetDeleteResponseDtoDataCliFlags("data-"),
		},
	}
}
func CastDeleteResponseDtoFromCli(c emigo.CliCastable) DeleteResponseDto {
	data := DeleteResponseDto{}
	if c.IsSet("data") {
		data.Data = CastDeleteResponseDtoDataFromCli(c)
	}
	return data
}

// The base class definition for deleteResponseDto
type DeleteResponseDto struct {
	Data DeleteResponseDtoData `json:"data" yaml:"data"`
}

func GetDeleteResponseDtoDataCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "item",
			Type:     "object",
			Children: GetDeleteResponseDtoDataItemCliFlags("item-"),
		},
	}
}
func CastDeleteResponseDtoDataFromCli(c emigo.CliCastable) DeleteResponseDtoData {
	data := DeleteResponseDtoData{}
	if c.IsSet("item") {
		data.Item = CastDeleteResponseDtoDataItemFromCli(c)
	}
	return data
}

// The base class definition for data
type DeleteResponseDtoData struct {
	Item DeleteResponseDtoDataItem `json:"item" yaml:"item"`
}

func GetDeleteResponseDtoDataItemCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "executed",
			Type: "bool",
		},
		{
			Name: prefix + "rows-affected",
			Type: "int64",
		},
	}
}
func CastDeleteResponseDtoDataItemFromCli(c emigo.CliCastable) DeleteResponseDtoDataItem {
	data := DeleteResponseDtoDataItem{}
	if c.IsSet("executed") {
		data.Executed = bool(c.Bool("executed"))
	}
	if c.IsSet("rows-affected") {
		data.RowsAffected = int64(c.Int64("rows-affected"))
	}
	return data
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
