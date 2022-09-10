package cdn

import "github.com/gin-gonic/gin"

type Image interface {
	GetRawImage(method, path string) ([]byte, error)
	GetImage(c *gin.Context)
}

const ImageCdnBaseUrl string = "/img"

func RegisterRoutes(i Image, r gin.IRouter) {
	r.GET(ImageCdnBaseUrl+"/*path", i.GetImage)
	r.HEAD(ImageCdnBaseUrl+"/*path", i.GetImage)
}
