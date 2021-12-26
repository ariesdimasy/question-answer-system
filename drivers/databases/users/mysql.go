package users

import (
	_userDomain "acp-final/business/users"
	"acp-final/helpers"
	"context"

	"gorm.io/gorm"
)

type UserRepositoryDatabase struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) _userDomain.UserRepository {
	return &UserRepositoryDatabase{
		db: database,
	}
}

func (repo *UserRepositoryDatabase) GetAllUser(ctx context.Context) ([]_userDomain.UserDomain, error) {

	var users []User
	result := repo.db.Find(&users)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_userDomain.UserDomain{}, helpers.ErrNotFound
		} else {
			return []_userDomain.UserDomain{}, helpers.ErrDbServer
		}
	}

	return ToListDomain(users), nil
}

func (repo *UserRepositoryDatabase) Get(ctx context.Context) ([]_userDomain.UserDomain, error) {

	var users []User
	result := repo.db.Find(&users)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return []_userDomain.UserDomain{}, helpers.ErrNotFound
		} else {
			return []_userDomain.UserDomain{}, helpers.ErrDbServer
		}
	}

	return ToListDomain(users), nil
}
