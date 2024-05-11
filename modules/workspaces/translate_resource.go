package workspaces

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type TranslationResourceCatalog struct {
	EntryPoint string   `json:"entryPoint" yaml:"entryPoint"`
	Languages  []string `json:"languages" yaml:"languages"`
	FileFormat string   `json:"fileFormat" yaml:"fileFormat"`
}

type ExtendStrategy struct {
	Resource string `json:"r" yaml:"r"`
}

type TranslationResource struct {
	Extends []ExtendStrategy            `json:"extends" yaml:"extends"`
	Content map[interface{}]interface{} `json:"content" yaml:"content"`
}

func convertToTypeScript(content map[interface{}]interface{}, indent string) string {
	var result strings.Builder
	result.WriteString("{\n")

	for key, value := range content {
		result.WriteString(fmt.Sprintf("%s%s: ", indent, key))

		switch v := value.(type) {
		case map[interface{}]interface{}:
			result.WriteString(convertToTypeScript(v, indent+"  "))
		case string:
			result.WriteString(fmt.Sprintf("\"%s\"", v))
		case int:
			result.WriteString(fmt.Sprintf("%d", v))
		case float64:
			result.WriteString(fmt.Sprintf("%f", v))
		default:
			result.WriteString("null")
		}

		result.WriteString(",\n")
	}

	result.WriteString(indent[:len(indent)-2])
	result.WriteString("}")

	return result.String()
}

func addMissingKeysRecursive(english, persian map[interface{}]interface{}) {
	for key, value := range english {
		_, exists := persian[key]
		if !exists {
			persian[key] = value
			continue
		}

		if englishMap, englishIsMap := english[key].(map[interface{}]interface{}); englishIsMap {
			if persianMap, persianIsMap := persian[key].(map[interface{}]interface{}); persianIsMap {
				addMissingKeysRecursive(englishMap, persianMap)
			}
		}
	}
}
func addMissingKeys(english, persian *TranslationResource) {
	addMissingKeysRecursive(english.Content, persian.Content)
}

func fullTsTranslationObject(lang string, body string) string {
	template := `/**
* Auto generated file by fireback language & translation manager.
*/
export const ?lang = ?content;`

	template = strings.ReplaceAll(template, "?content", body)
	template = strings.ReplaceAll(template, "?lang", lang)
	return template
}

func replaceLanguageValues(input string) string {
	// Define the regular expression to match 'strings-XX' where XX represents any two characters
	re := regexp.MustCompile(`strings-[a-z]{2}`)

	// Replace matches with 'strings-en'
	output := re.ReplaceAllString(input, "strings-en")

	return output
}

func TranslateResource(ctx TranslationResourceCatalog) {

	// Since the rule is english will be always the entry point, we use it
	// we replace it
	ctx.EntryPoint = replaceLanguageValues(ctx.EntryPoint)

	var resource TranslationResource
	if err := ReadYamlFile[TranslationResource](ctx.EntryPoint, &resource); err != nil {
		log.Fatalln("Cannot determine the entry point file at", ctx.EntryPoint)
	}

	// At least there should be english by default
	if !Contains(ctx.Languages, "en") {
		ctx.Languages = append(ctx.Languages, "en")
	}

	// Detect the other translation files.
	for _, lang := range ctx.Languages {

		// Do not regenerate for English
		if lang == "en" {
			continue
		}
		dist := filepath.Join(filepath.Dir(ctx.EntryPoint), "strings-"+lang+"."+ctx.FileFormat)

		// We create that resource with the keys in English if does not exists.
		if !Exists(dist) {
			body, err := yaml.Marshal(resource)
			if err != nil {
				log.Fatal(err)
			}
			os.WriteFile(dist, body, 0644)
		} else {
			// Else, it's necessary, to kinda add missing keys in the new second language
			// from primary language
			var secondayLanguage TranslationResource
			if err := ReadYamlFile[TranslationResource](dist, &secondayLanguage); err != nil {
				log.Fatalln("Cannot determine the resource file:", dist)
			}

			addMissingKeys(&resource, &secondayLanguage)

			body, err := yaml.Marshal(secondayLanguage)
			if err != nil {
				log.Fatal(err)
			}
			os.WriteFile(dist, body, 0644)
		}

	}

	CompileTranslation(&ctx)
}

func ReadResource(ctx *TranslationResourceCatalog, lang string, cwd string) *TranslationResource {
	var resource TranslationResource

	fileName := "strings-" + lang + "." + ctx.FileFormat
	dist := filepath.Join((cwd), fileName)

	if err := ReadYamlFile[TranslationResource](dist, &resource); err != nil {
		return nil
	}

	return &resource
}

func mergeMaps2(map1, map2 map[interface{}]interface{}) map[interface{}]interface{} {
	result := make(map[interface{}]interface{})

	// Copy map1 to result
	for k, v := range map1 {
		result[k] = v
	}

	// Merge map2 into result
	for k, v := range map2 {
		if existingVal, ok := result[k]; ok {
			// If key already exists in result
			switch existingVal.(type) {
			case map[interface{}]interface{}:
				// If value is a map, recursively merge
				if vMap, vIsMap := v.(map[interface{}]interface{}); vIsMap {
					result[k] = mergeMaps2(existingVal.(map[interface{}]interface{}), vMap)
				}
			default:
				// If value is not a map, overwrite with value from map2
				result[k] = v
			}
		} else {
			// If key doesn't exist in result, add from map2
			result[k] = v
		}
	}

	return result
}

func getResourceAsInterface(ctx *TranslationResourceCatalog, cwd string, lang string) map[interface{}]interface{} {
	resource := ReadResource(ctx, lang, cwd)
	var content = resource.Content
	if len(resource.Extends) > 0 {
		for _, item := range resource.Extends {
			newDir := path.Join(cwd, item.Resource)
			content = mergeMaps2(content, getResourceAsInterface(ctx, newDir, lang))
			// newDir := path.Join(cwd, item.Resource)
			// fmt.Println("Need to read also:", item)
			// resource2 := ReadResource(ctx, lang, newDir)
			// content = mergeMaps2(content, resource2.Content)
		}
	}

	return content
}

func CompileTranslation(ctx *TranslationResourceCatalog) {

	// Compile translations
	allData := ""
	entryPoint := filepath.Dir(ctx.EntryPoint)

	for _, lang := range ctx.Languages {
		content := getResourceAsInterface(ctx, entryPoint, lang)
		result := fullTsTranslationObject(lang, convertToTypeScript(content, "  "))
		allData += result
	}

	prev := []string{}
	for _, lang := range ctx.Languages {
		if lang == "en" {
			continue
		}
		prev = append(prev, "$"+lang+":"+lang)
	}
	allData += "\r\n export const strings = {...en, " + strings.Join(prev, ",") + "};\r\n"
	dist := filepath.Join(filepath.Dir(ctx.EntryPoint), "translations.ts")
	os.WriteFile(dist, []byte(allData), 0644)
}
