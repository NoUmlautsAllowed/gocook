package recipe

import (
	"net/http"
	"strings"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
)

type tmplRecipe struct {
	api.Recipe
	api.Search
	Instructions []string `json:"instructions"`
}

func (t *TemplateViewer) ShowRecipe(c *gin.Context) {
	recipe, err := t.api.Get(c.Param("recipe"))

	if err != nil {
		c.JSON(500, gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		})
	} else {
		tmplData := tmplRecipe{
			Recipe:       *recipe,
			Instructions: strings.Split(recipe.Instructions, "\n\n"),
		}
		c.HTML(http.StatusOK, t.recipeTemplate, tmplData)
	}
}
