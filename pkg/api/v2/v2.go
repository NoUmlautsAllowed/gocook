package v2

import (
	"net/http"
	"time"
)

const apiBaseRecipeUrl = "https://api.chefkoch.de/v2/recipes"
const apiBaseSearchUrl = "https://api.chefkoch.de/v2/search"

const previewImageFormat = "crop-480x600"

type V2Api struct {
	baseRecipeUrl string
	baseSearchUrl string
	defaultClient http.Client
}

func NewV2Api() *V2Api {
	return &V2Api{
		baseRecipeUrl: apiBaseRecipeUrl,
		baseSearchUrl: apiBaseSearchUrl,
		defaultClient: http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
