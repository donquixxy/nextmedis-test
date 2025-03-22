package payload

import "github.com/golang-jwt/jwt/v5"

type (
	LoginResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}

	TokenPayload struct {
		// user id
		ID string `json:"id"`
		jwt.RegisteredClaims
	}
)
