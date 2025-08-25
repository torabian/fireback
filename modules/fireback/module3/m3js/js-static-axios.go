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

// generates a static function, to developers prefer to make calls via axios
// axios context does not exists, uses the fetch native data
func AxiosStaticHelper(fetchctx fetchStaticFunctionContext, ctx mcore.MicroGenContext) (*mcore.CodeChunkCompiled, error) {
	claims := []mcore.JsFnArgument{
		{
			Key: "axios.clientInstance",
			Ts:  "clientInstance: AxiosInstance",
			Js:  "clientInstance",
		},
		{
			Key: "axios.config",
			Ts:  "config: AxiosRequestConfig<unknown>",
			Js:  "config",
		},
		{
			Key: "axios.request.generic",
			Js:  "",
			Ts:  "<unknown, AxiosResponse<unknown>, unknown>",
		},
	}

	claimsRendered := mcore.ClaimRender(claims, ctx)

	const tmpl = `
  	static Axios = (|@axios.clientInstance|, |@axios.config|) =>
		clientInstance
		.request|@axios.request.generic|(config)

		{{ if and .fetchctx.CastToJson .fetchctx.ResponseClass }}
		.then((res) => {
			return {
			...res,

			
			// if there is a output class, create instance out of it.
			data: new {{ .fetchctx.ResponseClass }}(res.data),
			};
		});
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
				Objects:  []string{"AxiosInstance", "AxiosRequestConfig", "AxiosResponse"},
				Location: "axios",
			},
		},
	}

	return res, nil
}
