package recipe

import (
	"github.com/NoUmlautsAllowed/gocook/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type tmplSearch struct {
	api.Search
	api.RecipeSearch
	ResultsPerPage int
	ResultsPerRow  int
}

const defaultResultsPerPage int = 12
const defaultResultsPerRow int = 3

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

		// use a multiple of 3 here
		// see search.tmpl and iterateRange template function
		// this is used to have a reasonable amount of recipes to show per page
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
		} else {

			tmplData := tmplSearch{
				Search:         search,
				RecipeSearch:   *recipeSearch,
				ResultsPerPage: defaultResultsPerPage,
				ResultsPerRow:  defaultResultsPerRow,
			}
			c.HTML(http.StatusOK, t.searchResultsTemplate, tmplData)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		})
	}
}
