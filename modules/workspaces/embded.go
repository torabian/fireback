package workspaces

import (
	"embed"
	"path"
	"strings"
)

func GetAllFilenames(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			res, err := GetAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}

	return
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

/**
*	Seeders are files such as yml, yaml, json and csv. We might not want to show
*	All files
**/
func GetSeederFilenames(fs *embed.FS, dir string) (out []string, err error) {
	out, err = GetAllFilenames(fs, dir)

	out = Filter(out, func(s string) bool {
		return strings.HasSuffix(s, ".yml") || strings.HasSuffix(s, ".yaml") || strings.HasSuffix(s, ".json") || strings.HasSuffix(s, ".csv")
	})
	return
}
