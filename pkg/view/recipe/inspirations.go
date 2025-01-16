package recipe

import (
	"net/http"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
)

type tmplInspirations struct {
	api.RecipeInspirationsMixed
	api.Search
}

func (t *TemplateViewer) ShowInspirations(c *gin.Context) {
	inspirationsMixed, err := t.api.Inspirations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: gin.ErrorTypeNu,
			Meta: nil,
		})
		return
	}

	c.HTML(http.StatusOK, t.inspirationsTemplate, tmplInspirations{
		RecipeInspirationsMixed: *inspirationsMixed,
	})
}
