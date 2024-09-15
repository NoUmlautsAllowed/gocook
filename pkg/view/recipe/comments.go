package recipe

import (
	"math"
	"net/http"

	"github.com/NoUmlautsAllowed/gocook/pkg/api"

	"github.com/gin-gonic/gin"
)

type tmplCommentData struct {
	Offset int
	Exists bool
}

type tmplComments struct {
	api.Comments
	api.CommentQuery
	Query          string
	ResultsPerPage int

	Previous tmplCommentData
	Next     tmplCommentData
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
			Previous: tmplCommentData{
				Offset: int(math.Max(0, float64(queryData.Offset-defaultCommentsPerPage))),
				Exists: queryData.Offset > 0,
			},
			Next: tmplCommentData{
				Offset: queryData.Offset + defaultCommentsPerPage,
				Exists: queryData.Offset+defaultCommentsPerPage < comments.Count,
			},
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
