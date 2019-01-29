package services

import (
	"github.com/dikumarweb/logger"
	"crypto/rsa"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "/Users/dinesh_kumar/go/src/github.com/dikumarweb/keys/sample_key.key"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "/Users/dinesh_kumar/go/src/github.com/dikumarweb/keys/sample_key.key" // openssl rsa -in app.rsa -pubout > app.rsa.pub

)

var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
	serverPort int
	// storing sample username/password pairs
	// don't do this on a real server
	users = map[string]string{
		"test": "known",
	}
)

//InitService read the key files before starting http handlers
func InitTokenService() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		logger.Info(err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		logger.Info(err)
	}
	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		logger.Info(err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		logger.Info(err)
	}

}

//CreateUserToken create user token
func CreateUserToken() (string, error) {
	token := &jwt.StandardClaims{
			Audience:  "aud",
			Id:        "4244bc8e-35f0-4612-8106-8ce16d5e15ca",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "https://dev.testexample.com",
			NotBefore: time.Now().Unix(),
			Subject:   "sub",
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		}
	return getSignedUserClaim(token)

}

// Remove Block Ended

func getSignedUserClaim(token *jwt.StandardClaims) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims = token

	// Creat token string
	return t.SignedString(signKey)
}

// GetValidTokenFromRequest parses a request for a JWT, validates it, and converts it to a common user JWT format
func GetValidTokenFromRequest(r *http.Request) (*jwt.StandardClaims, error) {
	// Get token from header
	rawToken := GetRawTokenFromRequest(r)
	if rawToken == "" {
		return nil, errors.New("jwt not found in header")
	}

	return ParseRawToken(rawToken)
}