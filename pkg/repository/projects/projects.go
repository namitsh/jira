package projects

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"jira-management/configs"
	"jira-management/pkg/models"
	"log"
)

type ProjectRepository interface {
	Save(context.Context, *models.Project) error
	FindById(context.Context, primitive.ObjectID) (*models.Project, error)
	FindAll(context.Context) ([]*models.Project, error)
	Update(context.Context, primitive.ObjectID) (*models.Project, error)
}

type projectRepository struct {
	client *mongo.Client
}

var projectCollection *mongo.Collection

func New(client *mongo.Client) ProjectRepository {
	projectCollection = configs.GetCollection(client, "jira", "projects")
	return &projectRepository{
		client: client,
	}
}

func (p *projectRepository) Save(ctx context.Context, project *models.Project) error {
	inserted, err := projectCollection.InsertOne(ctx, project)
	if err != nil {
		log.Printf("Error when creating issue:  %s", err.Error())
		return err
	}
	log.Printf("Successfully inserted project with id: %s", inserted.InsertedID)
	return nil
}
func (p *projectRepository) FindById(ctx context.Context, id primitive.ObjectID) (*models.Project, error) {
	var project *models.Project
	getBSON := bson.M{"_id": id}
	err := p.client.Database("jira").Collection("projects").FindOne(ctx, getBSON).Decode(&project)
	if err == mongo.ErrNoDocuments {
		log.Printf("No Documents found for project with ID: %s\n", id)
		return nil, err
	}
	if err != nil {
		log.Printf("Error occurred when getting project for id %s:  %s\n", id, err.Error())
		return nil, err
	}
	log.Println(project)
	return project, nil
}
func (p *projectRepository) FindAll(ctx context.Context) ([]*models.Project, error) {
	var projects []*models.Project
	allBSON := bson.M{}
	cursor, err := projectCollection.Find(ctx, allBSON)
	if err != nil {
		log.Printf("Error when getting all projects: %s\n", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var project models.Project
		if err := cursor.Decode(&project); err != nil {
			log.Println(err)
			return nil, err
		}
		projects = append(projects, &project)
	}
	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return projects, nil
}
func (p *projectRepository) Update(ctx context.Context, id primitive.ObjectID) (*models.Project, error) {
	return nil, nil
}
