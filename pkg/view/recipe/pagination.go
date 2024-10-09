package recipe

import (
	"math"
	"net/url"
	"strconv"
	"strings"
)

type tmplPagination struct {
	BaseURL string

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
	URL    string
}

func newTmplPageData(offset, page int, values url.Values) tmplPageData {
	values.Set("offset", strconv.Itoa(offset))

	return tmplPageData{
		Offset: offset,
		Page:   page,
		URL:    values.Encode(),
	}
}

func pagination(perPage, offset, count int, baseURL string, values url.Values) tmplPagination {
	var previousOffset int
	if offset == 0 {
		previousOffset = 0
	} else {
		previousOffset = offset - perPage
	}

	if !strings.HasSuffix(baseURL, "?") {
		baseURL += "?"
	}

	nextOffset := offset + perPage
	pageCount := int(math.Ceil(float64(count) / float64(perPage)))
	if count < perPage {
		nextOffset = 0
		pageCount = 1
	}

	return tmplPagination{
		BaseURL:        baseURL,
		PreviousButOne: newTmplPageData(previousOffset-perPage, (previousOffset-perPage)/perPage+1, values),
		Previous:       newTmplPageData(previousOffset, previousOffset/perPage+1, values),
		Current:        newTmplPageData(offset, offset/perPage+1, values),
		Next:           newTmplPageData(nextOffset, nextOffset/perPage+1, values),
		NextButOne:     newTmplPageData(nextOffset+perPage, nextOffset/perPage+2, values),
		LastButOne:     newTmplPageData((pageCount-2)*perPage, pageCount-1, values),
		Last:           newTmplPageData((pageCount-1)*perPage, pageCount, values),
	}
}
