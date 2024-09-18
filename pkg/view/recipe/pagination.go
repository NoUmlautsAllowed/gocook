package recipe

import "math"

type tmplPagination struct {
	BaseURL   string
	Separator string

	PreviousButOne tmplPageData
	Previous       tmplPageData
	Current        tmplPageData
	Next           tmplPageData
	NextButOne     tmplPageData
	LastButOne     tmplPageData
	Last           tmplPageData
}

type tmplPageData struct {
	Offset int
	Page   int
}

func pagination(perPage, offset, count int, baseURL string, urlHasParams bool) tmplPagination {
	var previousOffset int
	if offset == 0 {
		previousOffset = 0
	} else {
		previousOffset = offset - perPage
	}

	nextOffset := offset + perPage
	pageCount := int(math.Ceil(float64(count) / float64(perPage)))
	if count < perPage {
		nextOffset = 0
		pageCount = 1
	}

	separator := "?"
	if urlHasParams {
		separator = "&"
	}

	return tmplPagination{
		BaseURL:   baseURL,
		Separator: separator,
		PreviousButOne: tmplPageData{
			Offset: previousOffset - perPage,
			Page:   (previousOffset-perPage)/perPage + 1,
		},
		Previous: tmplPageData{
			Offset: previousOffset,
			Page:   previousOffset/perPage + 1,
		},
		Current: tmplPageData{
			Offset: offset,
			Page:   offset/perPage + 1,
		},
		Next: tmplPageData{
			Offset: nextOffset,
			Page:   nextOffset/perPage + 1,
		},
		NextButOne: tmplPageData{
			Offset: nextOffset + perPage,
			Page:   nextOffset/perPage + 2,
		},
		LastButOne: tmplPageData{
			Offset: (pageCount - 2) * perPage,
			Page:   pageCount - 1,
		},
		Last: tmplPageData{
			Offset: (pageCount - 1) * perPage,
			Page:   pageCount,
		},
	}
}
