package api

type RecipeApi interface {
	Get(id string) (*Recipe, error)
	Search(s Search) (*RecipeSearch, error)
}
