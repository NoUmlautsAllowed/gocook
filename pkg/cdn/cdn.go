package cdn

import (
	"github.com/gin-gonic/gin"
)

type Image interface {
	GetRawImage(method, path string) ([]byte, error)
	GetImage(c *gin.Context)
}

const RedirectImageCdnBasePath string = "/img"

func RegisterRoutes(i Image, r gin.IRouter) {
	r.GET(RedirectImageCdnBasePath+"/*path", i.GetImage)
	r.HEAD(RedirectImageCdnBasePath+"/*path", i.GetImage)
}
