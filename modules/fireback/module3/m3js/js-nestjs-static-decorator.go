// For each action, we produce a meta class to hold the method, default url,
// and such details, and provide a function to mimic the call with type safety.

package m3js

import (
	"bytes"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/module3/mcore"
)

type nestJsStaticFunctionUseCase int

const (
	RequestBody nestJsStaticFunctionUseCase = iota
	ResponseBody
	RequestHeaders
	ResponseHeaders
	QueryString
	QueryParams
)

type NestJsStaticDecoratorContext struct {

	// The class which will be created out of the request.
	ClassInstance string

	// represents location of the static function, to change the request section
	// which will be created
	NestJsStaticFunctionUseCase nestJsStaticFunctionUseCase
}

// If we add a static decorator to some classes, we can used them directly in nest.js
// decorators, and req, res, headers, query strings will become typesafe automatically
func JsNestJsStaticDecorator(ctxstatic NestJsStaticDecoratorContext, ctx mcore.MicroGenContext) (*mcore.CodeChunkCompiled, error) {

	// How to do it iterte and call Compile?

	const tmpl = `/**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@{{.className}}.Nest() headers: {{.className}}): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new {{ .ctx.ClassInstance }}( {{ .instanceArguments }} );
	},
  );

`

	// For different use cases, the argument might be different
	instanceArguments := "null"
	if ctxstatic.NestJsStaticFunctionUseCase == RequestHeaders {
		instanceArguments = `Object.entries(request.headers)`
	}

	if ctxstatic.NestJsStaticFunctionUseCase == RequestBody {
		instanceArguments = `request.body`
	}

	t := template.Must(template.New("qsclass").Funcs(mcore.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, gin.H{
		"className":         ctxstatic.ClassInstance,
		"instanceArguments": instanceArguments,
		"ctx":               ctxstatic,
	}); err != nil {
		return nil, err
	}

	res := &mcore.CodeChunkCompiled{
		ActualScript: []byte(buf.Bytes()),
	}

	res.CodeChunkDependenies = append(res.CodeChunkDependenies, mcore.CodeChunkDependency{
		Objects: []string{
			"createParamDecorator", "ExecutionContext",
		},
		Location: "@nestjs/common",
	})

	return res, nil
}
