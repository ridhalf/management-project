package controller

import (
	"github.com/gin-gonic/gin"
	"management-project/model/web"
	"management-project/service"
)

type ProjectController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Add(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
type ProjectControllerImpl struct {
	serviceProject service.ProjectService
}

func NewProjectController(serviceProject service.ProjectService) ProjectController {
	return &ProjectControllerImpl{
		serviceProject: serviceProject,
	}
}

func (controller ProjectControllerImpl) FindAll(ctx *gin.Context) {
	if !AllowReadProject(ctx) {
		return
	}
	projects, err := controller.serviceProject.FindAll(ctx)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	responses := web.ToPorjectResponses(projects)
	HandleRequestSuccess(ctx, "successfully fetched the data.", responses)
	return
}

func (controller ProjectControllerImpl) FindById(ctx *gin.Context) {
	if !AllowReadProject(ctx) {
		return
	}
	var request web.ProjectFindByIdRequest
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		HandleBindError(ctx)
		return
	}
	project, err := controller.serviceProject.FindById(ctx, request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	responses := web.ToProjectResponse(project)
	HandleRequestSuccess(ctx, "successfully fetched the data.", responses)
	return
}

func (controller ProjectControllerImpl) Add(ctx *gin.Context) {
	if !AllowReadProject(ctx) {
		return
	}
	request := web.ProjectCreateRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleBindError(ctx)
		return
	}
	project, err := controller.serviceProject.Add(ctx, request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	responses := web.ToProjectUpsertResponse(project)
	HandleRequestSuccess(ctx, "successfully added the data.", responses)
	return
}

func (controller ProjectControllerImpl) Update(ctx *gin.Context) {
	if !AllowReadProject(ctx) {
		return
	}
	request := web.ProjectUpdateRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		HandleBindError(ctx)
		return
	}
	project, err := controller.serviceProject.Update(ctx, request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	responses := web.ToProjectUpsertResponse(project)
	HandleRequestSuccess(ctx, "successfully updated the data.", responses)
	return
}

func (controller ProjectControllerImpl) Delete(ctx *gin.Context) {
	if !AllowReadProject(ctx) {
		return
	}
	request := web.ProjectFindByIdRequest{}
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		HandleBindError(ctx)
		return
	}
	err = controller.serviceProject.Delete(ctx, request)
	if err != nil {
		HandleServiceError(ctx, err)
		return
	}
	HandleRequestSuccess(ctx, "successfully deleted the data.", nil)
	return
}
