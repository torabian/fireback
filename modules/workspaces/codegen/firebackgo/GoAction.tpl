package {{ .m.Path }}

import (
	"github.com/gin-gonic/gin"
    "github.com/urfave/cli"
    {{ if ne .wsprefix "" }}
	"pixelplux.com/fireback/modules/workspaces"
    {{ end }}
)


{{ range .m.Actions }}

    {{ if .In.Fields }}


type {{ .Upper }}ActionReqDto struct {
    {{ template "definitionrow" (arr .In.Fields $.wsprefix) }}
}

func ( x * {{ .Upper }}ActionReqDto) RootObjectName() string {
	return "{{ $.m.Path }}"
}

var {{ .Upper }}CommonCliFlagsOptional = []cli.Flag{
  {{ template "dtoCliFlag" (arr .In.Fields "") }}
}


func {{ .Upper }}ActionReqValidator(dto *{{ .Upper }}ActionReqDto) *{{ $.wsprefix }}IError {
    err := {{ $.wsprefix }}CommonStructValidatorPointer(dto, false)

    {{ range .In.Fields }}
      {{ if  eq .Type "array"  }}
        if dto != nil && dto.{{ .UpperPlural }} != nil {
          {{ $.wsprefix }}AppendSliceErrors(dto.{{ .UpperPlural }}, false, "{{ .Name}}", err)
        }
      {{ end }}
    {{ end }}
    return err
  }


func Cast{{ .Upper }}FromCli (c *cli.Context) *{{ .Upper }}ActionReqDto {
	template := &{{ .Upper }}ActionReqDto{}

	{{ template "entityCliCastRecursive" (arr .In.Fields "")}}

	return template
}

    {{ end }}

    {{ if .Out.Fields }}

type {{ .Upper }}ActionResDto struct {
    {{ template "definitionrow" (arr .Out.Fields $.wsprefix) }}
}

func ( x * {{ .Upper }}ActionResDto) RootObjectName() string {
	return "{{ $.m.Path }}"
}


    {{ end }}

type {{ .Name }}ActionImpSig func(req {{ .ActionReqDto }}, q {{ $.wsprefix }}QueryDSL) ({{ .ActionResDto }}, *{{ $.wsprefix }}IError)
var {{ .Upper }}ActionImp {{ .Name }}ActionImpSig

func {{ .Upper }}ActionFn(req {{ .ActionReqDto }}, q {{ $.wsprefix }}QueryDSL) ({{ .ActionResDto }}, *{{ $.wsprefix }}IError) {

    if {{ .Upper }}ActionImp == nil {
        return nil, nil
    }

    return {{ .Upper }}ActionImp(req, q)
}

var {{ .Upper }}ActionCmd cli.Command = cli.Command{
	Name:  "{{ .ComputedCliName }}",
	Usage: "{{ .Description }}",
    {{ if .In.Fields }}
	Flags: {{ .Upper }}CommonCliFlagsOptional,
    {{ else if .In.Entity }}
	Flags: {{ .In.EntityPure }}CommonCliFlagsOptional,
    {{ end }}
	Action: func(c *cli.Context) {
		query := CommonCliQueryDSLBuilder(c)
        {{ if .In.Fields }}
		dto := Cast{{ .Upper }}FromCli(c)
        {{ else if .In.Entity }}
		dto := Cast{{ .In.EntityPure }}FromCli(c)
        {{ end }}
		result, err := {{ .Upper }}ActionFn(dto, query)

		HandleActionInCli(c, result, err, map[string]map[string]string{})
	},
}

{{ end }}

func {{ .m.PublicName }}CustomActions() []{{ $.wsprefix }}Module2Action {
	routes := []{{ $.wsprefix }}Module2Action{
        {{ range .m.Actions }}
		{
			Method: "{{ .MethodAllUpper }}",
			Url:    "{{ .ComputedUrl }}",
			Handlers: []gin.HandlerFunc{
                {{ if ne .SecurityModel.Model "public"}}
                WithAuthorization([]string{}),
                {{ end }}
				func(c *gin.Context) {
                    {{ $.wsprefix }}HttpPostEntity(c, {{ .Upper }}ActionFn)
                },
			},
			Format:         "{{ .FormatComputed }}",
			Action:         {{ .Upper }}ActionFn,
            {{ if .ComputeResponseEntity }}
			ResponseEntity: {{.ComputeResponseEntity}},
            {{ end }}
            {{ if .ComputeRequestEntity}}
			RequestEntity: {{.ComputeRequestEntity}},
            {{ end }}
		},
        {{ end }}
	}
    
	return routes
}

var {{ .m.Upper }}CustomActionsCli = []cli.Command {
{{ range .m.Actions }}
    {{ .Upper }}ActionCmd,
{{ end }}
}