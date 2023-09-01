package web

import (
	"net/http"

	"github.com/Broderick-Westrope/e-gommerce/internal/models"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(srv Server) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handleGetUsers(srv))
	router.Get("/{id}", handleGetUserByID(srv))
	router.Post("/", handleCreateUser(srv))
	router.Put("/{id}", handleUpdateUserByID(srv))
	router.Delete("/{id}", handleDeleteUserByID(srv))

	return router
}

func requestToUser(r *http.Request, id int) (*models.User, error) {
	var createUserReq models.CreateUserRequest
	err := parseJSONBody(r, &createUserReq)
	if err != nil {
		return nil, err
	}
	return createUserReq.ToUser(id), nil
}

//	@Summary		Get all users
//	@Description	Retrieves all users.
//	@ID				get-users
//	@Tags			users
//	@Produce		json
//	@Success		200	{array}		models.User		"Users"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/users [get]
func handleGetUsers(srv Server) http.HandlerFunc {
	return handleGetEntities[models.User](srv.Logger(), srv.Storage().GetUsers)
}

//	@Summary		Get a user
//	@Description	Retrieves a user by ID.
//	@ID				get-user
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int				true	"User ID"
//	@Success		200	{object}	models.User		"User"
//	@Failure		400	{object}	errorResponse	"Invalid parameter 'id'"
//	@Failure		404	{object}	errorResponse	"User not found"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/users/{id} [get]
func handleGetUserByID(srv Server) http.HandlerFunc {
	return handleGetEntityByID[models.User](srv.Logger(), srv.Storage().GetUser)
}

//	@Summary		Create a user
//	@Description	Creates a user.
//	@ID				create-user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.CreateUserRequest	true	"User"
//	@Success		201		{object}	idResponse					"User ID"
//	@Failure		500		{object}	errorResponse				"Internal Server Error"
//	@Router			/users [post]
func handleCreateUser(srv Server) http.HandlerFunc {
	return handleCreateEntity[models.CreateUserRequest](srv.Logger(), srv.Storage().CreateUser)
}

//	@Summary		Update a user
//	@Description	Updates a user.
//	@ID				update-user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int							true	"User ID"
//	@Param			user	body	models.CreateUserRequest	true	"User"
//	@Success		204
//	@Failure		400	{object}	errorResponse	"Invalid parameter 'id'"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/users/{id} [put]
func handleUpdateUserByID(srv Server) http.HandlerFunc {
	return handleUpdateEntityByID[models.User](srv.Logger(), srv.Storage().UpdateUser, requestToUser)
}

//	@Summary		Delete a user
//	@Description	Deletes a user.
//	@ID				delete-user
//	@Tags			users
//	@Param			id	path	int	true	"User ID"
//	@Success		204
//	@Failure		400	{object}	errorResponse	"Invalid parameter 'id'"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/users/{id} [delete]
func handleDeleteUserByID(srv Server) http.HandlerFunc {
	return handleDeleteEntityByID(srv.Logger(), srv.Storage().DeleteUser)
}
