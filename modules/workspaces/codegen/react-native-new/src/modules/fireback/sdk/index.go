package typescriptinclude

import "embed"

//go:embed *
var TypeScriptInclude embed.FS

// This will include everything in this directory and put into the target
