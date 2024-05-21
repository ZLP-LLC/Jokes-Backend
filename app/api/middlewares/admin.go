package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"jokes/constants"
	"jokes/domains"
	"jokes/lib"
	"jokes/models"
)

// RoleMiddleware middleware for role access
type RoleMiddleware struct {
	logger  lib.Logger
	service domains.UserService
	handler lib.RequestHandler
}

// NewRoleMiddleware creates new role middleware
func NewRoleMiddleware(
	logger lib.Logger,
	service domains.UserService,
	handler lib.RequestHandler,
) RoleMiddleware {
	return RoleMiddleware{
		logger:  logger,
		service: service,
		handler: handler,
	}
}

// Setup sets up admin middleware
func (m RoleMiddleware) Setup() {}

// Handler handles middleware functionality
func (m RoleMiddleware) Handler(roles []models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get(constants.UserID)
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		user, err := m.service.Get(userID.(uint))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		validRole := false
		for _, role := range roles {
			if user.Role == role {
				validRole = true
				break
			}
		}
		if !validRole {
			stringRoles := make([]string, len(roles))
			for i, role := range roles {
				stringRoles[i] = string(role)
			}
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Requires one of role: " + strings.Join(stringRoles, ", "),
			})
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
