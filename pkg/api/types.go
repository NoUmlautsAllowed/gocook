package api

import "time"

// API Type definitions
// These are the main API type definitions needed to query and search recipes.
// As for now, only API v2 is implemented, these types represent the API v2 types.
// This can change any time, so the real API types may need implementation in the API vX package.
// For easy type definitions, the JSON given by the API may be pasted into https://transform.tools/json-to-go

type Owner struct {
	ID                     string `json:"id"`
	Username               string `json:"username"`
	Rank                   int    `json:"rank"`
	Role                   string `json:"role"`
	HasAvatar              bool   `json:"hasAvatar"`
	AvatarImageURLTemplate string `json:"avatarImageUrlTemplate"`
	HasPaid                bool   `json:"hasPaid"`
	Deleted                bool   `json:"deleted"`
}

type Rating struct {
	Rating   float64 `json:"rating"`
	NumVotes int     `json:"numVotes"`
}

type Tag struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Count      int    `json:"count"`
	IsActive   bool   `json:"isActive"`
	IsDisabled bool   `json:"isDisabled"`
}

type Ingredient struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Unit         string  `json:"unit"`
	UnitID       string  `json:"unitId"`
	Amount       float64 `json:"amount"`
	IsBasic      bool    `json:"isBasic"`
	UsageInfo    string  `json:"usageInfo"`
	URL          any     `json:"url"`
	FoodID       string  `json:"foodId"`
	ProductGroup string  `json:"productGroup"`
	BlsKey       string  `json:"blsKey"`
}

type IngredientGroup struct {
	Header      string       `json:"header"`
	Ingredients []Ingredient `json:"ingredients"`
}

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
	PreviewImageOwner       Owner             `json:"previewImageOwner"`
	PreparationTime         int               `json:"preparationTime"`
	IsSubmitted             bool              `json:"isSubmitted"`
	IsRejected              bool              `json:"isRejected"`
	CreatedAt               time.Time         `json:"createdAt"`
	ImageCount              int               `json:"imageCount"`
	Editor                  Owner             `json:"editor"`
	SubmissionDate          time.Time         `json:"submissionDate"`
	IsPremium               bool              `json:"isPremium"`
	Status                  int               `json:"status"`
	Slug                    string            `json:"slug"`
	PreviewImageURLTemplate string            `json:"previewImageUrlTemplate"`
	IsPlus                  bool              `json:"isPlus"`
	Servings                int               `json:"servings"`
	KCalories               int               `json:"kCalories"`
	Nutrition               any               `json:"nutrition"`
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
	CategoryIDs             []string          `json:"categoryIds"`
	RecipeVideoID           any               `json:"recipeVideoId"`
	IsIndexable             bool              `json:"isIndexable"`
	AffiliateContent        string            `json:"affiliateContent"`
	SiteURL                 string            `json:"siteUrl"`
}

type RecipeSearchResult struct {
	Recipe Recipe `json:"recipe"`
	Score  int    `json:"score"`
}

type RecipeSearch struct {
	Count     int                  `json:"count"`
	QueryID   string               `json:"queryId"`
	Results   []RecipeSearchResult `json:"results"`
	TagGroups []TagGroup           `json:"tagGroups"`
}

type TagGroup struct {
	Key        string `json:"key"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	IsActive   bool   `json:"isActive"`
	IsDisabled bool   `json:"isDisabled"`
	Tags       []Tag  `json:"tags"`
}

type Search struct {
	Query  string `form:"query"`
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	Tags   string `form:"tags"`
}

type CommentResult struct {
	ID           string    `json:"id"`
	Owner        Owner     `json:"owner"`
	CreatedAt    time.Time `json:"createdAt"`
	Text         string    `json:"text"`
	Helpful      bool      `json:"helpful"`
	HelpfulCount int       `json:"helpfulCount"`
	ParentID     string    `json:"parentId"`
	Children     any       `json:"children"`
}

type Comments struct {
	Count         int             `json:"count"`
	Results       []CommentResult `json:"results"`
	MaxAge        int             `json:"maxAge"`
	SharedMaxAge  int             `json:"sharedMaxAge"`
	SurrogateKeys []string        `json:"surrogateKeys"`
}

type CommentQuery struct {
	RecipeID   string `form:"recipeId"`
	RecipeName string `form:"recipeName"`
	Limit      int    `form:"limit"`
	Offset     int    `form:"offset"`
}

type RecipeInspirations struct {
	Recipes []Recipe `json:"recipes"`
}

type RecipeInspirationsMixed struct {
	CookingRecipes []Recipe
	BakingRecipes  []Recipe
}
