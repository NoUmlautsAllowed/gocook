package img

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/env"

	"github.com/gin-gonic/gin"
)

type ImageCdn struct {
	cdnURL        string
	defaultClient http.Client
	userAgent     string
}

func NewImageCdn(e *env.Env) *ImageCdn {
	return &ImageCdn{
		cdnURL: e.CdnBaseURL(),
		defaultClient: http.Client{
			Timeout: 60 * time.Second,
		},
		userAgent: e.UserAgent(),
	}
}

func (c *ImageCdn) GetRawImage(method, imgPath string) ([]byte, error) {
	urlPath, err := url.JoinPath(c.cdnURL, imgPath)
	if err != nil {
		return nil, err
	}

	// head method is possible with the CDN
	// however, a proper cache control must be implemented on the client side
	// see GetImage function below where Cache-Control header is set
	if method != http.MethodGet && method != http.MethodHead {
		return nil, errors.New("only GET or HEAD method allowed")
	}

	req, err := http.NewRequest(method, urlPath, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)
	resp, err := c.defaultClient.Do(req)
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

	// set Cache-Control header so images are cached on the client for 7d
	ctx.Writer.Header().Set("Cache-Control", "max-age=604800")
	_, err = ctx.Writer.Write(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}
