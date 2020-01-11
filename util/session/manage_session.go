package session

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"teyake/entity"
	"teyake/util/token"
	"time"
)


const SessionKey="session_key"

func NewSession(id uint) *entity.Session {
	//Expires after a month for debugging
	tokenExpires := time.Now().AddDate(0,1,0).Unix()
	signingString, err := token.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: []byte(signingString),
		UUID:       id,
	}
}

// Create creates and sets sessionId cookie and sessionValueCookie
func SetCookies(claims jwt.Claims, sessionID uint, signingKey []byte, w http.ResponseWriter) {
	signedString, err := token.GenerateJWTClaim(signingKey, claims)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	sessionIdCookie := http.Cookie{
		Name:     SessionKey,
		Value:    string(sessionID),
		HttpOnly: true,
	}
	sessionValueCookie:=http.Cookie{
		Name:  string(sessionID),
		Value: signedString,
		HttpOnly:true,
	}

	http.SetCookie(w, &sessionIdCookie)
	http.SetCookie(w, &sessionValueCookie)
}


// Remove expires existing cookies
func RemoveCookies(sessionID uint, w http.ResponseWriter) {
	sessionIdCookie := http.Cookie{
		Name:     SessionKey,
		Value:   "",
		Expires: time.Unix(1, 0),
		MaxAge:  -1,
	}
	sessionValueCookie:=http.Cookie{
		Name:  string(sessionID),
		Value:   "",
		HttpOnly:true,
		Expires: time.Unix(1, 0),
		MaxAge:  -1,
	}

	http.SetCookie(w, &sessionIdCookie)
	http.SetCookie(w, &sessionValueCookie)
}
