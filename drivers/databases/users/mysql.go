package users

import (
	_userDomain "qa-system/business/users"

	"context"
	"qa-system/helpers"

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

func (repo *UserRepositoryDatabase) GetUserById(ctx context.Context, id_user string) (_userDomain.UserDomain, error) {
	var user User
	result := repo.db.Debug().Where("id = ?", id_user).Find(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _userDomain.UserDomain{}, helpers.ErrNotFound
		} else {
			return _userDomain.UserDomain{}, helpers.ErrDbServer
		}
	}

	return user.ToDomain(), nil
}

func (repo *UserRepositoryDatabase) Login(ctx context.Context, userLogin _userDomain.UserLoginDomain) (_userDomain.UserDomain, error) {
	var user User
	result := repo.db.Debug().Where("email = ? AND password = ? ", userLogin.Email, userLogin.Password).Find(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _userDomain.UserDomain{}, helpers.ErrNotFound
		} else {
			return _userDomain.UserDomain{}, helpers.ErrDbServer
		}
	}

	return user.ToDomain(), nil
}

func (repo *UserRepositoryDatabase) Register(ctx context.Context, userRegister _userDomain.UserRegisterDomain) (_userDomain.UserDomain, error) {
	var user User
	result := repo.db.Create(userRegister)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _userDomain.UserDomain{}, helpers.ErrNotFound
		} else {
			return _userDomain.UserDomain{}, helpers.ErrDbServer
		}
	}

	return user.ToDomain(), nil
}

func (repo *UserRepositoryDatabase) DeleteUser(ctx context.Context, id_user string) (_userDomain.UserDomain, error) {
	var user User
	result := repo.db.Delete(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return _userDomain.UserDomain{}, helpers.ErrNotFound
		} else {
			return _userDomain.UserDomain{}, helpers.ErrDbServer
		}
	}

	return user.ToDomain(), nil
}
