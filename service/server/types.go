package server

import (
	"github.com/dgrijalva/jwt-go"
	"RSS_FEEDS/pkg/rss"
)

// RSSUrls used to decode request body from rss
type RSSUrls struct {
	Urls []string `json:"rss_urls"`
}

// RSSItems used to encode response for rss
type RSSItems struct {
	Items []rss.RssItem `json:"items"`
}


// Credentials holds username and pasword
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims is a struct that will be encoded to a JWT.
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}


// JWT key used to create the signature
// Obviously this should be in cfg file and should not be exposed
var jwtKey = []byte("rss_feed")

// instead of database use map with dummy user accounts
var users = map[string]string{
	"user": "password",
}