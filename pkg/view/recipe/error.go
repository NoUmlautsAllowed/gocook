package recipe

import (
	"errors"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api/v2"

	"github.com/gin-gonic/gin"
)

type errorContext struct {
	StatusCode int
	Error      error
}

func mapErrorMessage(err error) string {
	switch {
	case errors.Is(err, v2.ErrRequestForbidden):
		return "Anfrage vom Upstream-Server verweigert"
	case errors.Is(err, v2.ErrRequestFailed):
		return "Anfrage fehlgeschlagen"
	}

	return err.Error()
}

func (t *TemplateViewer) ShowErrorPage(c *gin.Context, ctx errorContext) {
	tmplData := struct {
		ErrorCode    int
		ErrorMessage string
	}{
		ErrorCode:    ctx.StatusCode,
		ErrorMessage: mapErrorMessage(ctx.Error),
	}

	c.HTML(ctx.StatusCode, t.errorTemplate, tmplData)
}
