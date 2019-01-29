package server

import(
	"testing"

	"github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/assert"
)

type routerServiceMock struct {
	mock.Mock
}

func (m *routerServiceMock)
func Test_retDummy(t *testing.T){
	dummy := retDummyStr();
	assert.Equal(t,"dummy",dummy)
}