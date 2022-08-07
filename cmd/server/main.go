package main

import (
	"github.com/NoUmlautsAllowed/gocook/pkg/api"
	v2 "github.com/NoUmlautsAllowed/gocook/pkg/api/v2"
	"github.com/NoUmlautsAllowed/gocook/pkg/utils/tmpl"
	"github.com/NoUmlautsAllowed/gocook/pkg/view/recipe"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"iterateRange": tmpl.IterateRange[api.RecipeSearchResult],
	})
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	v := recipe.NewTemplateViewer(v2.NewV2Api())
	recipe.RegisterViewerRoutes(v, r)

	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}
