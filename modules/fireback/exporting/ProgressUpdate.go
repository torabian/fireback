package exporting

type ProgressUpdate struct {
	ItemsProcessed int
	Message        string
	Complete       bool
	Error          error
}
