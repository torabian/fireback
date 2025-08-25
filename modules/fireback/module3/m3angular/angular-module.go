package m3angular

// Combines multiple parts of an Module3Action definition into a single file and generates
// the webrequestX based class for communication

import (
	"errors"

	"github.com/torabian/fireback/modules/fireback/module3/m3js"
	"github.com/torabian/fireback/modules/fireback/module3/mcore"
)

// Combines entire features for a module, and creates a virtual map of the files
func AngularModuleFullVirtualFiles(module *mcore.Module3, ctx mcore.MicroGenContext) ([]mcore.VirtualFile, error) {

	files := []mcore.VirtualFile{}

	// step 1 - create the actions services
	if result, err := AngularActionsClass(module, ctx); err != nil {
		return nil, errors.New("angular actions class generation failed")
	} else {

		importsList := m3js.CombineImportsJsWorld(*result)

		files = append(files, mcore.VirtualFile{
			Name:         result.SuggestedFileName,
			ActualScript: importsList + "\r\n" + string(result.ActualScript),
			Extension:    result.SuggestedExtension,
		})
	}

	return files, nil
}
