package response

import (
	"fmt"
	"log"
	"net/http"
)

func respondBytes(w http.ResponseWriter, status int, payload []byte) {
	log.Printf("Response body:\t%s\n", string(payload))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(payload)
}

// RespondErrorMessage makes the error response with given message in json format
func RespondErrorMessage(w http.ResponseWriter, code int, message string) {
	respondBytes(w, code, []byte(fmt.Sprintf(`{"code":%v,"message":"%v"}`, code, message)))
}
