package m3js

// Combines multiple parts of an Module3Action definition into a single file and generates
// the webrequestX based class for communication

import (
	"github.com/torabian/fireback/modules/fireback/module3/mcore"
)

func AsFullDocument(x *mcore.CodeChunkCompiled) string {
	importsList := CombineImportsJsWorld(*x)
	var finalContent string = importsList + "\r\n" + string(x.ActualScript)

	return finalContent
}

// Combines entire features for a module, and creates a virtual map of the files
func JsModuleFullVirtualFiles(module *mcore.Module3, ctx mcore.MicroGenContext) ([]mcore.VirtualFile, error) {

	files := []mcore.VirtualFile{}

	// step1, is to compile all of the actions, since they are most important.
	for _, action := range module.Actions {

		actionRendered, err := JsActionClass(action, ctx)
		if err != nil {
			return nil, err
		}

		files = append(files, mcore.VirtualFile{
			Name:         actionRendered.SuggestedFileName,
			Extension:    actionRendered.SuggestedExtension,
			ActualScript: AsFullDocument(actionRendered),
		})

	}

	return files, nil
}
