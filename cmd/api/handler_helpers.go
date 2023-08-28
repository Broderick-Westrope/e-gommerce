package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/oklog/ulid/v2"
)

type errorResponse struct {
	ID    string `json:"error_id"`
	Error string `json:"error"`
}

type idResponse struct {
	ID int `json:"id"`
}

// respondWithJSON is a helper function to respond with the JSON payload.
// It also sets the Content-Type header to application/json.
// If the JSON payload cannot be encoded, it will write an Internal Server Error to the response.
func respondWithJSON(w http.ResponseWriter, logger config.Logger, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if payload == nil {
		w.WriteHeader(statusCode)
		return
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		messages := []string{"Failed to encode JSON payload", "encode_error", err.Error()}
		response := createErrorResponse(logger, messages...)
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			logger.Error("Failed to encode JSON payload for error response", "encode_error", err.Error())
		}
		return
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(buf.Bytes())
	if err != nil {
		// Since the response has already been written, we can only log the error.
		logger.Error("Failed to write JSON payload", "write_error", err.Error())
	}
}

// respondWithID is a helper function to respond with an idResponse.
// It also sets the Content-Type header to application/json.
func respondWithID(w http.ResponseWriter, logger config.Logger, statusCode int, id int) {
	response := idResponse{id}
	respondWithJSON(w, logger, statusCode, response)
}

// respondWithError is a helper function to respond with an errorResponse.
// It also sets the Content-Type header to application/json.
// It will log all messages to the logger. Check the logger implementation for more details.
func respondWithError(w http.ResponseWriter, logger config.Logger, statusCode int, messages ...string) {
	errResponse := createErrorResponse(logger, messages...)
	respondWithJSON(w, logger, statusCode, errResponse)
}

// parseJSONBody unmarshals the JSON payload and stores the result in the provided destination.
func parseJSONBody(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

// createErrorResponse is a helper function to create an errorResponse using all the first message.
// It will log all messages to the logger. Check the logger implementation for more details.
func createErrorResponse(logger config.Logger, messages ...string) errorResponse {
	errID := ulid.Make()
	response := errorResponse{ID: errID.String(), Error: messages[0]}

	args := []interface{}{"error_id", errID.String()}
	for _, s := range messages[1:] {
		if s != "" {
			args = append(args, s)
		}
	}

	logger.Error(messages[0], args...)
	return response
}
