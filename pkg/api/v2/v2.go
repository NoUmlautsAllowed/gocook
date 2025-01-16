package v2

import (
	"net/http"
	"time"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/env"
)

const (
	apiBaseRecipePath       = "/v2/recipes"
	apiBaseInspirationsPath = "/v2/inspiration"
	apiBaseSearchPath       = "/v2/search-gateway"
)

const previewImageFormat = "crop-480x600"

type API struct {
	baseRecipeURL       string
	baseSearchURL       string
	baseInspirationsURL string
	cdnBaseImageURL     string
	defaultClient       http.Client
	userAgent           string
}

func NewV2Api(e *env.Env, redirectImageCdnBasePath string) *API {
	return &API{
		baseRecipeURL:       e.APIBaseURL() + apiBaseRecipePath,
		baseSearchURL:       e.APIBaseURL() + apiBaseSearchPath,
		baseInspirationsURL: e.APIBaseURL() + apiBaseInspirationsPath,
		cdnBaseImageURL:     redirectImageCdnBasePath,
		defaultClient: http.Client{
			Timeout: 30 * time.Second,
		},
		userAgent: e.UserAgent(),
	}
}
