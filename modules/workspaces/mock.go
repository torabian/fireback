package workspaces

type MockQueryContext struct {
	WithPreloads []string
	Languages    []string
	Deep         bool
	ItemsPerPage int
}
