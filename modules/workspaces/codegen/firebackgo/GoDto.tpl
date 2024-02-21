package {{ .m.Path }}

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
)

{{ template "goimport" . }}

{{ range .children }}
type {{ .FullName }} struct {
    {{ template "definitionrow" (arr .CompleteFields $.wsprefix) }}
}
func ( x * {{ .FullName }}) RootObjectName() string {
	return "{{ $.e.DtoName }}"
}
{{ end }}

{{ template "dtoCastFromCli" . }}


var {{ .e.DtoName }}CommonCliFlagsOptional = []cli.Flag{
  {{ template "entityCommonCliFlag" (arr .e.CompleteFields "") }}
}

type {{ .e.DtoName }} struct {
    {{ template "definitionrow" (arr .e.CompleteFields $.wsprefix) }}
}

func (x* {{ .e.DtoName }}) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return (string(str))
	}
	return ""
}

func (x* {{ .e.DtoName }}) JsonPrint()  {
    fmt.Println(x.Json())
}
