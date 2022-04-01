package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"jira-management/configs"
	"jira-management/pkg/controllers"
	"jira-management/pkg/repository"
	"jira-management/pkg/services"
	"log"
	"net/http"
)

var (
	client          *mongo.Client               = configs.DB
	issueRepository repository.IssueRepository  = repository.New(client)
	issueService    services.IssueService       = services.New(issueRepository)
	issueController controllers.IssueController = controllers.New(issueService)
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
	issueRoutes := router.Group("/issues")
	{
		issueRoutes.POST("/", issueController.CreateIssue)
		issueRoutes.GET("/", issueController.GetAllIssues)
		issueRoutes.GET("/:issueId", issueController.GetIssue)
		issueRoutes.PATCH("/:issueId", issueController.UpdateIssue)
	}
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error running the server : %s", err.Error())
	}
}
