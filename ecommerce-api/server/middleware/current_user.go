package middleware

import (
	"ecommerce-api/commons"
	"ecommerce-api/modules/user/payload"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func ValidateJWT(opt commons.Model) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")

			if auth == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Bearer Token is required")
			}

			splittedTokens := strings.Split(auth, " ")
			bearerPart := splittedTokens[0]
			if len(splittedTokens) != 2 || strings.ToLower(bearerPart) != "bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Bearer Token is invalid")
			}

			userToken := splittedTokens[1]
			claimsUser := payload.TokenPayload{}
			claims, err := jwt.ParseWithClaims(userToken, &claimsUser, func(token *jwt.Token) (interface{}, error) {
				return []byte(opt.Config.JwtCfg.Secret), nil
			})

			if !claims.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is invalid")
			}

			if errors.Is(err, jwt.ErrSignatureInvalid) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is invalid")
			}

			if errors.Is(err, jwt.ErrTokenExpired) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is expired")
			}

			// Get user by token
			u, _ := opt.Service.User.Get(c.Request().Context(), payload.UserGet{
				Token: &userToken,
			})

			if u == nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is invalid")
			}

			c.Set("user", &claimsUser)
			return next(c)
		}
	}
}

func CurrentUser(c echo.Context) *payload.TokenPayload {
	user, ok := c.Get("user").(*payload.TokenPayload)

	if !ok {
		return new(payload.TokenPayload)
	}

	return user
}
