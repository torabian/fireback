package fireback

func GResponseSingleItem(v any) any {
	return map[string]any{
		"data": map[string]any{
			"item": v,
		},
	}
}
