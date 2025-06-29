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
		t.ShowErrorPage(c, errorContext{
			StatusCode: http.StatusBadRequest,
			Error:      err,
		})
		return
	}

	c.HTML(http.StatusOK, t.inspirationsTemplate, tmplInspirations{
		RecipeInspirationsMixed: *inspirationsMixed,
	})
}
