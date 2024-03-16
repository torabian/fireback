package workspaces

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	reactui "github.com/torabian/fireback/modules/workspaces/codegen/react-ui"
)

func RenderReactUiTemplate(
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
) ([]byte, error) {

	t, err := template.New("").Funcs(commonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"ctx": ctx,
	})

	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func ReactUiCodeGen(xapp *XWebServer, ctx *CodeGenContext) error {

	fmt.Println("@@", ctx.EntityPath)
	os.MkdirAll(ctx.Path, os.ModePerm)
	pathSplit := strings.Split(ctx.EntityPath, ".")
	entityName := ToUpper(pathSplit[len(pathSplit)-1])

	fmt.Println(ctx.EntityPath, pathSplit, "-", entityName)

	files, _ := GetAllFilenames(&reactui.ReactUITpl, ".")
	for _, file := range files {

		if strings.HasPrefix(file, "Template") {
			data, err := RenderReactUiTemplate(ctx, reactui.ReactUITpl, file)
			if err != nil {
				log.Fatalln(err)
			}

			newFile := strings.ReplaceAll(file, "Template", entityName)
			newFile = strings.ReplaceAll(newFile, ".tpl", "")
			exportPath := path.Join(ctx.Path, newFile+".tsx")
			os.WriteFile(exportPath, EscapeLines(data), 0644)
		}
	}
	fmt.Println("React UI Gen")

	return nil
}
