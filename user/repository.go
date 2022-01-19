package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
}

// private for package user
type repository struct {
	db *gorm.DB
}

// kalau ga salah tangkep
// dia akan return nilai pointer berupa alama pointer repository itu sendiri
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
