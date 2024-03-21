package workspaces

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

type McxBundle struct {
	To   string   `json:"to" yaml:"to"`
	From []string `json:"from" yaml:"from"`
}

type MvxManifest struct {
	Bundles []McxBundle `json:"bundles" yaml:"bundles"`
	Styles  []McxBundle `json:"styles" yaml:"styles"`
}

func compileStyles(content []string, to string) error {
	// all := ""

	// for _, file := range content {
	// 	data, _ := ioutil.ReadFile(file)

	// 	var context = runtime.NewContext()
	// 	var parser = parser.NewParser(context)
	// 	var stmts = parser.ParseScss(string(data))
	// 	var compiler2 = compiler.NewCompactCompiler(context)
	// 	var out = compiler2.Com(stmts)

	// 	all += ";\r\n" + out
	// }

	// fmt.Println("To::", to)

	// os.WriteFile(to, []byte(all), 0644)

	return nil
}

func minifyJSFiles(content []string, to string) error {
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)

	all := ""

	for _, file := range content {
		data, _ := ioutil.ReadFile(file)
		minified, err := m.String("text/javascript", string(data))
		if err != nil {
			return err
		}

		// Write the minified content back to the original file
		// For example, you can use ioutil.WriteFile or another method
		all += ";\r\n" + minified
	}

	fmt.Println("To::", to)

	os.WriteFile(to, []byte(all), 0644)

	return nil
}

func CompileMvxManifest(path string) {
	var data MvxManifest

	ReadYamlFile[MvxManifest](path, &data)

	MvxRunBundles(path, &data)
}

func getaccuratevalue(manifestPath string, url string) string {
	return filepath.Join(filepath.Dir(manifestPath), url)
}

func MvxRunBundles(path string, m *MvxManifest) {

	for _, item := range m.Bundles {
		items := []string{}
		for _, mv := range item.From {
			items = append(items, getaccuratevalue(path, mv))
		}

		err := minifyJSFiles(items, getaccuratevalue(path, item.To))
		if err != nil {
			panic(err)
		}
	}

	for _, item := range m.Styles {
		items := []string{}
		for _, mv := range item.From {
			items = append(items, getaccuratevalue(path, mv))
		}

		err := compileStyles(items, getaccuratevalue(path, item.To))
		if err != nil {
			panic(err)
		}
	}
}
