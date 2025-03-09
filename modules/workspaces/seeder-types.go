package workspaces

// Used in entity bundle, to define what is the seeder content
type Seeder[T any] struct {
	Items []T `json:"items" yaml:"items"`
}
