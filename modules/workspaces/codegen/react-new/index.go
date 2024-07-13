package firebackgonew

import "embed"

//go:embed .eslintignore .eslintrc.json .nojekyll craco.config.js package-lock.json package.json README.md tsconfig.json types.d.ts Makefile .gitignore src/* public/* .vscode src/apps/projectname/.env.bundle src/apps/projectname/.env.local.txt src/apps/projectname/.env.githubpages src/apps/projectname/.env.static
var FbReactjsNewTemplate embed.FS

/**
*	This directory includes a boilerplate for building react.js apps,
*	it's configured in a way that uses fireback as backend.
*   Nevertheless, it's not forced at all to use backend for fireback.
*	You can remove src/modules/fireback folder, adjust App.tsx a bit
*	and start your fully pure project in there.
 */
