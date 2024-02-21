package {{ .m.Path }}

import (
    "github.com/gin-gonic/gin"

    {{ if ne .m.Path "workspaces" }}
	"github.com/torabian/fireback/modules/workspaces"
	{{ end }}

	"log"
	"os"
	"fmt"
	"encoding/json"
	"strings"
	"github.com/schollz/progressbar/v3"
	"github.com/gookit/event"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	jsoniter "github.com/json-iterator/go"

    {{ if or (.e.Cte) (.e.Queries) }}
    queries "github.com/torabian/fireback/modules/{{ .m.Path }}/queries"
    {{ end }}

	"embed"
	reflect "reflect"

	"github.com/urfave/cli"
	{{ if .hasSeeders }}
	seeders "github.com/torabian/fireback/modules/{{ .m.Path }}/seeders/{{ .e.Upper }}"
	{{ end }}
	{{ if .hasMocks }}
	mocks "github.com/torabian/fireback/modules/{{ .m.Path }}/mocks/{{ .e.Upper }}"
	{{ end }}
	{{ if .hasMetas }}
	metas "github.com/torabian/fireback/modules/{{ .m.Path }}/metas"
	{{ end }}
   
)


{{ template "goimport" . }}

{{ range .children }}
type {{ .FullName }} struct {
    {{ template "defaultgofields" . }}
    {{ template "definitionrow" (arr .CompleteFields $.wsprefix) }}

	{{ if .LinkedTo }}
	LinkedTo *{{ .LinkedTo }} `yaml:"-" gorm:"-" json:"-" sql:"-"`
	{{ end }}
}

func ( x * {{ .FullName }}) RootObjectName() string {
	return "{{ $.e.EntityName }}"
}

{{ end }}

type {{ .e.EntityName }} struct {
    {{ template "defaultgofields" .e }}
    {{ template "definitionrow" (arr .e.CompleteFields $.wsprefix) }}

    {{ if .e.HasTranslations }}
    Translations     []*{{ .e.PolyglotName}} `json:"translations,omitempty" gorm:"foreignKey:LinkerId;references:UniqueId"`
    {{ end }}

    Children []*{{ .e.EntityName }} `gorm:"-" sql:"-" json:"children,omitempty" yaml:"children"`

    LinkedTo *{{ .e.EntityName }} `yaml:"-" gorm:"-" json:"-" sql:"-"`
}

var {{ .e.Upper }}PreloadRelations []string = []string{}


{{ template "eventsAndMeta" . }}

{{ template "polyglottable" . }}

{{ template "entitychildactions" . }}

{{ template "entityformatting" . }}

{{ template "mockingentity" . }}

{{ template "getEntityTranslateFields" . }}

{{ template "entitySeederInit" . }}

{{ template "entityAssociationCreate" . }}

{{ template "entityRelationContentCreation" . }}

{{ template "relationContentUpdate" . }}

{{ template "polyglot" . }}

{{ template "entityValidator" . }}

{{ template "entitySanitize" . }}

{{ template "entityBeforeCreateActions" . }}

{{ template "batchActionCreate" . }}

{{ template "entityActionCreate" . }}

{{ template "entityActionGetAndQuery" . }}

{{ template "queriesAndPivot" . }}

{{ template "entityUpdateExec" . }}

{{ template "entityUpdateAction" . }}

{{ template "entityRemoveAndCleanActions" . }}

{{ template "entityBulkUpdate" . }}

{{ template "entityExtensions" . }}

{{ template "entityMeta" . }}

{{ template "entityImportExport" . }}

{{ template "cliFlags" . }}

{{ template "entityCliCommands" . }}

{{ template "entityCastFromCli" . }}

{{ template "entityMockAndSeeders" . }}

{{ template "entityCliImportExportCmd" . }}

{{ template "entityCliActionsCmd" . }}

{{ template "entityHttp" . }}

{{ template "entityPermissions" . }}

{{ template "recursiveGetEnums" (arr .e.CompleteFields .e.Upper)}}

{{ template "entityDistinctOperations" . }}