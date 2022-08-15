package v2

import "testing"

func TestNewV2Api(t *testing.T) {
	api := NewV2Api()

	if api.baseSearchUrl != apiBaseSearchUrl || api.baseRecipeUrl != apiBaseRecipeUrl {
		t.Error("NewV2Api not populated with default api urls")
	}
}
