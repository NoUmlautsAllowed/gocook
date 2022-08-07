package v2

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type Recipe struct {
	ID                      string            `json:"id"`
	Type                    int               `json:"type"`
	Title                   string            `json:"title" form:"title"`
	Subtitle                string            `json:"subtitle"`
	Owner                   Owner             `json:"owner"`
	Rating                  Rating            `json:"rating"`
	Difficulty              int               `json:"difficulty"`
	HasImage                bool              `json:"hasImage"`
	HasVideo                bool              `json:"hasVideo"`
	PreviewImageID          string            `json:"previewImageId"`
	PreviewImageOwner       ImageOwner        `json:"previewImageOwner"`
	PreparationTime         int               `json:"preparationTime"`
	IsSubmitted             bool              `json:"isSubmitted"`
	IsRejected              bool              `json:"isRejected"`
	CreatedAt               time.Time         `json:"createdAt"`
	ImageCount              int               `json:"imageCount"`
	Editor                  Editor            `json:"editor"`
	SubmissionDate          time.Time         `json:"submissionDate"`
	IsPremium               bool              `json:"isPremium"`
	Status                  int               `json:"status"`
	Slug                    string            `json:"slug"`
	PreviewImageURLTemplate string            `json:"previewImageUrlTemplate"`
	IsPlus                  bool              `json:"isPlus"`
	Servings                int               `json:"servings"`
	KCalories               int               `json:"kCalories"`
	Nutrition               interface{}       `json:"nutrition"`
	Instructions            string            `json:"instructions"`
	MiscellaneousText       string            `json:"miscellaneousText"`
	IngredientsText         string            `json:"ingredientsText"`
	Tags                    []string          `json:"tags"`
	FullTags                []Tag             `json:"fullTags"`
	ViewCount               int               `json:"viewCount"`
	CookingTime             int               `json:"cookingTime"`
	RestingTime             int               `json:"restingTime"`
	TotalTime               int               `json:"totalTime"`
	IngredientGroups        []IngredientGroup `json:"ingredientGroups"`
	CategoryIds             []string          `json:"categoryIds"`
	RecipeVideoID           interface{}       `json:"recipeVideoId"`
	IsIndexable             bool              `json:"isIndexable"`
	AffiliateContent        string            `json:"affiliateContent"`
	SiteURL                 string            `json:"siteUrl"`
}

type RecipeSearchResult struct {
	Recipe Recipe `json:"recipe"`
	Score  int    `json:"score"`
}

type RecipeSearch struct {
	Count   int                  `json:"count"`
	QueryId string               `json:"queryId"`
	Results []RecipeSearchResult `json:"results"`
}

type tmplRecipe struct {
	Title            string            `json:"title"`
	Subtitle         string            `json:"subtitle"`
	IngredientGroups []IngredientGroup `json:"ingredientGroups"`
	Instructions     []string          `json:"instructions"`
}

func GetRecipe(c *gin.Context) {
	recipe := c.Param("recipe")
	u, _ := url.Parse(ApiBaseRecipeUrl)
	u.Path = path.Join(u.Path, recipe)
	resp, err := http.Get(u.String())
	log.Println(resp.StatusCode, u)
	if err != nil {
		c.JSON(500, gin.Error{
			Err:  err,
			Type: 0,
			Meta: nil,
		})
	} else {
		data, _ := io.ReadAll(resp.Body)
		var recipe Recipe
		_ = json.Unmarshal(data, &recipe)
		tmpl := tmplRecipe{
			Title:            recipe.Title,
			Subtitle:         recipe.Subtitle,
			IngredientGroups: recipe.IngredientGroups,
			Instructions:     strings.Split(recipe.Instructions, "\n\n"),
		}
		c.HTML(http.StatusOK, "recipe.tmpl", tmpl)
	}
}

func SearchRecipes(c *gin.Context) {
	type Search struct {
		Query string `form:"query"`
	}
	var search Search
	if c.ShouldBind(&search) == nil {
		// this is how the api call looks like
		// https://api.chefkoch.de/v2/search/recipe?query=lasagne%20vegan
		u, _ := url.Parse(ApiBaseSearchUrl)
		u.Path = path.Join(u.Path, "recipe")
		query := make(url.Values)
		query.Set("query", search.Query)
		u.RawQuery = query.Encode()

		resp, err := http.Get(u.String())
		log.Println(resp.StatusCode, u)
		if err != nil {
			c.JSON(500, gin.Error{
				Err:  err,
				Type: 0,
				Meta: nil,
			})
		} else {
			data, _ := io.ReadAll(resp.Body)
			var recipeSearch RecipeSearch
			_ = json.Unmarshal(data, &recipeSearch)
			c.JSON(http.StatusOK, recipeSearch)
		}
	}
}
