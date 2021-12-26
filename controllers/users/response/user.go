package response

import (
	_usersDomain "acp-final/business/users"
	"time"
)

type UserResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
}

func FromDomain(user _usersDomain.UserDomain) UserResponse {
	return UserResponse{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
	}
}

func ToListFromDomain(users []_usersDomain.UserDomain) []UserResponse {
	var result = []UserResponse{}
	for _, user := range users {
		result = append(result, FromDomain(user))
	}
	return result
}
