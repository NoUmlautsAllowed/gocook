package recipe

import (
	"fmt"
	"net/http"
	"net/url"

	"codeberg.org/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
)

type tmplComments struct {
	api.Comments
	api.CommentQuery
	Query          string
	ResultsPerPage int

	Pagination tmplPagination
}

const defaultCommentsPerPage int = 20

func (t *TemplateViewer) ShowComments(c *gin.Context) {
	var queryData api.CommentQuery
	if err := c.Bind(&queryData); err != nil {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: gin.ErrorTypeBind,
			Meta: nil,
		})
		return
	}

	queryData.RecipeID = c.Param("recipe")
	queryData.Limit = defaultCommentsPerPage
	if comments, err := t.api.Comments(queryData); err == nil {
		// remove all items that are replies to other comments
		commentsWithoutReplies := []api.CommentResult{}

		for _, c := range comments.Results {
			if c.ParentID == "" {
				commentsWithoutReplies = append(commentsWithoutReplies, c)
			}
		}

		comments.Results = commentsWithoutReplies

		tmplData := tmplComments{
			Comments:     *comments,
			CommentQuery: queryData,
			Query:        "",
			Pagination:   pagination(defaultCommentsPerPage, queryData.Offset, comments.Count, fmt.Sprintf("/recipes/%s/comments", queryData.RecipeID), make(url.Values)),
		}

		c.HTML(http.StatusOK, t.commentsTemplate, tmplData)
	} else {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: gin.ErrorTypeBind,
			Meta: nil,
		})
	}
}
