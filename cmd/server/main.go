package main

import (
	v2 "chefcook/pkg/api/v2"
	"chefcook/pkg/utils/tmpl"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"iterateRange": tmpl.IterateRange[v2.RecipeSearchResult],
	})
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})
	r.GET("/recipes", v2.SearchRecipes)
	r.GET("/recipes/:recipe", v2.GetRecipe)
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}
