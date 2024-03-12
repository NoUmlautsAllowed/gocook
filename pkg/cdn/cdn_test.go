package cdn

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"strings"
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	r := gin.Default()

	ctrl := gomock.NewController(t)
	m := NewMockImage(ctrl)

	RegisterRoutes(m, r)

	routes := r.Routes()

	if len(routes) != 2 {
		t.Error("expected 2 routes")
	}

	for _, route := range routes {
		if !strings.Contains(route.Path, RedirectImageCdnBasePath) {
			t.Error("expected " + RedirectImageCdnBasePath + " in route path")
		}
	}
}
