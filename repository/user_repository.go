package repository

import (
	"gorm.io/gorm"
	"management-project/model/domain"
)

type UserRepository interface {
	FindByUsername(username string) (domain.User, error)
	FindById(id int) (domain.User, error)
	Save(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Count(ID int) (int64, error)
}
type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (repository UserRepositoryImpl) FindByUsername(username string) (domain.User, error) {
	user := domain.User{}
	err := repository.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) FindById(id int) (domain.User, error) {
	user := domain.User{}
	err := repository.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	err := repository.db.Save(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (repository UserRepositoryImpl) Update(user domain.User) (domain.User, error) {
	err := repository.db.Save(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (repository UserRepositoryImpl) Count(ID int) (int64, error) {
	var count int64
	err := repository.db.Model(&domain.User{}).Where("id = ?", ID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
