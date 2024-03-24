package recipe

import (
	"github.com/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
)

type Viewer interface {
	ShowSearchResults(c *gin.Context)
	ShowRecipe(c *gin.Context)
}

type TemplateViewer struct {
	searchResultsTemplate string
	recipeTemplate        string
	api                   api.RecipeAPI
}

const (
	searchResultsPath = "/recipe"
	recipePath        = "recipes/:recipe"
)

func NewTemplateViewer(api api.RecipeAPI) *TemplateViewer {
	return &TemplateViewer{
		searchResultsTemplate: "results.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   api,
	}
}

func RegisterViewerRoutes(v Viewer, r gin.IRouter) {
	r.GET(searchResultsPath, v.ShowSearchResults)
	r.GET(recipePath, v.ShowRecipe)
}
