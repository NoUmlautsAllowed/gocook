package img

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"time"
)

const cdnBaseImageUrl = "https://img.chefkoch-cdn.de/"

type ImageCdn struct {
	cdnUrl        string
	defaultClient http.Client
}

func NewImageCdn() *ImageCdn {
	return &ImageCdn{
		cdnUrl: cdnBaseImageUrl,
		defaultClient: http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (c *ImageCdn) GetRawImage(method, imgPath string) ([]byte, error) {

	urlPath, err := url.JoinPath(c.cdnUrl, imgPath)
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	if method == http.MethodGet {
		resp, err = c.defaultClient.Get(urlPath)
	} else if method == http.MethodHead {
		// head method is possible with the CDN
		// however, a proper cache control must be implemented
		// it seems like the response header Cache-Control itself does not produce
		// the expected browser behavior
		resp, err = c.defaultClient.Head(urlPath)
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data []byte
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 && method != http.MethodHead {
		return nil, errors.New("no data received")
	}

	return data, nil
}

func (c *ImageCdn) GetImage(ctx *gin.Context) {

	if ctx.Request.Method != http.MethodGet && ctx.Request.Method != http.MethodHead {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "only GET allowed",
		})
		return
	}

	data, err := c.GetRawImage(ctx.Request.Method, ctx.Param("path"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// let gin itself handle the response headers
	// there may be an option to tell the requesting client to
	// use local caching for the image cdn
	// see also GetRawImage function
	_, err = ctx.Writer.Write(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

}
