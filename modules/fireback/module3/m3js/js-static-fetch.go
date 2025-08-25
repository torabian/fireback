// For each action, we produce a meta class to hold the method, default url,
// and such details, and provide a function to mimic the call with type safety.

package m3js

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/fireback/modules/fireback/module3/mcore"
)

type fetchStaticFunctionContext struct {
	RequestClass     string
	ResponseClass    string
	QueryStringClass string

	RequestHeadersClass string

	// The variable which will be used as default url
	DefaultUrlVariable string

	// For certain types, we need to make res.json() cast in fetch, if it's returning a dto,
	// or entity, or has fields. For text, html, or others, it does not require and makes no sense,
	// therefor needs to be casted res.text() from fetch perspective
	CastToJson bool
}

// generates a static function, to developers prefer to make calls via axios
func FetchStaticHelper(fetchctx fetchStaticFunctionContext, ctx mcore.MicroGenContext) (*mcore.CodeChunkCompiled, error) {
	claims := []mcore.JsFnArgument{
		{
			Key: "fetch.init",
			Ts:  "init?: TypedRequestInit<unknown, unknown> | undefined",
			Js:  "init",
		},
		{
			Key: "fetch.qs",
			Ts:  "qs?: " + fetchctx.QueryStringClass,
			Js:  "qs",
		},
		{
			Key: "fetch.overrideUrl",
			Ts:  "overrideUrl?: string",
			Js:  "overrideUrl",
		},
		{
			Key: "fetch.generic",
			Ts:  "<unknown, unknown, unknown>",
			Js:  "",
		},
	}

	claimsRendered := mcore.ClaimRender(claims, ctx)

	const tmpl = `
	static Fetch = (
		|@fetch.init|,
		|@fetch.qs|,
		|@fetch.overrideUrl|
	) =>
		fetchx|@fetch.generic|(
			new URL((overrideUrl {{ if .fetchctx.DefaultUrlVariable -}} ?? {{ .fetchctx.DefaultUrlVariable }} {{- end}} ) + '?' + qs?.toString()),
			init
		)

		{{ if .fetchctx.CastToJson }}
			.then((res) => res.json())
		{{ else }}
			.then((res) => res.text())
		{{ end }}

	
		{{ if .fetchctx.ResponseClass }}
			.then((data) => new {{ .fetchctx.ResponseClass }} (data));
		{{ end }}

	`

	t := template.Must(template.New("axioshelper").Funcs(mcore.CommonMap).Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, mcore.H{
		"claims":   claimsRendered,
		"fetchctx": fetchctx,
	}); err != nil {
		return nil, err
	}

	templateResult := buf.String()
	for key, value := range claimsRendered {
		templateResult = strings.ReplaceAll(templateResult, fmt.Sprintf("|@%v|", key), value)
	}

	res := &mcore.CodeChunkCompiled{
		ActualScript: []byte(templateResult),
		CodeChunkDependenies: []mcore.CodeChunkDependency{
			{
				Objects:  []string{"fetchx", "TypedRequestInit"},
				Location: INTERNAL_SDK_LOCATION,
			},
		},
	}

	return res, nil
}

// On final stage of compiling, this varialble will be replaced with context
// sdk location on the disk
var INTERNAL_SDK_LOCATION string = "./sdk"
