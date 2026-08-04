package firebackcapacitor

import "embed"

//go:embed set-ip.js switch-build.js README.md package.json capacitor.config.json .gitignore
var FbReactCapacitorNewTemplate embed.FS

/**
*	This directory includes a boilerplate for building react.js apps,
*	it's configured in a way that uses fireback as backend.
*   Nevertheless, it's not forced at all to use backend for fireback.
*	You can remove src/modules/fireback folder, adjust App.tsx a bit
*	and start your fully pure project in there.
 */
