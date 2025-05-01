package javainclude

import "embed"

//go:embed *
var SpringInclude embed.FS

// This will include everything in this directory and put into the target
