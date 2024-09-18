package v2

import (
	"testing"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/env"
)

func TestReplaceImageCdnUrl(t *testing.T) {
	// format tag as this is replaced for the CDN request
	// produce an error when replacing the cdn url
	api := NewV2Api(env.NewEnv(), "some raällly weird URL $$§ßßß???? \x01")
	input := "/cdn/<format>/xyz-img"

	result := api.replaceImageCdnURL(input)

	expected := setPreviewImageFormat(input)

	// no join of both urls will happen due to error parsing the cdn base image url
	// results only in replacing the preview image format tag <format>

	if result != expected {
		t.Error("expected", expected, "; got", result)
	}
}

func TestReplaceImageCdnUrl2(t *testing.T) {
	// format tag as this is replaced for the CDN request
	// produce an error when replacing the cdn url
	api := NewV2Api(env.NewEnv(), "some raällly weird URL $$§ßßß???? \x01")

	// a mal crafted input url
	input := "\x01"

	result := api.replaceImageCdnURL(input)

	if result != input {
		t.Error("expected", input, "; got", result)
	}
}
