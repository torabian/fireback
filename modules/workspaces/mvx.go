package workspaces

import (
	"os"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

type McxBundle struct {
	To   string   `json:"to" yaml:"to"`
	From []string `json:"from" yaml:"from"`
}

type MvxManifest struct {
	Bundles []McxBundle `json:"bundles" yaml:"bundles"`
}

func minifyJSFiles(files []string, to string) error {
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)

	for _, file := range files {
		minified, err := m.String("text/javascript", file)
		if err != nil {
			return err
		}
		os.WriteFile(to, []byte(minified), 0644)
		// Write the minified content back to the original file
		// For example, you can use ioutil.WriteFile or another method
	}

	return nil
}
func CompileMvxManifest(path string) {
	var data MvxManifest

	ReadYamlFile[MvxManifest](path, &data)

	MvxRunBundles(&data)
}

func MvxRunBundles(m *MvxManifest) {

	for _, item := range m.Bundles {
		err := minifyJSFiles(item.From, item.To)
		if err != nil {
			panic(err)
		}
	}
}
