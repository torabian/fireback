package fireback

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type QueryDSL struct {
	// It's a common string query againt database in text format.
	Query string `json:"query"`

	// Usefull for the paginated queries, it would add the start index
	// and in SQL becomes as offset
	StartIndex int `json:"startIndex"`

	// Numeric cursor
	Cursor *string `json:"cursor"`

	// Useful for paginated queries, similar to limit in SQL queries
	ItemsPerPage int `json:"itemsPerPage"`

	// It's gonna make left joins on the query, if the entity has
	// objects or arrays. It would slow down the query dramatically.
	Deep bool `json:"deep"`

	// It would indicate how the result to be sorted on SQL queries
	Sort string `json:"sort"`

	// Useful when querying against a single element, and by passing
	// uniqueId you can retrieve the record
	UniqueId string `json:"uniqueId"`

	// Sometimes you need to access the raw socket connection, here there is :)
	RawSocketConnection *websocket.Conn

	// It would indicate to the Gorm orm which tables to be included in the
	// SQL search.
	WithPreloads []string `json:"withPreloads"`

	// A modern JSON Query object to replace the 'query' section.
	// JsonQuery can allow for better and more complex queries inspired by
	// Elastic search json query format.
	JsonQuery string `json:"jsonQuery"`

	// this is gin context upon the request, which is being attached to the dsl
	// regularly, should not be accessed directly but in reality many times we need
	// to work low level and there is no reason framework do not allow it.
	c *gin.Context `json:"-" yaml:"-"`

	// The gorm transaction object. By setting the query Tx, you can connect
	// few Fireback actions to be done as transaction. Fireback also uses this
	// Object between it's own operations
	Tx *gorm.DB

	// This event will be trigged in the system, if that action is done
	TriggerEventName string `json:"-"`

	// The header authorization, the encrypted token is availble
	// in every request
	Authorization string `json:"authorization"`

	// Parsed languages
	AcceptLanguage []LangQ `json:"-"`

	// Automatically assigned UserId to the request after analising the token
	// This will be used to save each entity and determine the owner of the record
	UserId string `json:"-"`

	ResolveStrategy string `json:"-"`

	LinkerId string `json:"-"`

	// This is the person who is requesting, regardless of the workspace
	SearchPhrase string `json:"searchPhrase"`

	// This is the workspace which user is working inside, usually data belongs there
	WorkspaceId string `json:"-"`

	// Those capabilities which user has
	ActionRequires []PermissionInfo `json:"-"`

	// List of permissions that this request is affecting
	RequestAffectingScopes []string `json:"-"`

	// This is the capabilities that user has
	UserHas []string `json:"-"`

	UserAccessPerWorkspace *UserAccessPerWorkspaceDto `json:"-" yaml:"-"`

	// This is limitation of that workspace
	WorkspaceHas []string `json:"-"`

	InternalQuery string   `json:"-"`
	Language      string   `json:"-"`
	Region        string   `json:"-"`
	Preloads      []string `json:"-"`
}

func (x QueryDSL) Json() string {
	str, _ := json.MarshalIndent(x, "", "  ")
	return (string(str))

}

func (x QueryDSL) GetLanguage() string {
	return x.Language
}
