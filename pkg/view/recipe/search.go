package recipe

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
)

type tmplSearch struct {
	api.Search
	api.RecipeSearch
	ResultsPerPage int

	Pagination tmplPagination
}

var recipeURLRegex = regexp.MustCompile(`^(?:https:\/\/|)(?:www\.|)chefkoch\.de\/rezepte\/(\d+)\/[a-zA-Z-]+\.html[a-zA-Z0-9-_=&?#]*$`)

const defaultResultsPerPage int = 15

func (t *TemplateViewer) ShowSearchResults(c *gin.Context) {
	var search api.Search
	if err := c.Bind(&search); err == nil && len(search.Query) > 0 {
		// first of all, we check whether this is a direct url
		urlMatch := recipeURLRegex.FindStringSubmatch(search.Query)
		if len(urlMatch) > 0 {
			// redirect to direct recipe view
			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/recipes/%s", urlMatch[1]))
			return
		}

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

		tmplData := tmplSearch{
			Search:         search,
			RecipeSearch:   *recipeSearch,
			ResultsPerPage: defaultResultsPerPage,
			Pagination:     pagination(defaultResultsPerPage, offset, recipeSearch.Count, "/recipe?query="+search.Query, true),
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
