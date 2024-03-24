package main

import (
	"log"
	"net/http"

	v2 "github.com/NoUmlautsAllowed/gocook/pkg/api/v2"
	"github.com/NoUmlautsAllowed/gocook/pkg/cdn"
	"github.com/NoUmlautsAllowed/gocook/pkg/cdn/img"
	"github.com/NoUmlautsAllowed/gocook/pkg/env"
	"github.com/NoUmlautsAllowed/gocook/pkg/view/recipe"

	"github.com/gin-gonic/gin"
)

func main() {
	runEnv := env.NewEnv()
	log.Println("Using given environment configuration", runEnv)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	r.Static("static/", "static/")

	v := recipe.NewTemplateViewer(v2.NewV2Api(runEnv, cdn.RedirectImageCdnBasePath))
	recipe.RegisterViewerRoutes(v, r)

	imgCdn := img.NewImageCdn(runEnv)
	cdn.RegisterRoutes(imgCdn, r)

	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}
