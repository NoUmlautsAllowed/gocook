package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/NoUmlautsAllowed/gocook/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/splode/fname"
)

func generateRecipe(id string) (*api.Recipe, error) {
	rng := fname.NewGenerator(fname.WithDelimiter(" "), fname.WithCasing(fname.Title))
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
		CreatedAt:               time.Now(),
		ImageCount:              0,
		Editor:                  api.Editor{},
		SubmissionDate:          time.Now(),
		IsPremium:               false,
		Status:                  0,
		Slug:                    "",
		PreviewImageURLTemplate: id,
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
		CategoryIDs:             nil,
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

	recipeDB := []*api.Recipe{}

	for i := 0; i < 127; i++ {
		recipe, err := generateRecipe(strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		recipeDB = append(recipeDB, recipe)
	}

	v2api.GET("/recipes/:recipe", func(c *gin.Context) {
		rID := c.Param("recipe")
		for _, r := range recipeDB {
			if r.ID == rID {
				c.JSON(http.StatusOK, r)
				return
			}
		}

		// no recipe was found in db with matching id -> generate a new one
		r, err := generateRecipe(rID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.Error{
				Err:  err,
				Type: 0,
				Meta: nil,
			})
		}
		recipeDB = append(recipeDB, r)
		c.JSON(http.StatusOK, r)
	})

	v2api.GET("/search/recipe", func(c *gin.Context) {
		var search api.Search
		if err := c.Bind(&search); err == nil && len(search.Query) > 0 {
			recipeSearch := api.RecipeSearch{
				Count:   len(recipeDB),
				QueryID: "",
				Results: []api.RecipeSearchResult{},
			}

			limit, err := strconv.ParseInt(search.Limit, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.Error{
					Err:  err,
					Type: 0,
					Meta: errors.New("malformed limit value"),
				})
				return
			}
			offset, err := strconv.ParseInt(search.Offset, 10, 64)
			if err != nil {
				offset = 0
			}

			for i := range recipeDB {
				if i >= int(limit) || int64(i)+offset >= int64(len(recipeDB)) {
					break
				}
				recipeSearch.Results = append(recipeSearch.Results, api.RecipeSearchResult{
					Recipe: *recipeDB[int64(i)+offset],
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
