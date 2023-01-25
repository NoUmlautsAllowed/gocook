package v2

import (
	"encoding/json"
	"github.com/NoUmlautsAllowed/gocook/pkg/api"
	"io"
	"log"
	"net/url"
	"path"
)

func (a *V2Api) Get(id string) (*api.Recipe, error) {
	u, _ := url.Parse(a.baseRecipeUrl)
	u.Path = path.Join(u.Path, id)
	resp, err := a.defaultClient.Get(u.String())
	if err != nil {
		return nil, err
	} else {
		log.Println(resp.StatusCode, u)
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)

		var recipe api.Recipe
		if err = json.Unmarshal(data, &recipe); err != nil {
			return nil, err
		}

		recipe.PreviewImageURLTemplate = a.replaceImageCdnUrl(recipe.PreviewImageURLTemplate)
		return &recipe, nil
	}
}

func (a *V2Api) Search(s api.Search) (*api.RecipeSearch, error) {
	// this is how the api call looks like
	// https://api.chefkoch.de/v2/search/recipe?query=lasagne%20vegan
	// this is how the format should look like: crop-480x600, example
	// https://img.chefkoch-cdn.de/rezepte/2812481433250378/bilder/1185849/<format>/die-ultimative-vegane-lasagne.jpg
	// https://img.chefkoch-cdn.de/rezepte/2812481433250378/bilder/1185849/crop-480x600/die-ultimative-vegane-lasagne.jpg
	// this would be a query for pagination
	// https://api.chefkoch.de/v2/search-frontend/recipes?query=Lasagne+Vegan&limit=41&offset=41&analyticsTags=user,user_logged_out&enableClickAnalytics=true
	u, _ := url.Parse(a.baseSearchUrl)
	u.Path = path.Join(u.Path, "recipe")
	query := make(url.Values)
	query.Set("query", s.Query)
	query.Set("limit", s.Limit)
	query.Set("offset", s.Offset)
	u.RawQuery = query.Encode()

	resp, err := a.defaultClient.Get(u.String())
	if err != nil {
		return nil, err
	} else {
		log.Println(resp.StatusCode, u)
		defer resp.Body.Close()
		data, _ := io.ReadAll(resp.Body)
		var recipeSearch api.RecipeSearch
		err = json.Unmarshal(data, &recipeSearch)

		if err != nil {
			return nil, err
		}

		for i := range recipeSearch.Results {
			r := &recipeSearch.Results[i]
			r.Recipe.PreviewImageURLTemplate = a.replaceImageCdnUrl(r.Recipe.PreviewImageURLTemplate)
		}

		return &recipeSearch, nil
	}
}
