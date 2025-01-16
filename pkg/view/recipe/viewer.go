package recipe

import (
	"fmt"
	"net/http"
	"net/url"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
)

type Viewer interface {
	ShowSearchResults(c *gin.Context)
	ShowRecipe(c *gin.Context)
	ShowComments(c *gin.Context)
	ShowInspirations(c *gin.Context)
}

type TemplateViewer struct {
	searchResultsTemplate string
	recipeTemplate        string
	commentsTemplate      string
	inspirationsTemplate  string
	api                   api.RecipeAPI
}

const (
	searchResultsPath  = "recipe"
	redirectSearchPath = "rs/s:page/:query/Rezepte.html"
	recipePath         = "recipes/:recipe"
	redirectRecipePath = "rezepte/:recipe/*recipename"
	commentsPath       = "recipes/:recipe/comments"
	inspirationsPath   = "explore"
)

func NewTemplateViewer(api api.RecipeAPI) *TemplateViewer {
	return &TemplateViewer{
		searchResultsTemplate: "results.tmpl",
		recipeTemplate:        "recipe.tmpl",
		commentsTemplate:      "comments.tmpl",
		inspirationsTemplate:  "inspirations.tmpl",
		api:                   api,
	}
}

func RegisterViewerRoutes(v Viewer, r gin.IRouter) {
	r.GET(searchResultsPath, v.ShowSearchResults)
	r.GET(recipePath, v.ShowRecipe)
	r.GET(commentsPath, v.ShowComments)
	r.GET(inspirationsPath, v.ShowInspirations)

	// this path is used by chefkoch for displaying recipes and therefore this redirect makes URL rewrites simpler
	r.GET(redirectRecipePath, func(c *gin.Context) {
		recipe := c.Param("recipe")
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/recipes/%s", recipe))
	})

	// this path is used for search results at Chefkoch (and thus frequently returned by web search engines),
	// therefore this simplifies URL rewriting and improves compatibility with Chefkoch
	r.GET(redirectSearchPath, func(c *gin.Context) {
		query := url.QueryEscape(c.Param("query"))
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/recipe?query=%s", query))
	})
}
