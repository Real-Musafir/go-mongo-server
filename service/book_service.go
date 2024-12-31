package service

import "go.mongodb.org/mongo-driver/mongo"

type IBookService interface {
	CreateBook(data interface{}, sessionContext mongo.SessionContext) (interface{}, error)
}