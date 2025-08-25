package m3angular

// In Angular, on top of the existing pure javascript for each action,
// We need to create a service which would combine

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/torabian/fireback/modules/fireback/module3/mcore"
)

// represents each function in the service in Angular, which will be callable to make a
// action (http/socket/etc), we first compute that information
type angularServiceActionItem struct {
	FunctionSignature string

	// the function on angular http client service will be called, such this.http.post,get
	AngularHttpMethodFunction string
}

func AngularActionsClass(module *mcore.Module3, ctx mcore.MicroGenContext) (*mcore.CodeChunkCompiled, error) {

	const tmpl = `/**
* Angular service for actions
*/

@Injectable({ providedIn: 'root' })
export class {{ .className }} {

	  constructor(private http: HttpClient) {}

	{{ range .angularServiceActionItems }}
		{{ .FunctionSignature }} {
			/*
			// convert custom headers/query to Angular compatible objects
			const httpHeaders = new HttpHeaders(headers?.toObject() ?? {});
			const httpParams = new HttpParams({ fromObject: Object.fromEntries(query?.entries() ?? []) });
			return this.http.{{ .AngularHttpMethodFunction }}<CreateWorkspaceRes>(
				this.baseUrl + 'create',
				body,
				{ headers: httpHeaders, params: httpParams }
			);
			*/
		}
	{{end }}
}

`

	angularServiceActionItems := []angularServiceActionItem{}
	// compute the actions as functions to be placed inside the service
	for _, action := range module.Actions {
		angularServiceActionItems = append(angularServiceActionItems, angularServiceActionItem{
			FunctionSignature:         fmt.Sprintf("%v(body: any, options: any, overrideUrl?: string)", action.Name),
			AngularHttpMethodFunction: strings.ToLower(action.Method),
		})
	}

	t := template.Must(template.New("action").Funcs(mcore.CommonMap).Parse(tmpl))
	className := fmt.Sprintf("%vActionsService", mcore.ToUpper(module.Name))

	var buf bytes.Buffer
	if err := t.Execute(&buf, mcore.H{
		"shouldExport":              true,
		"angularServiceActionItems": angularServiceActionItems,
		"className":                 className,
	}); err != nil {
		return nil, err
	}

	res := &mcore.CodeChunkCompiled{
		ActualScript:       buf.Bytes(),
		SuggestedFileName:  className,
		SuggestedExtension: ".ts",
	}

	res.CodeChunkDependenies = append(res.CodeChunkDependenies, mcore.CodeChunkDependency{
		Objects: []string{
			"Injectable",
		},
		Location: "@angular/core",
	})

	res.CodeChunkDependenies = append(res.CodeChunkDependenies, mcore.CodeChunkDependency{
		Objects: []string{
			"HttpClient",
			"HttpHeaders",
			"HttpParams",
		},
		Location: "@angular/common/http",
	})

	res.CodeChunkDependenies = append(res.CodeChunkDependenies, mcore.CodeChunkDependency{
		Objects: []string{
			"Observable",
		},
		Location: "rxjs",
	})

	return res, nil
}
