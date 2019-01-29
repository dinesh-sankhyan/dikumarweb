package server

import(
	"testing"

    "github.com/stretchr/testify/assert"
)

type smsServiceMock struct {
	mock.Mock
}
func Test_retDummy(t *testing.T){
	dummy := retDummyStr();
	assert.Equal(t,"dummy",dummy)
}