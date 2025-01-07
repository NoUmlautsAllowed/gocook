package recipe

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

var testTagGroups = []api.TagGroup{
	{
		Key:        "forty",
		Name:       "forty group",
		IsActive:   true,
		IsDisabled: false,
		Tags: []api.Tag{
			{
				ID:         41,
				Name:       "forty one",
				Count:      0,
				IsActive:   false,
				IsDisabled: false,
			},
			{
				ID:         42,
				Name:       "the solution",
				Count:      0,
				IsActive:   true,
				IsDisabled: false,
			},
		},
	},
	{},
	{
		Key:        "other",
		Name:       "other tags",
		Icon:       "",
		IsActive:   false,
		IsDisabled: false,
		Tags: []api.Tag{
			{
				ID:         33,
				Name:       "thirty three",
				Count:      0,
				IsActive:   false,
				IsDisabled: false,
			},
		},
	},
	{
		Key:        "n",
		Name:       "ninety group",
		Icon:       "",
		IsActive:   true,
		IsDisabled: false,
		Tags: []api.Tag{
			{
				ID:         99,
				Name:       "ninety nine",
				Count:      0,
				IsActive:   true,
				IsDisabled: false,
			},
		},
	},
}

func TestTemplateViewer_ShowSearchResults(t *testing.T) {
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

	tests := []struct {
		name       string
		url        string
		method     string
		search     *api.Search
		mockResult *api.RecipeSearch
		mockErr    error
		statusCode int
	}{
		{
			name:       "show search results",
			url:        "http://127.0.0.1:8080/recipe?query=schnitzel",
			method:     http.MethodGet,
			search:     &api.Search{Query: "schnitzel", Limit: strconv.Itoa(defaultResultsPerPage)},
			mockResult: &api.RecipeSearch{},
			mockErr:    nil,
			statusCode: http.StatusOK,
		},
		{
			name:       "internal error",
			url:        "http://127.0.0.1:8080/recipe?query=pizza",
			method:     http.MethodGet,
			search:     &api.Search{Query: "pizza", Limit: strconv.Itoa(defaultResultsPerPage)},
			mockResult: nil,
			mockErr:    errors.New("sample error"),
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "empty query",
			url:        "http://127.0.0.1:8080/recipe?q=fries",
			method:     http.MethodGet,
			search:     nil,
			mockResult: nil,
			mockErr:    nil,
			statusCode: http.StatusMovedPermanently,
		},
		{
			name:       "URL in query",
			url:        "http://127.0.0.1:8080/recipe?query=https://www.chefkoch.de/rezepte/1983941321710773/Franzoesische-Apfeltarte.html",
			method:     http.MethodGet,
			search:     nil,
			mockResult: nil,
			mockErr:    nil,
			statusCode: http.StatusMovedPermanently,
		},
		{
			name:       "Offset 4",
			url:        "http://127.0.0.1:8080/recipe?query=schnitzel&offset=4",
			search:     &api.Search{Query: "schnitzel", Limit: strconv.Itoa(defaultResultsPerPage), Offset: "4"},
			method:     http.MethodGet,
			mockResult: &api.RecipeSearch{},
			mockErr:    nil,
			statusCode: http.StatusOK,
		},
		{
			name:       "Offset String",
			url:        "http://127.0.0.1:8080/recipe?query=schnitzel&offset=donotsetmeoff",
			search:     &api.Search{Query: "schnitzel", Limit: strconv.Itoa(defaultResultsPerPage), Offset: "donotsetmeoff"},
			method:     http.MethodGet,
			mockResult: &api.RecipeSearch{},
			mockErr:    nil,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:   "tags",
			url:    "http://127.0.0.1:8080/recipe?query=reis&tags=42,99",
			method: http.MethodGet,
			search: &api.Search{Query: "reis", Tags: "42,99", Limit: strconv.Itoa(defaultResultsPerPage)},
			mockResult: &api.RecipeSearch{
				TagGroups: testTagGroups,
			},
			mockErr:    nil,
			statusCode: http.StatusOK,
		},
		{
			name:       "Chefkoch search link redirect",
			url:        "http://127.0.0.1:8080/rs/s0/dampfnudeln/Rezepte.html",
			method:     http.MethodGet,
			search:     nil,
			mockResult: nil,
			mockErr:    nil,
			statusCode: http.StatusMovedPermanently,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := url.Parse(tt.url)
			if tt.search != nil {
				m.EXPECT().Search(*tt.search).Return(tt.mockResult, tt.mockErr)
			}

			w := httptest.ResponseRecorder{}
			req := http.Request{
				Method: tt.method,
				URL:    u,
			}

			r.ServeHTTP(&w, &req)

			if w.Code != tt.statusCode {
				t.Errorf("expected status %d, got %d", tt.statusCode, w.Code)
			}
		})
	}
}
