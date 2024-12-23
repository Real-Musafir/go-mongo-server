package repository

import (
	"context"

	"github.com/Real-Musafir/bookshop/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMongoRepository interface {
	Create(data interface{}, ctx mongo.SessionContext) (interface{}, error)
	FindOne(id string, ctx mongo.SessionContext) (interface{}, error)
	Update(id string, data interface{}, ctx mongo.SessionContext) (interface{}, error)
	Delete(id string, ctx mongo.SessionContext) (interface{}, error)
	FindAll(filter interface{}, ctx mongo.SessionContext) ([]map[string]interface{}, error)
	Aggregate(pipelines mongo.Pipeline, ctx mongo.SessionContext) ([]map[string]interface{}, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

func getUpSessionContext(sesionContext mongo.SessionContext) mongo.SessionContext {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if sesionContext == nil {
		return mongo.NewSessionContext(ctx, mongo.SessionFromContext(ctx))
	}
	return sesionContext
	
}

func (mr MongoRepository) Create(data interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getUpSessionContext(ctx)
	result, err := mr.collection.InsertOne(sessionContext, data)
	return result, err
}

func (mr MongoRepository) FindOne(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getUpSessionContext(ctx)
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	result := mr.collection.FindOne(sessionContext, bson.M{"_id": objectId})
	var  document map[string]interface{}
	if err := result.Decode(document); err != nil {
		return nil, err
	}

	return document, nil
}


func (mr MongoRepository) Update(id string, data interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getUpSessionContext(ctx)
	objectId, _ := primitive.ObjectIDFromHex(id)
	res, err := mr.collection.UpdateOne(sessionContext, bson.M{"_id":objectId}, data)
	return res, err
}

func (mr MongoRepository) Delete(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getUpSessionContext(ctx)
	objectId, _ := primitive.ObjectIDFromHex(id)
	res, err := mr.collection.DeleteOne(sessionContext, bson.M{"_id":objectId})
	return res, err
}

func (mr MongoRepository) FindAll(filter interface{}, ctx mongo.SessionContext) ([]map[string]interface{}, error) {
	sessionContext := getUpSessionContext(ctx)
	cursor, err := mr.collection.Find(sessionContext, filter)

	if err!= nil {
		return nil, err
	}

	defer cursor.Close(sessionContext)

	var results []map[string]interface{}
	for cursor.Next(sessionContext) {
		var document map[string]interface{}
		if err:= cursor.Decode(&document); err != nil {
			return nil, err
		}
		results = append(results, document)
	}

	return results, cursor.Err()
}

func (mr MongoRepository) Aggregate(pipelines mongo.Pipeline, ctx mongo.SessionContext) ([]map[string]interface{}, error) {
	sessionContext := getUpSessionContext(ctx)

	cursor, err := mr.collection.Aggregate(sessionContext, pipelines)

	if err!= nil {
		return nil, err
	}

	defer cursor.Close(sessionContext)

	var results []map[string]interface{}
	for cursor.Next(sessionContext) {
		var document map[string]interface{}
		if err:= cursor.Decode(&document); err != nil {
			return nil, err
		}
		results = append(results, document)
	}

	return results, cursor.Err()

}

func GetMongoRepository(dbName string, collectionName string) IMongoRepository {
	collection := config.GetDatabaseCollection(&dbName, collectionName)
	return &MongoRepository {
		collection: collection,
	}
}
