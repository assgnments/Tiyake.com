package token

import (
	"github.com/dgrijalva/jwt-go"
)

// CustomClaims specifies custom claims
type CustomClaims struct {
	SessionId string `json:"sessionId"`
	jwt.StandardClaims
}

//NewClaims returns claims used for generating jwt tokens
func NewClaims(sessionId string,expire int64) jwt.Claims {
	return CustomClaims{
		sessionId,
		jwt.StandardClaims{
			ExpiresAt: expire,
		},
	}
}

// Generates jwt token
func Generate(signingKey []byte, claims jwt.Claims) (string, error) {
	tn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := tn.SignedString(signingKey)
	return signedString, err
}


///Validate, parse, returns claim
func GetSessionIdFromToken(tkn string,keyFunc jwt.Keyfunc) string{
	token,err:=jwt.Parse(tkn,keyFunc)
	if err!=nil{
		return ""
	}
	sessionId:= token.Claims.(jwt.MapClaims)["sessionId"].(string)
	if !token.Valid {
		return ""
	}
	return sessionId
}

// CSRFToken Generates random string for CSRF
func NewCSRFToken(signingKey []byte) (string, error) {
	tn := jwt.New(jwt.SigningMethodHS256)
	signedString, err := tn.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// ValidCSRF checks if a given csrf token is valid
func ISValidCSRF(signedToken string, signingKey []byte) bool {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		return false
	}

	return true
}
