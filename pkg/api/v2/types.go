package v2

const ApiBaseRecipeUrl = "https://api.chefkoch.de/v2/recipes"
const ApiBaseSearchUrl = "https://api.chefkoch.de/v2/search"

type Owner struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Rank      int    `json:"rank"`
	Role      string `json:"role"`
	HasAvatar bool   `json:"hasAvatar"`
	HasPaid   bool   `json:"hasPaid"`
	Deleted   bool   `json:"deleted"`
}

type Rating struct {
	Rating   float64 `json:"rating"`
	NumVotes int     `json:"numVotes"`
}

type ImageOwner struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Rank      int    `json:"rank"`
	Role      string `json:"role"`
	HasAvatar bool   `json:"hasAvatar"`
	HasPaid   bool   `json:"hasPaid"`
	Deleted   bool   `json:"deleted"`
}

type Editor struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Rank      int    `json:"rank"`
	Role      string `json:"role"`
	HasAvatar bool   `json:"hasAvatar"`
	HasPaid   bool   `json:"hasPaid"`
	Deleted   bool   `json:"deleted"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Ingredient struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Unit         string      `json:"unit"`
	UnitID       string      `json:"unitId"`
	Amount       float64     `json:"amount"`
	IsBasic      bool        `json:"isBasic"`
	UsageInfo    string      `json:"usageInfo"`
	URL          interface{} `json:"url"`
	FoodID       string      `json:"foodId"`
	ProductGroup string      `json:"productGroup"`
	BlsKey       string      `json:"blsKey"`
}

type IngredientGroup struct {
	Header      string       `json:"header"`
	Ingredients []Ingredient `json:"ingredients"`
}
