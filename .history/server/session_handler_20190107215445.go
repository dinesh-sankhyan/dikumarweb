package server

import (
	"net/http"
)

type LoginParams struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {

	params := &LoginParams{}
	if err := decodeJSON(r, params); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	errSet := services.GetValidator().Struct(params)
	if errSet != nil {
		// errs := convertErrToAPICodes(errSet)
		// resp.Body.Err = errs
		WriteJSON(w, http.StatusBadRequest, errSet)
		return
	}
	resp, err := services.Login(params)
	fmt.Printf("%#v", resp.SessionXID)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err)
	} else {
		sessionCookie, err := services.GetSession(r, "session", resp.SessionXID)
		if err != nil {
			logger.Errorf("login error while setting cookie for userid %s, err %+v", params.Username, err)
		}

		if sessionCookie==nil || len(sessionCookie) == 0 {
			sessobj := make(map[interface{}]interface{})
			sessobj["accesstoken"] = resp.AccessToken
			sessobj["refreshtoken"] = resp.RefreshToken
			sessobj["sessionXID"] = resp.SessionXID
			sessobj["test"] = 1234
			services.SaveSession(r, w, "session", resp.SessionXID, sessobj)
		}
		response.WriteJSON(w, http.StatusOK, resp)
	}

}

func decodeJSON(r *http.Request, model interface{}) *errors.ErrorResponse {
	if cType := r.Header.Get("Content-Type"); cType != "application/json" {
		return errors.BadRequest("Content-Type must be application/json", nil)
	}

	if r.Body == nil {
		return errors.BadRequest("failed to parse JSON request: empty body", nil)
	}

	// Parse params
	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		return errors.BadRequest("failed to parse JSON request: "+err.Error(), nil)
	}

	return nil
}