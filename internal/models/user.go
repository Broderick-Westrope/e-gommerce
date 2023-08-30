package models

// User is a struct that defines the fields of a user.
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserRequest is a struct that defines the fields required to create a user.
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ToUser converts a CreateUserRequest to a User with the given id.
func (c *CreateUserRequest) ToUser(id int) *User {
	return &User{
		ID:       id,
		Email:    c.Email,
		Password: c.Password,
	}
}
