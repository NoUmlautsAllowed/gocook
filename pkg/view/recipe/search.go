package recipe

import (
	"fmt"
	"maps"
	"net/http"
	"regexp"
	"strconv"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"
	"codeberg.org/NoUmlautsAllowed/gocook/pkg/form"

	"github.com/gin-gonic/gin"
)

type tmplSearch struct {
	api.Search
	api.RecipeSearch
	ResultsPerPage int
	// this recipe array allows using the recipe grid template
	// because the results array has a different format than
	// being a simple array of recipes
	Recipes []api.Recipe

	Pagination        tmplPagination
	TagGroupTemplates []tmplTagGroup
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

		values, err := form.Values(search)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.Error{
				Err:  err,
				Type: 0,
				Meta: nil,
			})
		}

		var recipes []api.Recipe
		for _, result := range recipeSearch.Results {
			recipes = append(recipes, result.Recipe)
		}

		tmplData := tmplSearch{
			Search:            search,
			RecipeSearch:      *recipeSearch,
			Recipes:           recipes,
			ResultsPerPage:    defaultResultsPerPage,
			Pagination:        pagination(defaultResultsPerPage, offset, recipeSearch.Count, "/recipe", maps.Clone(values)),
			TagGroupTemplates: tagGroupTemplates("/recipe", *recipeSearch, maps.Clone(values)),
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
