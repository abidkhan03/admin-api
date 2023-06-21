package api

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/spongeling/admin-api/internal/errors"
	"github.com/spongeling/admin-api/internal/request"

	"github.com/homelight/json"
)

type API interface {
	Routes(r chi.Router)
}

// parseRequest gets json from request and fills the target model
func parseRequest(r *http.Request, target interface{}) error {
	tgt, ok := target.(request.Request)
	if !ok {
		return errors.New(errors.InvalidArgument, "parseRequest: target is expected to be of type *request.Request")
	}

	err := parseBody(r.Body, target)
	if err != nil {
		return err
	}

	return tgt.Validate()
}

func extractIdFromUrl(r *http.Request, name string) (uint64, error) {
	return strconv.ParseUint(chi.URLParam(r, name), 10, 0)
}

// parseBody gets json from request body and fills the target model
func parseBody(readcloser io.ReadCloser, target interface{}) error {
	body, err := io.ReadAll(readcloser)
	if err != nil {
		return err
	}
	log.Printf("Request body:\t%s\n", string(body))
	err = json.Unmarshal(body, target)
	return err
}

// structToCompactJson returns compact json string of the given struct
func structToCompactJson(source interface{}) (string, error) {
	jsonBytes, err := json.Marshal(source)
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, jsonBytes); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

// respondBytes makes the response with payload as an array of bytes
func respondBytes(w http.ResponseWriter, status int, payload []byte) {
	log.Printf("Response body:\t%s\n", string(payload))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(payload)
}

// respondAny makes the response with payload as an array of bytes
func respondAny(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.MarshalSafeCollections(payload)
	if err != nil {
		log.Println(err)
		respondError(w, http.StatusInternalServerError)
		return
	}
	respondBytes(w, status, response)
}

// respondOk makes the response with payload as json format
func respondOk(w http.ResponseWriter, payload interface{}) {
	respondAny(w, http.StatusOK, payload)
}

// respondErrorMessage makes the error response with given message in json format
func respondErrorMessage(w http.ResponseWriter, code int, message string) {
	respondBytes(w, code, []byte(fmt.Sprintf(`{"code":%v,"message":"%v"}`, code, message)))
}

// respondError makes the error response with default message in json format
func respondError(w http.ResponseWriter, code int) {
	var message string
	switch code {
	case http.StatusBadRequest:
		message = "The request had invalid inputs or otherwise cannot be served."
	case http.StatusUnauthorized:
		message = "Authorization information is missing or invalid."
	case http.StatusForbidden:
		message = "Access denied to the resource."
	case http.StatusNotFound:
		message = "Unable to find requested record."
	case http.StatusNotAcceptable:
		message = "Not acceptable for the database."
	case http.StatusRequestTimeout:
		message = "Request took too long to process."
	case http.StatusConflict:
		message = "A conflict has occurred."
	case http.StatusRequestedRangeNotSatisfiable:
		message = "No resource available, unable to fulfill the request."
	case http.StatusTooManyRequests:
		message = "Request rate too high, requests from this this user are throttled."
	case http.StatusInternalServerError:
		message = "An error was encountered."
	case http.StatusServiceUnavailable:
		message = "The service is unavailable, please try again later."
	case http.StatusGatewayTimeout:
		message = "The service timed out waiting for an upstream response. Try again later."
	}

	respondBytes(w, code, []byte(fmt.Sprintf(`{"code":%v,"message":"%v"}`, code, message)))
}
