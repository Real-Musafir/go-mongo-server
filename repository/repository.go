package repository

import "github.com/Real-Musafir/bookshop/config"

type Repository struct {
	UserRepository IMongoRepository
	BookRepository IMongoRepository
}


func GetRepository() *Repository {
	return &Repository{
		UserRepository: GetMongoRepository(config.GetEnvProperty("database_name"), "user"),
		BookRepository: GetMongoRepository(config.GetEnvProperty("database_name"), "book"),
	}
}