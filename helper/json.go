package helper

import (
	"encoding/json"
	"net/http"
)

func RequestDecode(request *http.Request, body interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(body)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
