package v2

import (
	"github.com/NoUmlautsAllowed/gocook/pkg/cdn"
	"testing"
)

func TestNewV2Api(t *testing.T) {
	api := NewV2Api(cdn.ImageCdnBaseUrl)

	if api.baseSearchUrl != apiBaseSearchUrl ||
		api.baseRecipeUrl != apiBaseRecipeUrl ||
		api.cdnBaseImageUrl != cdn.ImageCdnBaseUrl {
		t.Error("NewV2Api not populated with default api urls")
	}
}
