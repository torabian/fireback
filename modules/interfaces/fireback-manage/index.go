package swiftpl

import "embed"

//go:embed *
var FirebackManageTmpl embed.FS

/// Fireback manage is the same react-new project,
// but separately built for admin panel.
// In some projects if they do not want to use default ui,
// it's good to have the admin panel separate
// and they can continue on their project on front-end side on their own.
