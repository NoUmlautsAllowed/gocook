package main

import (
	v2 "chefcook/pkg/api/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})
	r.GET("/recipes", v2.SearchRecipes)
	r.GET("/recipes/:recipe", v2.GetRecipe)
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080
}
