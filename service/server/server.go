package server

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"log"

	"github.com/dgrijalva/jwt-go"
	"RSS_FEEDS/pkg/rss"
)

const (
	// ERRORS
	internalServerError = "Internal Server Error!"
	incorrectCredentials = "Incorrect Credentials!"
	badRequest = "Bad Request!"
	tokenValid = "Token is still valid!"

	// set expiration of token to 1 hour
	expiration = time.Hour;
	// time before expiration when u can renew
	renewBeforeExp = 5 * time.Minute
)

func Run() {
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/renew", func(w http.ResponseWriter, r *http.Request) {
		claim, err := jwtAuth(w, r)
		if err != nil {
			return
		}
		renewHandler(w, r, claim)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		claim, err := jwtAuth(w, r)
		if err != nil {
			return
		}
		homePageHandler(w, r, claim)
	})

	http.HandleFunc("/rss", func(w http.ResponseWriter, r *http.Request) {
		claim, err := jwtAuth(w, r)
		if err != nil {
			return
		}
		rssHandler(w, r, claim)
	})

	log.Println("Server starting ...")
	defer log.Println("Server shutting down ...")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("Server failed with error: ", err)
		return
	}
}

// signupHandler gets Jwt token
func signupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint: Signup")

	creds := &Credentials{}

	// Decode
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badRequest))
		return
	}

	// Check validity of credentials
	if expectedPassword, ok := users[creds.Username]; !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(incorrectCredentials))
		return
	}

	// set expiration of token
	expirationTime := time.Now().Add(expiration)

	// Declare the token with HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims {
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expirationTime.Unix(),
		},
	})

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(internalServerError))
		return
	}

	// Set Cookies with token and expiration time
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	log.Printf("User: %s with token: %s", creds.Username, tokenString)
}

// jwtAuth checks authentication via jwt token
func jwtAuth(w http.ResponseWriter, r *http.Request) (*Claims, error) {
	log.Println("JWT AUTHENTICATION")

	claims := &Claims{}

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(incorrectCredentials))
			return claims, err
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badRequest))
		return claims, err
	}

	tknStr := c.Value

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(incorrectCredentials))
			return claims, err
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badRequest))
		return claims, err
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(incorrectCredentials))
		return claims, err
	}

	return claims, nil

}

// renewHandler updates jwt token after are close to its expiration
func renewHandler(w http.ResponseWriter, r *http.Request, c *Claims) {
	log.Println("Endpoint: Renew")

	// Refresh only in the last 5 minutes
	if time.Unix(c.ExpiresAt, 0).Sub(time.Now()) > renewBeforeExp {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tokenValid))
		return
	}

	// Refresh token and expiration time
	expirationTime := time.Now().Add(renewBeforeExp)
	c.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(internalServerError))
		return
	}

	// Set Cookies with new token and expiration time
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

// homePageHandler
func homePageHandler(w http.ResponseWriter, r *http.Request, c *Claims) {
	log.Println("Endpoint: homePage")

	w.Write([]byte(fmt.Sprintf("Welcome %s!", c.Username)))
}

// rssHandler
func rssHandler(w http.ResponseWriter, r *http.Request, c *Claims) {
	log.Println("Endpoint: RSS")

	rfs := &RSSUrls{}
	err := json.NewDecoder(r.Body).Decode(rfs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(badRequest))
		return
	}

	ris := &RSSItems{
		rss.Parse(rfs.Urls),
	}

	json.NewEncoder(w).Encode(ris)
}

