package services

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dikumarweb/logger"
	"gopkg.in/jose.v1/crypto"
	"gopkg.in/jose.v1/jws"
)

const (
	privKeyPath = "/Users/dinesh_kumar/go/src/github.com/dikumarweb/keys/sample_key.key" // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "/Users/dinesh_kumar/go/src/github.com/dikumarweb/keys/sample_key.key.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub

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

//JwtValidateToken validate jwt token
func JwtValidateToken(r *http.Request) (isValidToken bool,  err error) {
	token, err := jws.ParseJWTFromRequest(r)
	if err == nil {
		if err = token.Validate(verifyKey, crypto.SigningMethodRS256); err != nil &&
			(err.Error() == "token is expired" || err.Error() == "token is not yet valid") {
			if err = token.Claims().Validate(time.Now(), 120*time.Second, 120*time.Second); err != nil {
				logger.Errorf("Token is not valid even after adding leeway %s", err)
				return
			}
		}
		isValidToken = true
		logger.Info("token decoded successfully")

	} else {

		logger.Error("Unauthorized access to this resource " + err.Error())
	}
	return
}
