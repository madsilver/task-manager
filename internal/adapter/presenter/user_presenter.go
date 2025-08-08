package presenter

import "github.com/madsilver/task-manager/internal/domain"

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ToUserResponse converts a domain.User to a UserResponse.
// user: the domain.User to convert
// returns *UserResponse
func ToUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
