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

func TestAddValue(t *testing.T) {
	s := testServer()
	//Init
	fromdate := sTime1 //domain.CustomTime{Time: sTime1}
	todate := eTime1   //domain.CustomTime{Time: eTime1}
	//jwtToken = GetJwtToken()
	//Create Event Validation
	calendarEventModel := domain.CalendarEventModel{CalendarEventBaseModel: domain.CalendarEventBaseModel{Classid: "", UID: ""}, Title: "",
		Metada: nil, ShowStudent: "true1212", StartDate: fromdate, EndDate: todate}
	reqJSON, _ := json.Marshal(&calendarEventModel)
	req, _ := http.NewRequest("POST", "/v1/calendars/event", bytes.NewBuffer(reqJSON))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))
	req.Header.Add("Authorization", "bearer "+jwtToken)

	resp, body := mockHTTPServer(t, s.RegisterHandlers(), req)
	respModel := []domain.APIResponse{}
	json.Unmarshal(body, &respModel)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	//Title max length and invalid dates
	eTimeW, _ := time.Parse("2006-01-02", "")
	fromdateW := eTimeW //domain.CustomTime{Time: eTimeW}
	todateW := eTimeW   //domain.CustomTime{Time: eTimeW}
	calendarEventModel = domain.CalendarEventModel{CalendarEventBaseModel: domain.CalendarEventBaseModel{Classid: "", UID: ""},
		Title:  "  This title is more than sixty character long. This title is more than sixty character long. test  ",
		Metada: nil, ShowStudent: "true1212", StartDate: fromdateW, EndDate: todateW}
	reqJSON, _ = json.Marshal(&calendarEventModel)
	req, _ = http.NewRequest("POST", "/v1/calendars/event", bytes.NewBuffer(reqJSON))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))
	req.Header.Add("Authorization", "bearer "+jwtToken)

	resp, body = mockHTTPServer(t, s.RegisterHandlers(), req)
	respModel = []domain.APIResponse{}
	json.Unmarshal(body, &respModel)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	//Start date greater than end date
	sTimeW, _ := time.Parse("2006-01-02", "2017-02-05")
	eTimeW, _ = time.Parse("2006-01-02", "2017-02-03")
	fromdateW = sTimeW //domain.CustomTime{Time: sTimeW}
	todateW = eTimeW   //domain.CustomTime{Time: eTimeW}
	calendarEventModel = domain.CalendarEventModel{CalendarEventBaseModel: domain.CalendarEventBaseModel{Classid: "", UID: ""},
		Title:  "  This title is more than sixty character long. This title is more than sixty character long. test  ",
		Metada: nil, ShowStudent: "true1212", StartDate: fromdateW, EndDate: todateW}
	reqJSON, _ = json.Marshal(&calendarEventModel)
	req, _ = http.NewRequest("POST", "/v1/calendars/event", bytes.NewBuffer(reqJSON))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))
	req.Header.Add("Authorization", "bearer "+jwtToken)

	resp, body = mockHTTPServer(t, s.RegisterHandlers(), req)
	respModel = []domain.APIResponse{}
	json.Unmarshal(body, &respModel)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	metadata := map[string]string{"key": "val"}
	//Create Event
	calendarEventModel = domain.CalendarEventModel{CalendarEventBaseModel: domain.CalendarEventBaseModel{Classid: classID, UID: "dk"}, Title: "test title",
		Metada: metadata, ShowStudent: "true", StartDate: fromdate, EndDate: todate}
	reqJSON, _ = json.Marshal(&calendarEventModel)
	req, _ = http.NewRequest("POST", "/v1/calendars/event", bytes.NewBuffer(reqJSON))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))
	req.Header.Add("Authorization", "bearer "+jwtToken)

	resp, body = mockHTTPServer(t, s.RegisterHandlers(), req)
	respModel = []domain.APIResponse{}
	json.Unmarshal(body, &respModel)
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	caleventResp := domain.APIResponse{}.Body
	json.Unmarshal(body, &caleventResp)

	eventid := caleventResp.Data.UUID
	fmt.Println(eventid) //Read Event
	//Validations
	//End date before start date
	req, _ = http.NewRequest("GET", "/v1/calendars/event?classid="+classID+"&startdate="+eDate+"&enddate="+sDate, nil)
	req.Header.Add("Authorization", "bearer "+jwtToken)
	resp, body = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	//classID missing
	req, _ = http.NewRequest("GET", "/v1/calendars/event?classid=&startdate="+eDate+"&enddate="+sDate, nil)
	req.Header.Add("Authorization", "bearer "+jwtToken)
	resp, body = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	//Invlaid start date
	req, _ = http.NewRequest("GET", "/v1/calendars/event?classid="+classID+"&startdate=2017-21-21&enddate="+sDate, nil)
	req.Header.Add("Authorization", "bearer "+jwtToken)
	resp, body = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	//Invlaid End date
	req, _ = http.NewRequest("GET", "/v1/calendars/event?classid="+classID+"&startdate=2017-21-21&enddate="+sDate, nil)
	req.Header.Add("Authorization", "bearer "+jwtToken)
	resp, body = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	//correct input params
	req, _ = http.NewRequest("GET", "/v1/calendars/event?classid="+classID+"&startdate="+sDate+"&enddate="+eDate, nil)
	req.Header.Add("Authorization", "bearer "+jwtToken)
	resp, body = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// caleventResp := domain.APIResponse{}.Body
	// json.Unmarshal(body, &caleventResp)

	// eventid := caleventResp.Data.UUID
	// fmt.Println(eventid)
	//eventid := "123"

	//Update Event
	//Invalid params
	calendarEventModel = domain.CalendarEventModel{CalendarEventBaseModel: domain.CalendarEventBaseModel{Classid: classID, UID: ""}, Title: "",
		Metada: nil, ShowStudent: "true", StartDate: fromdate, EndDate: todate}
	reqJSON, _ = json.Marshal(&calendarEventModel)
	req, _ = http.NewRequest("PUT", "/v1/calendars/event/"+eventid, bytes.NewBuffer(reqJSON))
	req.Header.Add("Authorization", "bearer "+jwtToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))

	resp, _ = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	////Update Non Existing record
	calendarEventModel = domain.CalendarEventModel{CalendarEventBaseModel: domain.CalendarEventBaseModel{Classid: classID, UID: "dk1221"}, Title: "test title123",
		Metada: metadata, ShowStudent: "true", StartDate: fromdate, EndDate: todate}
	reqJSON, _ = json.Marshal(&calendarEventModel)
	req, _ = http.NewRequest("PUT", "/v1/calendars/event/df27139e-d38f-11e7-9296-cec278b6b50b", bytes.NewBuffer(reqJSON))
	req.Header.Add("Authorization", "bearer "+jwtToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))

	resp, _ = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)

	//correct params
	calendarEventModel = domain.CalendarEventModel{CalendarEventBaseModel: domain.CalendarEventBaseModel{Classid: classID, UID: "dk1221"}, Title: "test title123",
		Metada: metadata, ShowStudent: "true", StartDate: fromdate, EndDate: todate}
	reqJSON, _ = json.Marshal(&calendarEventModel)
	req, _ = http.NewRequest("PUT", "/v1/calendars/event/"+eventid, bytes.NewBuffer(reqJSON))
	req.Header.Add("Authorization", "bearer "+jwtToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))

	resp, _ = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	//Delete Event
	//Invalid params
	calendarDelModel := domain.CalendarEventDeleteModel{UID: ""}
	reqJSON, _ = json.Marshal(&calendarDelModel)

	req, _ = http.NewRequest("DELETE", "/v1/calendars/event/123213213213", bytes.NewBuffer(reqJSON))
	req.Header.Add("Authorization", "bearer "+jwtToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))

	resp, _ = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	//Delete Non Existing record
	calendarDelModel = domain.CalendarEventDeleteModel{UID: "dk"}
	reqJSON, _ = json.Marshal(&calendarDelModel)

	req, _ = http.NewRequest("DELETE", "/v1/calendars/event/df27139e-d38f-11e7-9296-cec278b6b50a", bytes.NewBuffer(reqJSON))
	req.Header.Add("Authorization", "bearer "+jwtToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))

	resp, _ = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)

	//correct params
	calendarDelModel = domain.CalendarEventDeleteModel{UID: "dk"}
	reqJSON, _ = json.Marshal(&calendarDelModel)

	req, _ = http.NewRequest("DELETE", "/v1/calendars/event/"+eventid, bytes.NewBuffer(reqJSON))
	req.Header.Add("Authorization", "bearer "+jwtToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))

	resp, _ = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	dao.DelCalEvent(eventid)

	//Validate jwt
	calendarDelModel = domain.CalendarEventDeleteModel{UID: "dk"}
	reqJSON, _ = json.Marshal(&calendarDelModel)

	req, _ = http.NewRequest("DELETE", "/v1/calendars/event/"+eventid, bytes.NewBuffer(reqJSON))
	req.Header.Add("Authorization", "bearer "+"some random token string")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(reqJSON)))

	resp, _ = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	//OPtion Call
	req, _ = http.NewRequest("OPTIONS", "/v1/calendars/event", nil)
	req.Header.Add("Authorization", "bearer "+jwtToken)
	resp, body = mockHTTPServer(t, s.RegisterHandlers(), req)
	require.Equal(t, http.StatusOK, resp.StatusCode)

}