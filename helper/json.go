package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequest(r *http.Request, result interface{}) {

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(result)
	ErrorPanic(err)

}

func WriteRequest(w http.ResponseWriter, response interface{}) {

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	ErrorPanic(err)

}
