package middleware

import (
	"fmt"
	"net/http"
	"sayamphoo/microservice/enum"
	"sayamphoo/microservice/models/domain"
	"sayamphoo/microservice/repository"
	"sayamphoo/microservice/security/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

var bypassPath = []string{
	"/login",
	"/register",
	"/image",
}

func byPass(c *gin.Context) bool {
	path := c.Request.URL.Path

	for _, s := range bypassPath {
		if strings.Contains(path, s) {
			return true
		}
	}

	return false
}

func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bypass logic
		if byPass(c) {
			return
		}

		headers := c.Request.Header

		// Print all headers
		fmt.Println("\n\n------------- START HEARDER ----------\n")
		for key, values := range headers {
			fmt.Printf("%s: %s\n", key, values)
		}
		fmt.Println("\n------------- END HEARDER ----------\n\n")

		// Check Authorization header
		token := c.GetHeader("Authorization")

		if !(isValidBearerToken(token)) {
			unauthorizedError()
		}

		// Verify JWT token
		token = strings.Trim(token[7:], " ")
		result, err := jwt.VerifyToken(token)
		if err != nil {
			unauthorizedError()
		}

		// Extract user ID from token claims
		Id, err := result.Claims.GetSubject()
		if err != nil {
			unauthorizedError()
		}

		// Check access based on user role
		checkAccess(Id, c.Request.URL.Path, c.Request.Method)

		// Set user ID in the context
		c.Set(enum.REQUEST_USER_ID, Id)
		c.Next()
	}
}

func isValidBearerToken(token string) bool {
	return len(token) >= 7 && strings.ToUpper(token[0:6]) == "BEARER"
}

func unauthorizedError() {
	panic(domain.UtilityModel{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	})
}

func checkAccess(userId string, path string, method string) {
	pathSegments := strings.Split(path, "/")

	if userId == "owner" {
		if pathSegments[1] == "api-customer" && method != "GET" {
			accessDeniedError()
		}
	} else {
		if pathSegments[1] == "api-customer" {
			repoMember := &repository.MemberRepo{}
			member, err := repoMember.FindById(userId)
			if err != nil || member == nil {
				memberNotFoundError()
			}
		} else {
			accessDeniedError()
		}
	}
}

func accessDeniedError() {
	panic(domain.UtilityModel{
		Code:    http.StatusUnauthorized,
		Message: "Access to this path is denied for the System.",
	})
}

func memberNotFoundError() {
	panic(domain.UtilityModel{
		Code:    http.StatusNotFound,
		Message: "Member Not Found",
	})
}
