package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// CustomClaims specifies custom claims
type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
// Claims returns claims used for generating jwt tokens
//Expires after a month for debugging
func Claims(email string,) jwt.Claims {
	return CustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0,1,0).Unix(),
		},
	}
}

// Generate generates jwt token
func GenerateJWTClaim(signingKey []byte, claims jwt.Claims) (string, error) {
	tn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := tn.SignedString(signingKey)
	return signedString, err
}

// Valid validates a given token
func ValidateToken(signedToken string, signingKey []byte) (bool, error) {
	token, err := jwt.ParseWithClaims(signedToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(*CustomClaims); !ok || !token.Valid {
		return false, err
	}

	return true, nil
}



// CSRFToken Generates random string for CSRF
func CSRFToken(signingKey []byte) (string, error) {
	tn := jwt.New(jwt.SigningMethodHS256)
	signedString, err := tn.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// ValidCSRF checks if a given csrf token is valid
func ValidCSRF(signedToken string, signingKey []byte) (bool, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		return false, err
	}

	return true, nil
}
