package query

import (
	"github.com/caclm10/simpletodo-api/internal/model"
	"gorm.io/gorm"
)

func FindUserByEmail(db *gorm.DB, email string) (model.User, bool, error) {
	user := model.User{}

	result := db.Where("email = ?", email).Limit(1).Find(&user)
	if result.Error != nil {
		return user, false, result.Error
	}

	if result.RowsAffected > 0 {
		return user, true, nil
	}

	return user, false, nil
}
