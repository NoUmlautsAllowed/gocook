package recipe

import (
	"errors"
	"github.com/NoUmlautsAllowed/gocook/pkg/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestTemplateViewer_ShowRecipe(t *testing.T) {

	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeApi(ctrl)
	m.EXPECT().Get("1234567890").Return(&api.Recipe{}, nil)

	v := TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   m,
	}

	r := gin.Default()
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/recipes/1234567890")

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

func TestTemplateViewer_ShowRecipe_InternalError(t *testing.T) {

	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeApi(ctrl)
	m.EXPECT().Get("1234567890").Return(nil, errors.New("sample error"))

	v := TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   m,
	}

	r := gin.Default()
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/recipes/1234567890")

	w := httptest.ResponseRecorder{}
	req := http.Request{
		Method: http.MethodGet,
		URL:    u,
	}

	r.ServeHTTP(&w, &req)

	if w.Code != http.StatusInternalServerError {
		t.Error("expected status 500")
	}

}
