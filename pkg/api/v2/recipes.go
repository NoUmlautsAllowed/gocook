package v2

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Search struct {
	Query string `form:"query"`
}

type tmplRecipe struct {
	Recipe
	Search
	Instructions []string `json:"instructions"`
}

type tmplSearch struct {
	Search
	RecipeSearch
}

func GetRecipe(c *gin.Context) {
	recipe := c.Param("recipe")
	u, _ := url.Parse(ApiBaseRecipeUrl)
	u.Path = path.Join(u.Path, recipe)
	resp, err := http.Get(u.String())
	log.Println(resp.StatusCode, u)
	if err != nil {
		c.JSON(500, gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		})
	} else {
		data, _ := io.ReadAll(resp.Body)
		var recipe Recipe
		_ = json.Unmarshal(data, &recipe)
		tmpl := tmplRecipe{
			Recipe:       recipe,
			Instructions: strings.Split(recipe.Instructions, "\n\n"),
		}
		tmpl.PreviewImageURLTemplate = setPreviewImageFormat(tmpl.PreviewImageURLTemplate)
		c.HTML(http.StatusOK, "recipe.tmpl", tmpl)
	}
}

func SearchRecipes(c *gin.Context) {

	var search Search
	if c.ShouldBind(&search) == nil {
		// this is how the api call looks like
		// https://api.chefkoch.de/v2/search/recipe?query=lasagne%20vegan
		// this is how the format should look like crop-480x600
		// https://img.chefkoch-cdn.de/rezepte/2812481433250378/bilder/1185849/<format>/die-ultimative-vegane-lasagne.jpg
		// https://img.chefkoch-cdn.de/rezepte/2812481433250378/bilder/1185849/crop-480x600/die-ultimative-vegane-lasagne.jpg
		// this would be a query for pagination
		// https://api.chefkoch.de/v2/search-frontend/recipes?query=Lasagne+Vegan&limit=41&offset=41&analyticsTags=user,user_logged_out&enableClickAnalytics=true
		u, _ := url.Parse(ApiBaseSearchUrl)
		u.Path = path.Join(u.Path, "recipe")
		query := make(url.Values)
		query.Set("query", search.Query)
		u.RawQuery = query.Encode()

		resp, err := http.Get(u.String())
		log.Println(resp.StatusCode, u)
		if err != nil {
			c.JSON(500, gin.Error{
				Err:  err,
				Type: 0,
				Meta: nil,
			})
		} else {
			data, _ := io.ReadAll(resp.Body)
			var recipeSearch RecipeSearch
			_ = json.Unmarshal(data, &recipeSearch)
			for i, _ := range recipeSearch.Results {
				r := &recipeSearch.Results[i]
				r.Recipe.PreviewImageURLTemplate = setPreviewImageFormat(r.Recipe.PreviewImageURLTemplate)
			}

			t := tmplSearch{
				Search:       search,
				RecipeSearch: recipeSearch,
			}
			c.HTML(http.StatusOK, "results.tmpl", t)
		}
	}
}
