package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type routerServiceMock struct {
	mock.Mock
}

func (m *routerServiceMock) AddValue(a, b int) int {
	fmt.Println("AddValue function")
	fmt.Printf("Value passed in: %d and %d", a,b)
	// this records that the method was called and passes in the value
	// it was called with
	args := m.Called(a,b)
	// it then returns whatever we tell it to return
	// in this case true to simulate an SMS Service Notification
	// sent out
	return args.Int(5)
}
func Test_retDummy(t *testing.T) {
	dummy := retDummyStr()
	assert.Equal(t, "dummy", dummy)
}

func Test_AddValue(t *testing.T){
   routerService := routerServiceMock{}
   routerService.On
   result := routerService.AddValue(3,4)
   assert.Equal(t, 7, result)
}