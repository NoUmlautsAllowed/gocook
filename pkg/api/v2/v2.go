package v2

import (
	"net/http"
	"time"

	"github.com/NoUmlautsAllowed/gocook/pkg/env"
)

const (
	apiBaseRecipePath = "/v2/recipes"
	apiBaseSearchPath = "/v2/search"
)

const previewImageFormat = "crop-480x600"

type V2Api struct {
	baseRecipeURL   string
	baseSearchURL   string
	cdnBaseImageURL string
	defaultClient   http.Client
	userAgent       string
}

func NewV2Api(e *env.Env, redirectImageCdnBasePath string) *V2Api {
	return &V2Api{
		baseRecipeURL:   e.APIBaseURL() + apiBaseRecipePath,
		baseSearchURL:   e.APIBaseURL() + apiBaseSearchPath,
		cdnBaseImageURL: redirectImageCdnBasePath,
		defaultClient: http.Client{
			Timeout: 30 * time.Second,
		},
		userAgent: e.UserAgent(),
	}
}
