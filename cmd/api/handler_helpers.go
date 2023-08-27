package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Broderick-Westrope/e-gommerce/internal/config"
)

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
		errMsg := "Failed to encode JSON payload"
		logger.Error(errMsg)
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(createErrorResponse(errMsg))
		if err != nil {
			logger.Error("Failed to encode JSON payload for error response: " + err.Error())
		}
		return
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(buf.Bytes())
	if err != nil {
		logger.Error("Failed to write JSON payload: " + err.Error())
	}
}

// respondWithError is a helper function to respond with an error.
// It also sets the Content-Type header to application/json.
func respondWithError(w http.ResponseWriter, logger config.Logger, statusCode int, message string) {
	logger.Error(message)
	mapResponse := createErrorResponse(message)
	respondWithJSON(w, logger, statusCode, mapResponse)
}

// parseJSONBody unmarshals the JSON payload and stores the result in the provided destination.
func parseJSONBody(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

// createErrorResponse is a helper function to create an error response map.
func createErrorResponse(message string) map[string]string {
	return map[string]string{"error": message}
}
