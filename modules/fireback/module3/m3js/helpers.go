package m3js

import (
	"fmt"
	"sort"
	"strings"

	"github.com/torabian/fireback/modules/fireback/module3/mcore"
)

// In this file we are going to put small functions which can generate code for different languages

func CombineImportsJsWorld(chunk mcore.CodeChunkCompiled) string {
	// group by location
	locMap := map[string]map[string]struct{}{}

	for _, dep := range chunk.CodeChunkDependenies {
		if _, ok := locMap[dep.Location]; !ok {
			locMap[dep.Location] = map[string]struct{}{}
		}
		for _, obj := range dep.Objects {
			locMap[dep.Location][obj] = struct{}{}
		}
	}

	// build final import statements
	var importsList []string
	for loc, objs := range locMap {
		// sort objects for deterministic output
		objSlice := make([]string, 0, len(objs))
		for obj := range objs {
			objSlice = append(objSlice, obj)
		}
		sort.Strings(objSlice)
		statement := fmt.Sprintf("import { %s } from '%s';", strings.Join(objSlice, ", "), loc)
		importsList = append(importsList, statement)
	}

	// sort imports by location for consistency
	sort.Strings(importsList)

	// combine with actual script
	return strings.Join(importsList, "\r\n")
}
