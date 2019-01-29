package server

import (
	"github.com/stretchr/testify/require"
	"encoding/json"
	"net/http"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type calcServiceMock struct {
	mock.Mock
}

func (m *calcServiceMock) AddValue(a, b float64) float64 {
	fmt.Println("AddValue function")
	fmt.Printf("Value passed in: %f and %f", a, b)
	// this records that the method was called and passes in the value
	// it was called with
	args := m.Called(a, b)
	// it then returns whatever we tell it to return

	val := args[0].(float64)

	//valFloat,_:= strconv.ParseFloat(val, 64)

	return val
}
func Test_retDummy(t *testing.T) {
	dummy := retDummyStr()
	assert.Equal(t, "dummy", dummy)
}

func Test_AddValue(t *testing.T) {
	caclService := CalcService{}
	result := caclService.AddValue(3.41, 4.11)
	assert.Equal(t, 7.5200000000000005, result)
}

func Test_AddValueMock(t *testing.T) {
	caclService := calcServiceMock{}
	caclService.On("AddValue", 3.4, 4.1).Return(7.52)
	result := caclService.AddValue(3.4, 4.1)
	assert.Equal(t, 7.52, result)
}

func Test_AddValueMock1(t *testing.T) {
	caclService := calcServiceMock{}
	caclService.On("AddValue", 3.4, 4.1).Return(7.52)
	result := caclService.AddValue(3.4, 4.2)
	assert.Equal(t, 7.52, result)
}

func TestAddValue(t *testing.T) {
	s := testServer()
	//Init
	req, _ := http.NewRequest("GET", "/calc/add/?op1=10&op2=10", nil)

	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, body := mockHTTPServer(t, s.RegisterHandlers(), req)

	respModel := ResultSet{}
	json.Unmarshal(body, &respModel)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	assert.Equal(t,float64(20),respModel.Value)
	assert.NotNil(t,body)
}