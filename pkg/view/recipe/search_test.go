package recipe

import (
	"errors"
	"github.com/NoUmlautsAllowed/gocook/pkg/api"
	"github.com/NoUmlautsAllowed/gocook/pkg/utils/tmpl"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestTemplateViewer_ShowSearchResults(t *testing.T) {

	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeApi(ctrl)
	m.EXPECT().Search(api.Search{Query: "schnitzel"}).Return(&api.RecipeSearch{}, nil)

	v := TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   m,
	}

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"iterateRange": tmpl.IterateRange[api.RecipeSearchResult],
	})
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/recipe?query=schnitzel")

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

func TestTemplateViewer_ShowSearchResults_InternalError(t *testing.T) {

	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeApi(ctrl)
	m.EXPECT().Search(api.Search{Query: "pizza"}).Return(nil, errors.New("sample error"))

	v := TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   m,
	}

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"iterateRange": tmpl.IterateRange[api.RecipeSearchResult],
	})
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/recipe?query=pizza")

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

func TestTemplateViewer_ShowSearchResults_BadRequest(t *testing.T) {

	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeApi(ctrl)

	v := TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   m,
	}

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"iterateRange": tmpl.IterateRange[api.RecipeSearchResult],
	})
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/recipe?q=fries")

	w := httptest.ResponseRecorder{}
	req := http.Request{
		Method: http.MethodGet,
		URL:    u,
	}

	r.ServeHTTP(&w, &req)

	if w.Code != http.StatusBadRequest {
		t.Error("expected status 400")
	}
}
