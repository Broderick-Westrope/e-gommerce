package web

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Broderick-Westrope/e-gommerce/internal/config"
	"github.com/Broderick-Westrope/e-gommerce/internal/storage"
	"github.com/go-chi/chi/v5"
)

func handleGetEntities[T any](logger config.Logger, getFunc func() (*[]T, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entities, err := getFunc()
		if err != nil {
			messages := []string{"Failed to get entities", "get_entities_error", err.Error()}
			respondWithError(w, logger, http.StatusInternalServerError, messages...)
			return
		}

		respondWithJSON(w, logger, http.StatusOK, entities)
	}
}

func handleGetEntityByID[T any](logger config.Logger, getFunc func(int) (*T, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			messages := []string{"Invalid parameter 'id'", "atoi_error", err.Error()}
			respondWithError(w, logger, http.StatusBadRequest, messages...)
			return
		}

		entity, err := getFunc(id)
		if err != nil {
			var notFoundErr *storage.NotFoundError
			if errors.As(err, &notFoundErr) {
				messages := []string{"Entity not found", "get_entity_error", notFoundErr.Error()}
				respondWithError(w, logger, http.StatusNotFound, messages...)
				return
			}
			messages := []string{"Failed to get entity", "get_entity_error", err.Error()}
			respondWithError(w, logger, http.StatusInternalServerError, messages...)
			return
		}

		respondWithJSON(w, logger, http.StatusOK, entity)
	}
}

func handleCreateEntity[T any](logger config.Logger, createFunc func(*T) (int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createEntityReq T
		err := parseJSONBody(r, &createEntityReq)
		if err != nil {
			messages := []string{"Failed to parse JSON payload", "parse_json_body_error", err.Error()}
			respondWithError(w, logger, http.StatusInternalServerError, messages...)
			return
		}

		var id int
		id, err = createFunc(&createEntityReq)
		if err != nil {
			messages := []string{"Failed to create entity", "create_entity_error", err.Error()}
			respondWithError(w, logger, http.StatusInternalServerError, messages...)
			return
		}

		respondWithID(w, logger, http.StatusCreated, id)
	}
}

func handleUpdateEntityByID[T any](
	logger config.Logger,
	updateFunc func(*T) error,
	reqToEntityFunc func(*http.Request, int) (*T, error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			messages := []string{"Invalid parameter 'id'", "atoi_error", err.Error()}
			respondWithError(w, logger, http.StatusBadRequest, messages...)
			return
		}

		var entity *T
		entity, err = reqToEntityFunc(r, id)
		if err != nil {
			messages := []string{"Failed to parse JSON payload", "parse_json_body_error", err.Error()}
			respondWithError(w, logger, http.StatusInternalServerError, messages...)
			return
		}

		err = updateFunc(entity)
		if err != nil {
			var notFoundErr *storage.NotFoundError
			if errors.As(err, &notFoundErr) {
				messages := []string{"Entity not found", "update_entity_error", notFoundErr.Error()}
				respondWithError(w, logger, http.StatusNotFound, messages...)
				return
			}
			messages := []string{"Failed to update entity", "update_entity_error", err.Error()}
			respondWithError(w, logger, http.StatusInternalServerError, messages...)
			return
		}

		respondWithJSON(w, logger, http.StatusNoContent, nil)
	}
}

func handleDeleteEntityByID(logger config.Logger, deleteFunc func(int) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			messages := []string{"Invalid parameter 'id'", "atoi_error", err.Error()}
			respondWithError(w, logger, http.StatusBadRequest, messages...)
			return
		}

		err = deleteFunc(id)
		if err != nil {
			var notFoundErr *storage.NotFoundError
			if errors.As(err, &notFoundErr) {
				messages := []string{"Entity not found", "delete_entity_error", notFoundErr.Error()}
				respondWithError(w, logger, http.StatusNotFound, messages...)
				return
			}
			messages := []string{"Failed to delete entity", "delete_entity_error", err.Error()}
			respondWithError(w, logger, http.StatusInternalServerError, messages...)
			return
		}

		respondWithJSON(w, logger, http.StatusNoContent, nil)
	}
}
