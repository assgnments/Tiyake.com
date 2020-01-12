package session

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"teyake/entity"
	"teyake/util/token"
	"time"
)

const SessionKey = "session_key"

func NewSession(id uint) *entity.Session {
	//Expires after a month for debugging
	tokenExpires := time.Now().AddDate(0, 1, 0).Unix()
	signingString, err := token.GenerateRandomString(32)
	sessionId, err := token.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	return &entity.Session{
		SessionId:  sessionId,
		Expires:    tokenExpires,
		SigningKey: []byte(signingString),
		UUID:       id,
	}
}

// Create creates and sets sessionId cookie and sessionValueCookie
func SetCookies(claims jwt.Claims, expire int64, signingKey []byte, w http.ResponseWriter) {
	signedString, err := token.Generate(signingKey, claims)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	cookie := http.Cookie{
		Name:     SessionKey,
		Value:    signedString,
		HttpOnly: true,
		Expires:  time.Unix(expire, 0),
	}

	http.SetCookie(w, &cookie)
}

// Remove expires existing session
func RemoveCookies(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:    SessionKey,
		Value:   "",
		Expires: time.Unix(1, 0),
		MaxAge:  -1,
	}
	http.SetCookie(w, &cookie)
}
