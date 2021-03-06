package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Using asymetric crypto/RSA keys
// location of private/publice key files
const (
	// openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	// openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

// Private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

// Read the key files before starting http handlers
func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}
}

// Generate JWT token
func GenerateJWT(name, role string) (string, error) {
	// Set claims for JWT token
	claims := jwt.MapClaims{}
	claims["iss"] = "admin"
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	// Create a signer for rsa 256
	// t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Extract JWT from request header
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	tokenFromRequest := ExtractToken(r)
	// Validate the token
	token, err := jwt.Parse(tokenFromRequest, func(token *jwt.Token) (interface{}, error) {
		// Verify the token with public key, which is the counter part of private key
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: // JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Token!",
					401,
				)
				return
			
			default:
				DisplayAppError(
					w,
					err,
					"Error while parsing the Access Token!",
					500,
				)
				return
			}

		default:
			DisplayAppError(
				w,
				err,
				"Error while parsing Access Token!",
				500,
			)
			return
		}
	}

	if token.Valid {
		next(w, r)
	} else {
		DisplayAppError(
			w,
			err,
			"Invalid Access Token",
			401,
		)
	}
}
