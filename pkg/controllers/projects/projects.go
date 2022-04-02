package projects

import (
	"github.com/gin-gonic/gin"
	"jira-management/pkg/models/req_res"
	"jira-management/pkg/services/projects"
	"log"
	"net/http"
)

type ProjectController interface {
	CreateProject(ctx *gin.Context)
	GetProject(ctx *gin.Context)
}

type projectController struct {
	svc projects.ProjectService
}

func New(svc projects.ProjectService) ProjectController {
	return &projectController{
		svc: svc,
	}
}

func (c *projectController) CreateProject(ctx *gin.Context) {
	var projectRequest req_res.CreateProjectRequest
	if err := ctx.ShouldBindJSON(&projectRequest); err != nil {
		log.Printf("Error validating the create project request: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	project, err := c.svc.CreateProject(&projectRequest)
	if err != nil {
		log.Printf("Error occurred creating a project: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, project)
}

func (c *projectController) GetProject(ctx *gin.Context) {
	projectId := ctx.Param("projectId")
	project, err := c.svc.GetProject(projectId)
	if err != nil {
		log.Printf("Error fetching the project for id %s: %s", projectId, err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, project)
}
