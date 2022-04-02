package issues

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jira-management/pkg/models"
	"jira-management/pkg/models/req_res"
	"jira-management/pkg/repository/issues"
	"log"
	"time"
)

type issueService struct {
	repo issues.IssueRepository
}
type IssueService interface {
	CreateIssue(*req_res.CreateIssueRequest) (*models.Issue, error)
	GetIssue(id string) (*models.Issue, error)
	GetIssues() ([]*models.Issue, error)
	UpdateIssue(id string, req *req_res.PatchIssueRequest) (*models.Issue, error)
}

func New(rep issues.IssueRepository) IssueService {
	return &issueService{
		repo: rep,
	}
}

func (svc *issueService) CreateIssue(req *req_res.CreateIssueRequest) (*models.Issue, error) {
	log.Println("createIssue service")
	log.Println(req)
	issue := &models.Issue{}
	issue.IssueType = req.Type
	issue.Description = req.Description
	issue.Summary = req.Summary
	issue.ID = primitive.NewObjectID()
	issue.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	issue.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	issue.Name = "ISSUE-1"
	issue.Status = "New-Issue"
	log.Println(issue)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := svc.repo.Save(ctx, issue)
	if err != nil {
		log.Printf("Error occurred while saving the issue to database: %s", err.Error())
		return nil, err
	}
	return issue, nil
}
func (svc *issueService) GetIssue(id string) (*models.Issue, error) {
	//	convert into ObjectID
	issueId, _ := primitive.ObjectIDFromHex(id)
	//	send to repo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	issue, err := svc.repo.FindById(ctx, issueId)
	if err != nil {
		log.Printf("Error occurred in database while getting issue for ID %s: %s", issueId, err)
		return nil, err
	}
	return issue, nil

}
func (svc *issueService) GetIssues() ([]*models.Issue, error) {
	log.Println("Getting all issues.")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	issues, err := svc.repo.FindAll(ctx)
	if err != nil {
		log.Printf("Error occurred listing all issues: %s", err.Error())
		return nil, err
	}
	return issues, nil
}
func (svc *issueService) UpdateIssue(id string, req *req_res.PatchIssueRequest) (*models.Issue, error) {
	issueId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	issue, err := svc.repo.Update(ctx, issueId, req)
	if err != nil {
		return nil, err
	}
	return issue, nil
}
