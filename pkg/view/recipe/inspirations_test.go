package recipe

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestTemplateViewer_ShowInspirations(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeAPI(ctrl)
	m.EXPECT().Inspirations().Return(&api.RecipeInspirationsMixed{}, nil)

	v := TemplateViewer{
		commentsTemplate: "inspirations.tmpl",
		api:              m,
	}

	r := gin.Default()
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/explore")

	w := httptest.ResponseRecorder{}
	req := http.Request{
		Method: http.MethodGet,
		URL:    u,
	}

	r.ServeHTTP(&w, &req)

	if w.Code != http.StatusOK {
		t.Error("expected status 200")
	}
}
