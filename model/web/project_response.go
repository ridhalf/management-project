package web

import (
	"management-project/model/domain"
)

type ProjectResponse struct {
	ID          int                  `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	StartDate   string               `json:"start_date"`
	EndDate     string               `json:"end_date"`
	Authored    UserFindByIdResponse `json:"authored"`
}

func ToProjectResponse(project domain.Project) ProjectResponse {
	//startDate, _ := time.Parse(time.RFC3339, project.StartDate)
	//endDate, _ := time.Parse(time.RFC3339, project.EndDate)
	return ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		StartDate:   project.StartDate.Format("2006-01-02"),
		EndDate:     project.EndDate.Format("2006-01-02"),
		Authored: UserFindByIdResponse{
			Id:       project.User.ID,
			Username: project.User.Username,
			Email:    project.User.Email,
			Role:     project.User.Role,
		},
	}
}
func ToPorjectResponses(projects []domain.Project) []ProjectResponse {
	var projectResponses []ProjectResponse
	for _, project := range projects {
		projectResponse := ToProjectResponse(project)
		projectResponses = append(projectResponses, projectResponse)
	}
	return projectResponses
}
