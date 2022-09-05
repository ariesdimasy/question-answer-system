package questions

import (
	_questionDomain "qa-system/business/questions"

	"gorm.io/gorm"
)

type UserRepositoryDatabase struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) _questionDomain.QuestionRepository {
	return &UserRepositoryDatabase{
		db: database,
	}
}
