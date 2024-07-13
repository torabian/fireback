package workspaces

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/gin-gonic/gin"
)

func getTranslationKeys(entity *Module2Entity) map[string]string {
	pluralize2 := pluralize.NewClient()

	dic := map[string]string{}
	dic["edit"+ToUpper(entity.Name)] = "Edit " + CamelCaseToWords(entity.Name)
	dic["new"+ToUpper(entity.Name)] = "New " + CamelCaseToWords(entity.Name)
	dic["archiveTitle"] = ToUpper(CamelCaseToWords(pluralize2.Plural(entity.Name)))

	for _, field := range entity.Fields {
		dic[field.Name] = ToUpper(CamelCaseToWords(field.Name))
		dic[field.Name+"Hint"] = ToUpper(CamelCaseToWords(field.Name))

	}

	return dic
}

func ReactUIParams(xapp *XWebServer, ctx *CodeGenContext, entityName string) map[string]any {

	pathSplit := strings.Split(ctx.EntityPath, ".")
	modulePath := pathSplit[0 : len(pathSplit)-1]

	pluralize2 := pluralize.NewClient()
	templtes := ToLower(pluralize2.Plural(entityName))
	template := ToLower(entityName)
	templateDashed := CamelCaseToWordsDashed(entityName)
	templatesDashed := CamelCaseToWordsDashed(templtes)

	e := FindModule2Entity(xapp, ctx.EntityPath)

	return gin.H{
		"ctx":             ctx,
		"Template":        entityName,
		"SdkDir":          "fireback",
		"ModuleDir":       strings.Join(modulePath, "/"),
		"templates":       templtes,
		"templatesDashed": templatesDashed,
		"templateDashed":  templateDashed,
		"template":        template,
		"e":               e,
	}
}

func RenderReactUiTemplate(
	xapp *XWebServer,
	ctx *CodeGenContext,
	fs embed.FS,
	fname string,
	entityName string,
) ([]byte, error) {

	t, err := template.New("").Funcs(CommonMap).ParseFS(fs, fname, "SharedSnippets.tpl")
	if err != nil {
		return []byte{}, err
	}
	var tpl bytes.Buffer

	pathSplit := strings.Split(ctx.EntityPath, ".")
	modulePath := pathSplit[0 : len(pathSplit)-1]

	pluralize2 := pluralize.NewClient()
	templtes := ToLower(pluralize2.Plural(entityName))
	template := ToLower(entityName)
	templateDashed := CamelCaseToWordsDashed(entityName)
	templatesDashed := CamelCaseToWordsDashed(templtes)

	e := FindModule2Entity(xapp, ctx.EntityPath)

	err = t.ExecuteTemplate(&tpl, fname, gin.H{
		"ctx":             ctx,
		"Template":        entityName,
		"SdkDir":          "fireback",
		"ModuleDir":       strings.Join(modulePath, "/"),
		"templates":       templtes,
		"templatesDashed": templatesDashed,
		"templateDashed":  templateDashed,
		"template":        template,
		"e":               e,
	})

	if err != nil {
		return []byte{}, err
	}

	return tpl.Bytes(), nil
}

func ReactUiCodeGen(xapp *XWebServer, ctx *CodeGenContext, refDir embed.FS) error {

	os.MkdirAll(ctx.Path, os.ModePerm)
	pathSplit := strings.Split(ctx.EntityPath, ".")
	entityName := ToUpper(pathSplit[len(pathSplit)-1])
	e := FindModule2Entity(xapp, ctx.EntityPath)

	pluralize2 := pluralize.NewClient()
	templtes := ToLower(pluralize2.Plural(entityName))

	jo := map[string]interface{}{}
	jo[templtes] = getTranslationKeys(e)

	u, _ := json.MarshalIndent(jo, "", "  ")
	fmt.Println(string(u))

	err2 := fs.WalkDir(refDir, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fmt.Println("path:", path)

		return nil
	})

	if err2 != nil {
		panic(err2)
	}

	files, _ := GetAllFilenames(&refDir, ".")
	for _, file := range files {
		fmt.Println("File:", file)
		if strings.HasPrefix(file, "Template") {
			data, err := RenderReactUiTemplate(xapp, ctx, refDir, file, entityName)
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
