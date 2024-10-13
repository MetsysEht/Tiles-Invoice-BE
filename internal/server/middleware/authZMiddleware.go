package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// Authorize determines if current subject has been authorized to take an action on an object.
func AuthzMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		u, ok := c.Get("role")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		role, _ := u.(string)
		// casbin enforces policy
		ok, err := enforce(role, c.Request.RequestURI, c.Request.Method)
		//ok, err := enforce(val.(string), obj, act, adapter)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(500, "error occurred when authorizing user")
			return
		}
		if !ok {
			c.AbortWithStatusJSON(403, "forbidden")
			return
		}
		c.Next()
	}
}

func enforce(sub string, obj string, act string) (bool, error) {
	enforcer, err := casbin.NewEnforcer("internal/server/middleware/authz_model.conf", "config/authz_policy.csv")
	if err != nil {
		return false, fmt.Errorf("failed to load policy: %w", err)
	}
	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, fmt.Errorf("failed to enforce policy: %w", err)
	}
	return ok, nil
}
