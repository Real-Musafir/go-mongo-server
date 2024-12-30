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
	FindOneByKey(key string, value interface{}, ctx mongo.SessionContext) (interface{}, error)
	Update(id string, data interface{}, ctx mongo.SessionContext) (interface{}, error)
	Delete(id string, ctx mongo.SessionContext) (interface{}, error)
	FindAll(filter interface{}, ctx mongo.SessionContext) ([]map[string]interface{}, error)
	Aggregate(pipelines mongo.Pipeline, ctx mongo.SessionContext) ([]map[string]interface{}, error)
	
}

type MongoRepository struct {
	collection *mongo.Collection
}

func getSessionContext(sesionContext mongo.SessionContext) mongo.SessionContext {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	cont := context.Background()
	if sesionContext == nil {
		return mongo.NewSessionContext(cont, mongo.SessionFromContext(cont))
	}
	return sesionContext
	
}

func (mr MongoRepository) Create(data interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	result, err := mr.collection.InsertOne(sessionContext, data)
	return result, err
}

func (mr MongoRepository) FindOne(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
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

func (mr MongoRepository) FindOneByKey(key string, value interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	objId, err := primitive.ObjectIDFromHex(value.(string))

	var result *mongo.SingleResult

	if err != nil {
		result = mr.collection.FindOne(sessionContext, bson.M{key: value})
	}else {
		result = mr.collection.FindOne(sessionContext, bson.M{key: objId})
	}
	
	var  document map[string]interface{}
	
	if err := result.Decode(&document); err != nil {
		return nil, err
	}

	// fmt.Println("%s Check the documnt", document)

	return document, nil
}


func (mr MongoRepository) Update(id string, data interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	objectId, _ := primitive.ObjectIDFromHex(id)
	res, err := mr.collection.UpdateOne(sessionContext, bson.M{"_id":objectId}, data)
	return res, err
}

func (mr MongoRepository) Delete(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	objectId, _ := primitive.ObjectIDFromHex(id)
	res, err := mr.collection.DeleteOne(sessionContext, bson.M{"_id":objectId})
	return res, err
}

func (mr MongoRepository) FindAll(filter interface{}, ctx mongo.SessionContext) ([]map[string]interface{}, error) {
	sessionContext := getSessionContext(ctx)
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
	sessionContext := getSessionContext(ctx)

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
