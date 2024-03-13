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
	baseRecipeUrl   string
	baseSearchUrl   string
	cdnBaseImageUrl string
	defaultClient   http.Client
	userAgent       string
}

func NewV2Api(e *env.Env, redirectImageCdnBasePath string) *V2Api {
	return &V2Api{
		baseRecipeUrl:   e.ApiBaseUrl() + apiBaseRecipePath,
		baseSearchUrl:   e.ApiBaseUrl() + apiBaseSearchPath,
		cdnBaseImageUrl: redirectImageCdnBasePath,
		defaultClient: http.Client{
			Timeout: 30 * time.Second,
		},
		userAgent: e.UserAgent(),
	}
}
