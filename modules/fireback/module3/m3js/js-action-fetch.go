// For each action, we produce a meta class to hold the method, default url,
// and such details, and provide a function to mimic the call with type safety.

package m3js

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/module3/mcore"
)

func findTokenByName(realms []mcore.GeneratedScriptToken, name string) *mcore.GeneratedScriptToken {

	for _, item := range realms {
		if item.Name == name {
			return &item
		}
	}

	return nil
}

func JsActionFetchAndMetaData(action *mcore.Module3Action, realms jsActionRealms, ctx mcore.MicroGenContext) (*mcore.CodeChunkCompiled, error) {
	className := fmt.Sprintf("Fetch%vAction", mcore.ToUpper(action.Name))

	res := &mcore.CodeChunkCompiled{}
	// How to do it iterte and call Compile?

	const tmpl = `/**
 * {{.className}}
 */

export class {{ .className }} {
  static URL = '{{ .action.Url }}';
  static Method = '{{ .action.Method }}';

  {{ .axiosStaticFunction }}
  
  {{ .fetchStaticFunction }}
}
`

	fetchctx := fetchStaticFunctionContext{
		DefaultUrlVariable: fmt.Sprintf("%v.URL", className),
	}

	if realms.ResponseClass != nil {
		responseClassToken := findTokenByName(realms.ResponseClass.Tokens, TOKEN_ROOT_CLASS)
		if responseClassToken != nil {

			// Not sure about this yet. Primitives also can be a class, right?
			// therefor they might not need to cast to json, but still you need to create a class out of them.
			fetchctx.CastToJson = true
			fetchctx.ResponseClass = responseClassToken.Value
		}
	}

	if realms.RequestHeadersClass != nil {
		requestHeadersClassToken := findTokenByName(realms.RequestHeadersClass.Tokens, TOKEN_ROOT_CLASS)
		if requestHeadersClassToken != nil {
			fetchctx.RequestHeadersClass = requestHeadersClassToken.Value
		}
	}

	if realms.QueryStringClass != nil {
		qsClassToken := findTokenByName(realms.QueryStringClass.Tokens, TOKEN_ROOT_CLASS)
		if qsClassToken != nil {
			fetchctx.QueryStringClass = qsClassToken.Value
		}
	}

	// Add the axios helper function to the response depencencies,
	axiosStaticFunction, err := AxiosStaticHelper(fetchctx, ctx)
	if err != nil {
		return nil, err
	}
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, axiosStaticFunction.CodeChunkDependenies...)

	// add the native fetch function to the axios
	fetchStaticFunction, err := FetchStaticHelper(fetchctx, ctx)
	if err != nil {
		return nil, err
	}
	res.CodeChunkDependenies = append(res.CodeChunkDependenies, fetchStaticFunction.CodeChunkDependenies...)

	t := template.Must(template.New("qsclass").Funcs(mcore.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, gin.H{
		"action":              action,
		"axiosStaticFunction": string(axiosStaticFunction.ActualScript),
		"fetchStaticFunction": string(fetchStaticFunction.ActualScript),
		"shouldExport":        true,
		"className":           className,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = []byte(buf.Bytes())

	return res, nil
}
