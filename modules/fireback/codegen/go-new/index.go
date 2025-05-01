package firebackgonew

import "embed"

//go:embed README.md Makefile.tpl go.sum.tpl go.mod.tpl .gitignore.tpl modules cmd .vscode cmd docs/blog docs/docs docs/src docs/src docs/static docs/.gitignore docs/docusaurus.config.js docs/package-lock.json docs/package.json docs/README.md docs/sidebars.js e2e/cypress e2e/cypress.config.js e2e/package-lock.json e2e/package.json .github
var FbGoNewTemplate embed.FS
