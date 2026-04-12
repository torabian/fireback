package fireback

type NextPageInfo struct {
	Cursor string `json:"cursor"`
}

type GoogleResponseData[V any] struct {
	Item                V            `json:"item"`
	Items               []V          `json:"items"`
	Next                NextPageInfo `json:"next"`
	TotalItems          int64        `json:"totalItems"`
	TotalAvailableItems int64        `json:"totalAvailableItems"`
	StartIndex          int64        `json:"startIndex"`
	ItemsPerPage        int64        `json:"itemsPerPage"`
}

type GoogleResponse[V any] struct {
	Data GoogleResponseData[V] `json:"data"`
}

func GResponseSingleItem[T any](v T) GoogleResponse[T] {
	return GoogleResponse[T]{
		Data: GoogleResponseData[T]{Item: v},
	}
}

func GResponseQuery[T any](v []T, meta *QueryResultMeta, q *QueryDSL) GoogleResponse[T] {

	res := GoogleResponse[T]{
		Data: GoogleResponseData[T]{Items: v},
	}

	if meta != nil {
		if meta.Cursor != nil {
			res.Data.Next.Cursor = *meta.Cursor
		}
		res.Data.TotalItems = meta.TotalItems
		res.Data.TotalAvailableItems = meta.TotalAvailableItems
		res.Data.StartIndex = int64(q.StartIndex)
		res.Data.ItemsPerPage = int64(q.ItemsPerPage)
	}

	return res
}
