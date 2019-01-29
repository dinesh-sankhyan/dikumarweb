package server

import(
	"testing"

	"github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/assert"
)

type routerServiceMock struct {
	mock.Mock
}

func (m *routerServiceMock) AddValue(a,b int) int
func Test_retDummy(t *testing.T){
	dummy := retDummyStr();
	assert.Equal(t,"dummy",dummy)
}