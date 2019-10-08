package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ResponseOK(dto interface{}, writer *http.ResponseWriter) {
	responseWriter := *writer
	responseWriter.WriteHeader(http.StatusOK)

	responseWriter.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(*writer).Encode(dto)
}
func ResponseError(dto interface{}, status int, writer *http.ResponseWriter) {
	responseWriter := *writer
	responseWriter.WriteHeader(status)
	responseWriter.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(*writer).Encode(dto)
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}
func ReadUrlEncode(vars string, request *http.Request) string {
	var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(request.Body)
	return string(bodyBytes)
}
