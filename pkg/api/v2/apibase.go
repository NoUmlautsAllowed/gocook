package v2

import "strings"

const ApiBaseRecipeUrl = "https://api.chefkoch.de/v2/recipes"
const ApiBaseSearchUrl = "https://api.chefkoch.de/v2/search"

const previewImageFormat = "crop-480x600"

func setPreviewImageFormat(in string) string {
	return strings.ReplaceAll(in, "<format>", previewImageFormat)
}
