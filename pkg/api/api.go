package api

type RecipeAPI interface {
	Get(id string) (*Recipe, error)
	Comments(c CommentQuery) (*Comments, error)
	Search(s Search) (*RecipeSearch, error)
	Inspirations() (*RecipeInspirationsMixed, error)
}
