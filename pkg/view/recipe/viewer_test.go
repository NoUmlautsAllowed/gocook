package recipe

import (
	"testing"

	"github.com/NoUmlautsAllowed/gocook/pkg/api"

	"go.uber.org/mock/gomock"
)

func TestNewTemplateViewer(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := api.NewMockRecipeApi(ctrl)
	v := NewTemplateViewer(m)

	if v.searchResultsTemplate != "results.tmpl" {
		t.Error("expected results.tmpl as template")
	}

	if v.recipeTemplate != "recipe.tmpl" {
		t.Error("expected recipe.tmpl as template")
	}
}
