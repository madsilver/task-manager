package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/madsilver/task-manager/internal/adapter/presenter"
	"net/http"
)

func HasRole(rolesRequired ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesVal, exists := c.Get("roles")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, presenter.ErrorResponse{Error: "roles not found in context"})
			return
		}

		roles, ok := rolesVal.([]interface{})
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, presenter.ErrorResponse{Error: "invalid roles format"})
			return
		}

		for _, required := range rolesRequired {
			for _, r := range roles {
				if r == required {
					c.Next()
					return
				}
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, presenter.ErrorResponse{Error: "insufficient role"})
	}
}
