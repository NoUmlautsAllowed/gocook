package main

import (
	"github.com/NoUmlautsAllowed/gocook/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/splode/fname"
	"log"
	"net/http"
	"strconv"
	"time"
)

func generateRecipe(id string) (*api.Recipe, error) {

	rng := fname.NewGenerator(fname.WithDelimiter(" "))
	name, err := rng.Generate()

	if err != nil {
		return nil, err
	}

	return &api.Recipe{
		ID:                      id,
		Type:                    0,
		Title:                   name,
		Subtitle:                "",
		Owner:                   api.Owner{},
		Rating:                  api.Rating{},
		Difficulty:              0,
		HasImage:                false,
		HasVideo:                false,
		PreviewImageID:          id,
		PreviewImageOwner:       api.ImageOwner{},
		PreparationTime:         0,
		IsSubmitted:             false,
		IsRejected:              false,
		CreatedAt:               time.Time{},
		ImageCount:              0,
		Editor:                  api.Editor{},
		SubmissionDate:          time.Time{},
		IsPremium:               false,
		Status:                  0,
		Slug:                    "",
		PreviewImageURLTemplate: "",
		IsPlus:                  false,
		Servings:                0,
		KCalories:               0,
		Nutrition:               nil,
		Instructions:            "",
		MiscellaneousText:       "",
		IngredientsText:         "",
		Tags:                    nil,
		FullTags:                nil,
		ViewCount:               0,
		CookingTime:             0,
		RestingTime:             0,
		TotalTime:               0,
		IngredientGroups:        nil,
		CategoryIds:             nil,
		RecipeVideoID:           nil,
		IsIndexable:             false,
		AffiliateContent:        "",
		SiteURL:                 "",
	}, nil
}

func main() {
	r := gin.Default()

	r.GET("/v2/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	v2api := r.Group("/v2")

	recipe, err := generateRecipe(strconv.Itoa(0))
	if err != nil {
		log.Fatal(err)
	}
	recipeDb := []*api.Recipe{
		recipe,
	}

	v2api.GET("/recipes/:recipe", func(c *gin.Context) {
		rId := c.Param("recipe")
		for _, r := range recipeDb {
			if r.ID == rId {
				c.JSON(http.StatusOK, r)
				return
			}
		}

		// no recipe was found in db with matching id -> generate a new one
		r, err := generateRecipe(rId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.Error{
				Err:  err,
				Type: 0,
				Meta: nil,
			})
		}
		recipeDb = append(recipeDb, r)
		c.JSON(http.StatusOK, r)
	})

	v2api.GET("/search/recipe", func(c *gin.Context) {
		var search api.Search
		if err := c.Bind(&search); err == nil && len(search.Query) > 0 {
			recipeSearch := api.RecipeSearch{
				Count:   len(recipeDb),
				QueryId: "",
				Results: []api.RecipeSearchResult{},
			}

			for _, r := range recipeDb {
				recipeSearch.Results = append(recipeSearch.Results, api.RecipeSearchResult{
					Recipe: *r,
					Score:  0,
				})
			}

			c.JSON(http.StatusOK, recipeSearch)
		} else {
			c.JSON(http.StatusBadRequest, gin.Error{
				Err:  err,
				Type: 0,
				Meta: nil,
			})
		}
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no route",
		})
	})

	log.Fatal(r.Run(":8082"))
}
