package recipe

import (
	"net/url"
	"slices"
	"strconv"
	"strings"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"
)

type tmplTag struct {
	api.Tag
	URL string
}

type tmplTagGroup struct {
	api.TagGroup
	TagTemplates []tmplTag
}

func tagGroupTemplates(baseURL string, search api.RecipeSearch, values url.Values) []tmplTagGroup {
	tmpl := make([]tmplTagGroup, 0, len(search.TagGroups))

	if !strings.HasSuffix(baseURL, "?") {
		baseURL += "?"
	}

	currentTags := strings.Split(values.Get("tags"), ",")

	for _, group := range search.TagGroups {
		tmplGroup := tmplTagGroup{
			TagGroup:     group,
			TagTemplates: make([]tmplTag, 0, len(group.Tags)),
		}

		for _, tag := range group.Tags {
			activeTags := make([]string, len(currentTags), len(currentTags)+1)
			copy(activeTags, currentTags)
			if i := slices.Index(activeTags, strconv.Itoa(tag.ID)); i >= 0 && tag.IsActive {
				activeTags = append(activeTags[:i], activeTags[i+1:]...)
			} else {
				activeTags = append(activeTags, strconv.Itoa(tag.ID))
			}

			if len(activeTags) > 0 {
				values.Set("tags", strings.Trim(strings.Join(activeTags, ","), ","))
			} else {
				values.Del("tags")
			}

			tmplGroup.TagTemplates = append(tmplGroup.TagTemplates, tmplTag{
				Tag: tag,
				URL: baseURL + values.Encode(),
			})
		}

		tmpl = append(tmpl, tmplGroup)
	}

	return tmpl
}
