package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
	"github.com/go-loremipsum/loremipsum"
	"github.com/splode/fname"
)

func generateRecipe(id int) (*api.Recipe, error) {
	rng := fname.NewGenerator(fname.WithDelimiter(" "), fname.WithCasing(fname.Title))
	lig := loremipsum.New()
	name, err := rng.Generate()
	ids := strconv.Itoa(id)
	if err != nil {
		return nil, err
	}

	return &api.Recipe{
		ID:                      ids,
		Type:                    0,
		Title:                   name,
		Subtitle:                "",
		Owner:                   api.Owner{},
		Rating:                  api.Rating{},
		Difficulty:              0,
		HasImage:                id%2 == 0,
		HasVideo:                false,
		PreviewImageID:          ids,
		PreviewImageOwner:       api.Owner{},
		PreparationTime:         0,
		IsSubmitted:             false,
		IsRejected:              false,
		CreatedAt:               time.Now(),
		ImageCount:              0,
		Editor:                  api.Owner{},
		SubmissionDate:          time.Now(),
		IsPremium:               false,
		Status:                  0,
		Slug:                    "",
		PreviewImageURLTemplate: ids,
		IsPlus:                  false,
		Servings:                0,
		KCalories:               0,
		Nutrition:               nil,
		Instructions:            lig.Paragraph(),
		MiscellaneousText:       "",
		IngredientsText:         "",
		Tags:                    nil,
		FullTags:                nil,
		ViewCount:               0,
		CookingTime:             0,
		RestingTime:             0,
		TotalTime:               0,
		IngredientGroups: []api.IngredientGroup{
			{
				Header: lig.Words(4),
				Ingredients: []api.Ingredient{
					{Name: lig.Words(2), Amount: 1},
					{Name: lig.Words(2), Amount: 1},
					{Name: lig.Words(2), Amount: 1},
				},
			},
			{
				Header: lig.Words(4),
				Ingredients: []api.Ingredient{
					{Name: lig.Words(2), Amount: 1},
					{Name: lig.Words(2), Amount: 1},
					{Name: lig.Words(2), Amount: 1},
				},
			},
		},
		CategoryIDs:      nil,
		RecipeVideoID:    nil,
		IsIndexable:      false,
		AffiliateContent: "",
		SiteURL:          "",
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
		recipe, err := generateRecipe(i)
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
		id, err := strconv.Atoi(rID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.Error{
				Err:  err,
				Type: 0,
				Meta: nil,
			})
		}
		r, err := generateRecipe(id)
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
