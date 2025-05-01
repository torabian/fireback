package angularinclude

import "embed"

//go:embed *
var AngularInclude embed.FS

// This will include everything in this directory and put into the target
