package server

import(
	"testing"

	"github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/assert"
)

type routerServiceMock struct {
	mock.Mock
}

func (m *routerServiceMock) AddValue(a,b int) int{
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)
  // this records that the method was called and passes in the value
  // it was called with
  args := m.Called(value)
  // it then returns whatever we tell it to return
  // in this case true to simulate an SMS Service Notification
  // sent out 
	return args.Bool(0)
}
func Test_retDummy(t *testing.T){
	dummy := retDummyStr();
	assert.Equal(t,"dummy",dummy)
}