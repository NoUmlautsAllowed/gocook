package recipe

import (
	"math"
	"net/http"
	"strconv"

	"github.com/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
)

type tmplPageData struct {
	Offset int
	Page   int
}

type tmplSearch struct {
	api.Search
	api.RecipeSearch
	ResultsPerPage int

	PreviousButOne tmplPageData
	Previous       tmplPageData
	Current        tmplPageData
	Next           tmplPageData
	NextButOne     tmplPageData
	LastButOne     tmplPageData
	Last           tmplPageData
}

const defaultResultsPerPage int = 15

func (t *TemplateViewer) ShowSearchResults(c *gin.Context) {
	var search api.Search
	if err := c.Bind(&search); err == nil && len(search.Query) > 0 {
		// this is how the api call looks like
		// https://api.chefkoch.de/v2/search/recipe?query=lasagne%20vegan
		// this is how the format should look like crop-480x600
		// https://img.chefkoch-cdn.de/rezepte/2812481433250378/bilder/1185849/<format>/die-ultimative-vegane-lasagne.jpg
		// https://img.chefkoch-cdn.de/rezepte/2812481433250378/bilder/1185849/crop-480x600/die-ultimative-vegane-lasagne.jpg
		// this would be a query for pagination
		// https://api.chefkoch.de/v2/search-frontend/recipes?query=Lasagne+Vegan&limit=41&offset=41&analyticsTags=user,user_logged_out&enableClickAnalytics=true

		// use a multiple of 3 and 5 here
		// this is used to have a reasonable amount of recipes to show per page on tablet and widescreen
		// the user is not allowed to set another value here
		// this value is handed over directly to the api
		search.Limit = strconv.Itoa(defaultResultsPerPage)

		recipeSearch, err := t.api.Search(search)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.Error{
				Err:  err,
				Type: 0,
				Meta: nil,
			})
			return
		}

		offset := 0
		if search.Offset != "" {
			var err error
			offset, err = strconv.Atoi(search.Offset)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.Error{
					Err:  err,
					Type: 0,
					Meta: nil,
				})
				return
			}
		}

		var previousOffset int
		if offset == 0 {
			previousOffset = 0
		} else {
			previousOffset = offset - defaultResultsPerPage
		}

		nextOffset := offset + defaultResultsPerPage
		pageCount := int(math.Ceil(float64(recipeSearch.Count) / float64(defaultResultsPerPage)))
		if recipeSearch.Count < defaultResultsPerPage {
			nextOffset = 0
			pageCount = 1
		}

		tmplData := tmplSearch{
			Search:         search,
			RecipeSearch:   *recipeSearch,
			ResultsPerPage: defaultResultsPerPage,
			PreviousButOne: tmplPageData{
				Offset: previousOffset - defaultResultsPerPage,
				Page:   (previousOffset-defaultResultsPerPage)/defaultResultsPerPage + 1,
			},
			Previous: tmplPageData{
				Offset: previousOffset,
				Page:   previousOffset/defaultResultsPerPage + 1,
			},
			Current: tmplPageData{
				Offset: offset,
				Page:   offset/defaultResultsPerPage + 1,
			},
			Next: tmplPageData{
				Offset: nextOffset,
				Page:   nextOffset/defaultResultsPerPage + 1,
			},
			NextButOne: tmplPageData{
				Offset: nextOffset + defaultResultsPerPage,
				Page:   nextOffset/defaultResultsPerPage + 2,
			},
			LastButOne: tmplPageData{
				Offset: (pageCount - 2) * defaultResultsPerPage,
				Page:   pageCount - 1,
			},
			Last: tmplPageData{
				Offset: (pageCount - 1) * defaultResultsPerPage,
				Page:   pageCount,
			},
		}
		c.HTML(http.StatusOK, t.searchResultsTemplate, tmplData)
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: gin.ErrorTypeBind,
			Meta: nil,
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
