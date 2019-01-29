package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type calcServiceMock struct {
	mock.Mock
}

func (m *routerServiceMock) AddValue(a, b int) int {
	fmt.Println("AddValue function")
	fmt.Printf("Value passed in: %d and %d", a,b)
	// this records that the method was called and passes in the value
	// it was called with
	args := m.Called(a,b)
	// it then returns whatever we tell it to return

	return args.Int(0)
}
func Test_retDummy(t *testing.T) {
	dummy := retDummyStr()
	assert.Equal(t, "dummy", dummy)
}

func Test_AddValue(t *testing.T){
   routerService := routerServiceMock{}
   routerService.On("AddValue",3.4,4.1).Return(5.0)
   result := routerService.AddValue(3.4,4.1)
   assert.Equal(t, 5.0, result)
}