package recipe

import (
	"github.com/NoUmlautsAllowed/gocook/pkg/api"
	"github.com/golang/mock/gomock"
	"testing"
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
