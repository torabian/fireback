package swiftinclude

import "embed"

//go:embed *
var SwiftInclude embed.FS

// This will include everything in this directory and put into the target
