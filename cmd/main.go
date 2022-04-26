package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"jira-management/configs"
	issues3 "jira-management/pkg/controllers/issues-controller"
	"jira-management/pkg/controllers/projects-controller"
	"jira-management/pkg/repository/issues"
	projects2 "jira-management/pkg/repository/projects"
	issues2 "jira-management/pkg/services/issue-service"
	projects3 "jira-management/pkg/services/project-service"
	"log"
	"net/http"
)

var (
	client            *mongo.Client               = configs.DB
	issueRepository   issues.IssueRepository      = issues.New(client)
	projectRepository projects2.ProjectRepository = projects2.New(client)

	issueService   issues2.IssueService     = issues2.New(issueRepository)
	projectService projects3.ProjectService = projects3.New(projectRepository)

	issueController   issues3.IssueController               = issues3.New(issueService)
	projectController projects_controller.ProjectController = projects_controller.New(projectService)
)

func main() {
	fmt.Println("Hello World")
	router := gin.Default()
	// route testing
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	// TODO : create these routes in different modules , pkg/routes maybe
	issueRoutes := router.Group("/issues")
	{
		issueRoutes.POST("/", issueController.CreateIssue)
		issueRoutes.GET("/", issueController.GetAllIssues)
		issueRoutes.GET("/:issueId", issueController.GetIssue)
		issueRoutes.PATCH("/:issueId", issueController.UpdateIssue)
	}

	projectRoutes := router.Group("/projects")
	{
		projectRoutes.POST("/", projectController.CreateProject)
		projectRoutes.GET("/:projectId", projectController.GetProject)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error running the server : %s", err.Error())
	}
}
