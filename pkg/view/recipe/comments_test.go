package recipe

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestTemplateViewer_ShowComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeAPI(ctrl)
	m.EXPECT().Comments(api.CommentQuery{RecipeID: "1983941321710773", Limit: 20}).Return(&api.Comments{}, nil)

	v := TemplateViewer{
		commentsTemplate: "comments.tmpl",
		api:              m,
	}

	r := gin.Default()
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/recipes/1983941321710773/comments")

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

func TestTemplateViewer_ShowComments_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeAPI(ctrl)
	m.EXPECT().Comments(api.CommentQuery{RecipeID: "1234567890", Limit: 20}).Return(nil, errors.New("sample error"))

	v := TemplateViewer{
		commentsTemplate: "comments.tmpl",
		api:              m,
	}

	r := gin.Default()
	r.LoadHTMLGlob("../../../templates/*")
	RegisterViewerRoutes(&v, r)

	u, _ := url.Parse("http://127.0.0.1:8080/recipes/1234567890/comments")

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
