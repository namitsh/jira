package issues

import (
	"github.com/gin-gonic/gin"
	"jira-management/pkg/models/req_res"
	"jira-management/pkg/services/issues"
	"log"
	"net/http"
)

type IssueController interface {
	CreateIssue(ctx *gin.Context)
	GetIssue(ctx *gin.Context)
	GetAllIssues(ctx *gin.Context)
	UpdateIssue(ctx *gin.Context)
}

type issueController struct {
	service issues.IssueService
}

func New(svc issues.IssueService) IssueController {
	return &issueController{
		service: svc,
	}
}
func (c *issueController) CreateIssue(ctx *gin.Context) {

	//step 1 : put the values in the createIssueRequest struct and validate the values
	// step 2 : if not validated correctly , return error with 400 request
	// step 3 : send the create share request struct to the service
	// step 4 : send the response
	log.Println("In create issue controller")
	var issueRequest req_res.CreateIssueRequest
	// unmarshall the value
	if err := ctx.ShouldBindJSON(&issueRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// go to service
	//var *issue models.Issue
	issue, err := c.service.CreateIssue(&issueRequest)
	if err != nil {
		log.Println("Error occurred while creating issue in service module")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": issue,
	})
	return
}
func (c *issueController) GetIssue(ctx *gin.Context) {
	//	get the id of issue
	issueId := ctx.Param("issueId")
	log.Printf("Fetching the issue details for id: %s", issueId)
	issue, err := c.service.GetIssue(issueId)
	if err != nil {
		log.Printf("Error fetching getting issue: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, issue)
}
func (c *issueController) GetAllIssues(ctx *gin.Context) {
	//
	log.Println("In issue controller")
	issues, err := c.service.GetIssues()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, issues)
}
func (c *issueController) UpdateIssue(ctx *gin.Context) {
	log.Println("Patch method called for issue")
	issueId := ctx.Param("issueId")

	var req *req_res.PatchIssueRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if (*req == req_res.PatchIssueRequest{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No data provided to update",
		})
		return
	}

	log.Printf("Updating the issue details for id: %s", issueId)
	issue, err := c.service.UpdateIssue(issueId, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, issue)
}
