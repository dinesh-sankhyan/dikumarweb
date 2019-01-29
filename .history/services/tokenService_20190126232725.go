package services

import (
	"github.com/dikumarweb/logger"
	"crypto/rsa"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

///Remove or modify these lines of code and constant
const (
	privKeyPath = "/Users/dinesh_kumar/go/src/github.mheducation.com/MHEducation/dle-session-management/app/test/sample_key"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "/Users/dinesh_kumar/go/src/github.mheducation.com/MHEducation/dle-session-management/app/test/sample_key.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub

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

// read the key files before starting http handlers
func init() {
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
func CreateUserToken(userID string) (string, error) {&jwt.StandardClaims{
			Audience:  "aud",
			Id:        "4244bc8e-35f0-4612-8106-8ce16d5e15ca",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "https://idm-dev.mheducation.com",
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
