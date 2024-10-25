package repository

import (
	"gorm.io/gorm"
	"management-project/model/domain"
)

type ProjectRepository interface {
	FindAll() ([]domain.Project, error)
	FindByID(id int) (domain.Project, error)
	Create(project domain.Project) (domain.Project, error)
	Update(project domain.Project) (domain.Project, error)
	Count(id int) (int, error)
	Delete(id int) error
}

type ProjectRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &ProjectRepositoryImpl{
		db: db,
	}
}

func (repository ProjectRepositoryImpl) FindAll() ([]domain.Project, error) {
	var projects []domain.Project
	err := repository.db.Model(&domain.Project{}).Preload("User").Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (repository ProjectRepositoryImpl) FindByID(id int) (domain.Project, error) {
	project := domain.Project{}
	err := repository.db.First(&project, id).Error
	if err != nil {
		return domain.Project{}, err
	}
	return project, nil
}

func (repository ProjectRepositoryImpl) Create(project domain.Project) (domain.Project, error) {
	err := repository.db.Save(&project).Error
	if err != nil {
		return domain.Project{}, err
	}
	return project, nil
}

func (repository ProjectRepositoryImpl) Update(project domain.Project) (domain.Project, error) {
	err := repository.db.Save(&project).Error
	if err != nil {
		return domain.Project{}, err
	}
	return project, nil
}

func (repository ProjectRepositoryImpl) Count(id int) (int, error) {
	var count int64
	err := repository.db.Model(&domain.Project{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (repository ProjectRepositoryImpl) Delete(id int) error {
	err := repository.db.Delete(&domain.Project{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
