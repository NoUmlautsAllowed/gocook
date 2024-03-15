package v2

import (
	"testing"

	"github.com/NoUmlautsAllowed/gocook/pkg/cdn"
	"github.com/NoUmlautsAllowed/gocook/pkg/env"
)

func TestNewV2Api(t *testing.T) {
	runEnv := env.NewEnv()
	api := NewV2Api(runEnv, cdn.RedirectImageCdnBasePath)

	if api.baseSearchURL != runEnv.APIBaseURL()+apiBaseSearchPath ||
		api.baseRecipeURL != runEnv.APIBaseURL()+apiBaseRecipePath ||
		api.cdnBaseImageURL != cdn.RedirectImageCdnBasePath {
		t.Error("NewV2Api not populated with default api urls")
	}
}
