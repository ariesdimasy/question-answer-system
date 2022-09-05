package users

import (
	_userDomain "qa-system/business/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
	Email     string `gorm:"not null;unique"`
	Password  string
}

func (user *User) ToDomain() _userDomain.UserDomain {
	return _userDomain.UserDomain{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func ToListDomain(users []User) []_userDomain.UserDomain {
	var result = []_userDomain.UserDomain{}
	for _, user := range users {
		result = append(result, user.ToDomain())
	}
	return result
}
