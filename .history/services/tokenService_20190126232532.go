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
func CreateUserToken(userID string) (token *jwt.StandardClaims, err error) {
	scope := []string{"test1", "test2"}
	token = &models.CustomJWT{
		&jwt.StandardClaims{
			Audience:  "aud",
			Id:        "4244bc8e-35f0-4612-8106-8ce16d5e15ca",
			IssuedAt:  time.Now().Unix(),
			Issuer:    "https://idm-dev.mheducation.com",
			NotBefore: time.Now().Unix(),
			Subject:   "sub",
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"",
		"",           //Person Xid
		userID,   //Username
		"urn:com.mheducation.openlearning:enterprise.identity:qastg.us-east-1:session:1ef87379-a4bd-462b-b8d2-5bb296f0b7c1", //Session XID
		"1ef87379-a4bd-462b-b8d2-5bb296f0b7c1",  //Session XID
		"urn:com.mheducation.openlearning:lms:qastg.us-east-1.connect2-qastg.mheducation.com",  //Client XID
		"urn:com.mheducation.openlearning:enterprise.identity:qastg.global:person:42f951f4-67dc-11e7-bc26-0e11a86e63f4",        //XID
		scope,        //Scope
		"EngradeStageClient",        //cid
	}
	accessToken, err := getSignedUserClaim(token)
	token.Token = accessToken

	return
}

// Remove Block Ended

func getSignedUserClaim(token *models.CustomJWT) (string, error) {
	// create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims = token

	// Creat token string
	return t.SignedString(signKey)
}
