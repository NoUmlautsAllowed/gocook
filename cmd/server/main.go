package main

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"

	v2 "codeberg.org/NoUmlautsAllowed/gocook/pkg/api/v2"
	"codeberg.org/NoUmlautsAllowed/gocook/pkg/cdn"
	"codeberg.org/NoUmlautsAllowed/gocook/pkg/cdn/img"
	"codeberg.org/NoUmlautsAllowed/gocook/pkg/env"
	"codeberg.org/NoUmlautsAllowed/gocook/pkg/view/recipe"
	"codeberg.org/NoUmlautsAllowed/gocook/web"

	"github.com/gin-gonic/gin"
)

func main() {
	runEnv := env.NewEnv()
	log.Println("Using given environment configuration", runEnv)

	r := gin.Default()
	r.SetHTMLTemplate(template.Must(template.ParseFS(web.Templates, "templates/*")))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	staticFS, err := fs.Sub(web.Static, "static")
	if err != nil {
		log.Fatal(err)
	}
	r.StaticFS("static/", http.FS(staticFS))

	v := recipe.NewTemplateViewer(v2.NewV2Api(runEnv, cdn.RedirectImageCdnBasePath))
	recipe.RegisterViewerRoutes(v, r)

	imgCdn := img.NewImageCdn(runEnv)
	cdn.RegisterRoutes(imgCdn, r)

	log.Fatal(r.Run(runEnv.BindAddress())) // listen and serve
}
