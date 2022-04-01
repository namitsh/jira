package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"jira-management/configs"
	"jira-management/pkg/models"
	"jira-management/pkg/models/req_res"
	"log"
	"time"
)

type IssueRepository interface {
	Save(context.Context, *models.Issue) error
	FindAll(context.Context) ([]*models.Issue, error)
	Update(context.Context, primitive.ObjectID, *req_res.PatchIssueRequest) (*models.Issue, error)
	FindById(context.Context, primitive.ObjectID) (*models.Issue, error)
}

type issueRepository struct {
	client *mongo.Client
}

var issueCollection *mongo.Collection

func New(client *mongo.Client) IssueRepository {
	issueCollection = configs.GetCollection(client, "jira", "issues")
	return &issueRepository{
		client: client,
	}
}

func (is *issueRepository) Save(ctx context.Context, issue *models.Issue) error {
	inserted, err := issueCollection.InsertOne(ctx, issue)
	if err != nil {
		log.Printf("Error when creating issue: %s\n", err.Error())
		return err
	}
	log.Printf("Successfully inserted issue with id: %s", inserted.InsertedID)
	return nil

}
func (is *issueRepository) FindAll(ctx context.Context) ([]*models.Issue, error) {
	var issues []*models.Issue
	allBson := bson.M{}
	cursor, err := issueCollection.Find(ctx, allBson)
	if err != nil {
		log.Printf("Error when getting all issues: %s\n", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var issue models.Issue
		if err := cursor.Decode(&issue); err != nil {
			log.Println(err)
			return nil, err
		}
		issues = append(issues, &issue)
	}
	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return issues, nil
}

func (is *issueRepository) Update(ctx context.Context, id primitive.ObjectID, issue *req_res.PatchIssueRequest) (*models.Issue, error) {
	filter := bson.M{"_id": id}
	update := bson.M{}
	//update, err := bson.Marshal(&issue)
	//if err != nil {
	//	log.Printf("Error marshalling the document to bson: %s", err.Error())
	//	return nil, err
	//}
	if issue.Summary != nil {
		update["summary"] = issue.Summary
	}
	if issue.Description != nil {
		update["description"] = issue.Description
	}
	if issue.Status != nil {
		update["status"] = issue.Status
	}
	update["updatedat"], _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	log.Println(update)
	result, err := issueCollection.UpdateOne(ctx, filter,
		bson.M{"$set": update})
	if err != nil {
		log.Printf("Error when updating issue for ID %s: %s", id, err.Error())
		return nil, err
	}
	var updatedIssue *models.Issue
	if result.MatchedCount == 1 {
		err := is.client.Database("jira").Collection("issues").FindOne(ctx, filter).Decode(&updatedIssue)
		if err != nil {
			log.Printf("Error when getting the issue for ID %s : %s ", id, err.Error())
			return nil, err
		}
	}
	log.Println(updatedIssue)
	return updatedIssue, nil
}
func (is *issueRepository) FindById(ctx context.Context, id primitive.ObjectID) (*models.Issue, error) {
	var issue *models.Issue
	getBson := bson.M{"_id": id}
	err := is.client.Database("jira").Collection("issues").FindOne(ctx, getBson).Decode(&issue)
	if err == mongo.ErrNoDocuments {
		log.Printf("No Documents found for issue with ID: %s\n", id)
		return nil, err
	}
	if err != nil {
		log.Printf("Error occurred when getting issue for id: %s wrong %s\n", id, err.Error())
		return nil, err
	}
	log.Println(issue)
	return issue, nil
}
