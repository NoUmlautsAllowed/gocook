package main

import (
	"chefcook/pkg/api"
	v2 "chefcook/pkg/api/v2"
	"chefcook/pkg/utils/tmpl"
	"chefcook/pkg/view/recipe"
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
