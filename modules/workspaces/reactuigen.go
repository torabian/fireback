package workspaces

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/gin-gonic/gin"
	reactui "github.com/torabian/fireback/modules/workspaces/codegen/react-ui"
)

func getTranslationKeys(entity *Module2Entity) map[string]string {
	dic := map[string]string{}
	dic["edit"+entity.Name] = "Edit " + CamelCaseToWords(entity.Name)
	dic["new"+entity.Name] = "New" + CamelCaseToWords(entity.Name)
	dic["archiveTitle"] = CamelCaseToWords(entity.Name)

	for _, field := range entity.Fields {
		dic[field.Name] = CamelCaseToWords(field.Name)
		dic[field.Name+"Hint"] = CamelCaseToWords(field.Name) + " hint"

	}

	return dic
}

func RenderReactUiTemplate(
	xapp *XWebServer,
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
	entityName string,
) ([]byte, error) {

	t, err := template.New("").Funcs(commonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	pathSplit := strings.Split(ctx.EntityPath, ".")
	modulePath := pathSplit[0 : len(pathSplit)-1]

	pluralize2 := pluralize.NewClient()
	templtes := strings.ToLower(pluralize2.Plural(entityName))

	e := FindModule2Entity(xapp, ctx.EntityPath)

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"ctx":       ctx,
		"Template":  entityName,
		"SdkDir":    "fireback",
		"ModuleDir": strings.Join(modulePath, "/"),
		"templates": templtes,
		"e":         e,
	})

	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func ReactUiCodeGen(xapp *XWebServer, ctx *CodeGenContext) error {

	os.MkdirAll(ctx.Path, os.ModePerm)
	pathSplit := strings.Split(ctx.EntityPath, ".")
	entityName := ToUpper(pathSplit[len(pathSplit)-1])
	e := FindModule2Entity(xapp, ctx.EntityPath)

	pluralize2 := pluralize.NewClient()
	templtes := strings.ToLower(pluralize2.Plural(entityName))

	jo := map[string]interface{}{}
	jo[templtes] = getTranslationKeys(e)

	u, _ := json.MarshalIndent(jo, "", "  ")
	fmt.Println(string(u))

	files, _ := GetAllFilenames(&reactui.ReactUITpl, ".")
	for _, file := range files {

		if strings.HasPrefix(file, "Template") {
			data, err := RenderReactUiTemplate(xapp, ctx, reactui.ReactUITpl, file, entityName)
			if err != nil {
				log.Fatalln(err)
			}

			newFile := strings.ReplaceAll(file, "Template", entityName)
			newFile = strings.ReplaceAll(newFile, ".tpl", "")
			exportPath := path.Join(ctx.Path, newFile+".tsx")
			os.WriteFile(exportPath, EscapeLines(data), 0644)
		}
	}

	return nil
}
