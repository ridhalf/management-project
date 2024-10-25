package controller

import (
	"github.com/gin-gonic/gin"
	"management-project/model/api"
	"management-project/model/constants"
	"management-project/model/domain"
	"net/http"
)

func AllowReadProject(ctx *gin.Context) bool {
	user := ctx.MustGet("user").(domain.User)
	if user.Role != constants.PROJECT_MANAGER {
		ValidateMessage(ctx)
		return false
	}
	return true
}
func ValidateMessage(ctx *gin.Context) {
	response := api.APIResponse("sorry, but you do not have access to this area. Contact support for further assistance", http.StatusForbidden, "Forbidden", nil)
	ctx.JSON(http.StatusForbidden, response)
	return
}
func HandleBindError(ctx *gin.Context) {
	response := api.APIResponse("failed to bind request", http.StatusBadRequest, "BadRequest", nil)
	ctx.JSON(http.StatusBadRequest, response)
}
func HandleServiceError(ctx *gin.Context, err error) {
	response := api.APIResponse(err.Error(), http.StatusBadRequest, "BadRequest", nil)
	ctx.JSON(http.StatusBadRequest, response)
}
func HandleRequestSuccess(ctx *gin.Context, msg string, data interface{}) {
	apiResponse := api.APIResponse(msg, http.StatusOK, "Success", data)
	ctx.JSON(http.StatusOK, apiResponse)
}
func HandleGenerateTokenError(ctx *gin.Context) {
	response := api.APIResponse("generate token is failed", http.StatusBadRequest, "BadRequest", nil)
	ctx.JSON(http.StatusBadRequest, response)
}
