package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"management-project/model/domain"
	"management-project/model/web"
	"management-project/repository"
	"time"
)

type ProjectService interface {
	FindAll(ctx *gin.Context) ([]domain.Project, error)
	FindById(ctx *gin.Context, request web.ProjectFindByIdRequest) (domain.Project, error)
	Add(ctx *gin.Context, request web.ProjectCreateRequest) (domain.Project, error)
	Update(ctx *gin.Context, request web.ProjectUpdateRequest) (domain.Project, error)
	Delete(ctx *gin.Context, request web.ProjectFindByIdRequest) error
}
type ProjectServiceImpl struct {
	repositoryProject repository.ProjectRepository
}

func NewProjectService(repositoryProject repository.ProjectRepository) ProjectService {
	return &ProjectServiceImpl{
		repositoryProject: repositoryProject,
	}
}

func (service ProjectServiceImpl) FindAll(ctx *gin.Context) ([]domain.Project, error) {
	projects, err := service.repositoryProject.FindAll()
	if err != nil {
		return nil, errors.New("an unexpected error occurred while accessing the database")
	}
	return projects, nil
}

func (service ProjectServiceImpl) FindById(ctx *gin.Context, request web.ProjectFindByIdRequest) (domain.Project, error) {
	project, err := service.repositoryProject.FindByID(request.ID)
	if err != nil {
		return domain.Project{}, errors.New("an unexpected error occurred while accessing the database")
	}
	return project, nil
}

func (service ProjectServiceImpl) Add(ctx *gin.Context, request web.ProjectCreateRequest) (domain.Project, error) {
	user := ctx.MustGet("user").(domain.User)
	startDate, _ := time.Parse("2006-01-02", request.StartDate)
	endDate, _ := time.Parse("2006-01-02", request.EndDate)
	project := domain.Project{
		Name:        request.Name,
		Description: request.Description,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedBy:   user.ID,
	}
	create, err := service.repositoryProject.Create(project)
	if err != nil {
		return domain.Project{}, errors.New("an unexpected error occurred while accessing the database")
	}
	return create, nil
}

func (service ProjectServiceImpl) Update(ctx *gin.Context, request web.ProjectUpdateRequest) (domain.Project, error) {
	startDate, _ := time.Parse(time.RFC3339, request.StartDate)
	endDate, _ := time.Parse(time.RFC3339, request.EndDate)
	project := domain.Project{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		StartDate:   startDate,
		EndDate:     endDate,
	}
	update, err := service.repositoryProject.Update(project)
	if err != nil {
		return domain.Project{}, errors.New("an unexpected error occurred while accessing the database")
	}
	return update, nil
}

func (service ProjectServiceImpl) Delete(ctx *gin.Context, request web.ProjectFindByIdRequest) error {
	count, err := service.repositoryProject.Count(request.ID)
	if err != nil {
		return errors.New("an unexpected error occurred while accessing the database")
	}
	if count == 0 {
		return errors.New("data does not exist")
	}
	err = service.repositoryProject.Delete(request.ID)
	if err != nil {
		return errors.New("an unexpected error occurred while accessing the database")
	}
	return nil
}
