package projects

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jira-management/pkg/models"
	"jira-management/pkg/models/req_res"
	"jira-management/pkg/repository/projects"
	"log"
	"time"
)

type ProjectService interface {
	CreateProject(*req_res.CreateProjectRequest) (*models.Project, error)
	GetProject(id string) (*models.Project, error)
	GetProjects() ([]*models.Project, error)
	UpdateProject(string, *models.Project) (*models.Project, error)
}

type projectService struct {
	repo projects.ProjectRepository
}

func New(rep projects.ProjectRepository) ProjectService {
	return &projectService{
		repo: rep,
	}
}

func (svc *projectService) CreateProject(req *req_res.CreateProjectRequest) (*models.Project, error) {
	log.Println("Creating Project in service")
	project := &models.Project{}
	project.ID = primitive.NewObjectID()
	project.Name = req.Name
	project.Description = req.Description
	project.Lead = req.Lead
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := svc.repo.Save(ctx, project)
	if err != nil {
		log.Printf("Error creating a project: %s", err.Error())
		return nil, err
	}
	return project, nil
}
func (svc *projectService) GetProject(id string) (*models.Project, error) {
	projectId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	project, err := svc.repo.FindById(ctx, projectId)
	if err != nil {
		log.Printf("Error fetching the project for id: %s: %s", projectId, err.Error())
		return nil, err
	}
	return project, nil
}
func (svc *projectService) GetProjects() ([]*models.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	projects, err := svc.repo.FindAll(ctx)
	if err != nil {
		log.Printf("Error fetching the projects: %s", err.Error())
		return nil, err
	}
	return projects, nil
}

// UpdateProject TODO : this method is broken right now, need to change it.
func (svc *projectService) UpdateProject(id string, req *models.Project) (*models.Project, error) {
	projectId, _ := primitive.ObjectIDFromHex(id)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	project, err := svc.repo.Update(ctx, projectId)
	if err != nil {
		log.Printf("Error updating the document %s : %s", projectId, err.Error())
		return nil, err
	}
	return project, nil
}
