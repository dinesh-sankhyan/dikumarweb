package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"encoding/json"
	"strings"

	"github.com/stretchr/testify/assert"
	"github.mheducation.com/MHEducation/dle-planner-api/config"
	"github.mheducation.com/MHEducation/dle-planner-api/domain"
)

func testServer() *Server {
	config.InitConfig()
	return New()
}

func mockHTTPServer(t *testing.T, h http.Handler, req *http.Request) (*http.Response, []byte) {
	assert := assert.New(t)
	ts := httptest.NewServer(h)
	defer ts.Close()

	// Add the started server's URL to the front of the original request URL
	u, err := url.Parse(ts.URL + req.URL.String())
	assert.NoError(err, "Invalid URL: %s", req.URL.String())
	req.URL = u

	// Make the request without following redirects
	client := http.DefaultClient
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	resp, err := client.Do(req)
	assert.NoError(err, "Error making HTTP request")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err, "Error reading response body")

	return resp, body
}
