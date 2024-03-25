package recipe

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestTemplateViewer_ShowSearchResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeAPI(ctrl)
	m.EXPECT().Search(api.Search{Query: "schnitzel", Limit: strconv.Itoa(defaultResultsPerPage)}).Return(&api.RecipeSearch{}, nil)

	v := TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   m,
	}

	r := gin.Default()
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
	m := api.NewMockRecipeAPI(ctrl)
	m.EXPECT().Search(api.Search{Query: "pizza", Limit: strconv.Itoa(defaultResultsPerPage)}).Return(nil, errors.New("sample error"))

	v := TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   m,
	}

	r := gin.Default()
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

func TestTemplateViewer_ShowSearchResults_EmptyQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeAPI(ctrl)

	v := TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   m,
	}

	r := gin.Default()
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/recipe?q=fries")

	w := httptest.ResponseRecorder{}
	req := http.Request{
		Method: http.MethodGet,
		URL:    u,
	}

	r.ServeHTTP(&w, &req)

	if w.Code != http.StatusMovedPermanently {
		t.Error("expected status 301")
	}
}

func TestTemplateViewer_ShowSearchResults_Offset(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeAPI(ctrl)
	m.EXPECT().Search(api.Search{Query: "schnitzel", Limit: strconv.Itoa(defaultResultsPerPage), Offset: "4"}).Return(&api.RecipeSearch{}, nil)

	v := TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   m,
	}

	r := gin.Default()
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/recipe?query=schnitzel&offset=4")

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

func TestTemplateViewer_ShowSearchResults_Offset2(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeAPI(ctrl)
	m.EXPECT().Search(api.Search{Query: "schnitzel", Limit: strconv.Itoa(defaultResultsPerPage), Offset: "donotsetmeoff"}).Return(&api.RecipeSearch{}, nil)

	v := TemplateViewer{
		searchResultsTemplate: "search.tmpl",
		recipeTemplate:        "recipe.tmpl",
		api:                   m,
	}

	r := gin.Default()
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/recipe?query=schnitzel&offset=donotsetmeoff")

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
