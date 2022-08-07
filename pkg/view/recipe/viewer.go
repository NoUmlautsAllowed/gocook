package recipe

import (
	"chefcook/pkg/api"
	"github.com/gin-gonic/gin"
)

type Viewer interface {
	ShowSearchResults(c *gin.Context)
	ShowRecipe(c *gin.Context)
}

type TemplateViewer struct {
	searchResultsTemplate string
	recipeTemplate        string
	api                   api.RecipeApi
}

func NewTemplateViewer(api api.RecipeApi) *TemplateViewer {
	return &TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   api,
	}
}

func RegisterViewerRoutes(v Viewer, r gin.IRouter) {
	r.GET("/recipe", v.ShowSearchResults)
	r.GET("/recipes/:recipe", v.ShowRecipe)
}
