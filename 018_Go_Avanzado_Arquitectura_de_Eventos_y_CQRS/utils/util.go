package utils

import (
	"encoding/json"
	"net/http"
)

type Util interface {
	//MuxErrorResponse(w http.ResponseWriter, err error)
	WriteJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) (err error)
}

type util struct {
	env Env
}

func NewUtil(
	env Env) Util {
	return &util{
		env: env,
	}
}

/*
func (u *util) MuxErrorResponse(w http.ResponseWriter, err error) {
	fmt.Println(err.Error())
	u.logger.WriteLogEntry(logger.Error, err.Error())
	errResponse := errs.HttpResponseBody(&err, false)
	writeErr := u.WriteJSON(w, errResponse.Error.StatusCode, errResponse, nil)

	if writeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

*/

func (u *util) WriteJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) (err error) {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return
}
