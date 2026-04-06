package fireback

type GoogleResponseData[V any] struct {
	Item V `json:"item"`
}

type GoogleResponse[V any] struct {
	Data GoogleResponseData[V] `json:"data"`
}

func GResponseSingleItem[T any](v T) GoogleResponse[T] {
	return GoogleResponse[T]{
		Data: GoogleResponseData[T]{Item: v},
	}
}
