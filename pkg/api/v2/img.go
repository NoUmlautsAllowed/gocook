package v2

import (
	"net/url"
	"strings"
)

func setPreviewImageFormat(in string) string {
	return strings.ReplaceAll(in, "<format>", previewImageFormat)
}

func (a *V2Api) replaceImageCdnUrl(in string) string {
	out := setPreviewImageFormat(in)

	u, err := url.Parse(out)
	if err != nil {
		// should not happen as teh input comes directly out of the API
		// anyway, return the CDN with set preview image format by default
		return out
	}

	p, err := url.JoinPath(a.cdnBaseImageUrl, u.Path)
	if err != nil {
		// same as above
		// fallback is default CDN
		return out
	}

	// host part is stripped away at this time
	return p
}
