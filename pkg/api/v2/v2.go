package v2

const apiBaseRecipeUrl = "https://api.chefkoch.de/v2/recipes"
const apiBaseSearchUrl = "https://api.chefkoch.de/v2/search"

const previewImageFormat = "crop-480x600"

type V2Api struct {
	baseRecipeUrl string
	baseSearchUrl string
}

func NewV2Api() *V2Api {
	return &V2Api{
		baseRecipeUrl: apiBaseRecipeUrl,
		baseSearchUrl: apiBaseSearchUrl,
	}
}
