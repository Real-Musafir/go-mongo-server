package dto

import (
	"time"
)

type CreateBookDto struct {
	BookName 		string 				`bson:"book_name" json:"book_name"`
	BookCode 		string 				`bson:"book_code" json:"book_code"`
	Description 	string 				`bson:"description,omitempty" json:"description,omitempty"` //omitempty --> this is not mendatroy
	Author 			string 				`bson:"author" json:"author"`
	Price 			float32 			`bson:"price" json:"price"`
	PageCount 		int 				`bson:"page_count" json:"page_count"`
	PublishedYear 	int 				`bson:"published_year" json:"published_year"`
	Status 			bool 				`bson:"status" json:"status"`
	CreatedBy 		string			 	`bson:"created_by" json:"created_by"`
	ModifiedBy 		string			 	`bson:"modified_by" json:"modified_by"`
	CreatedDate 	time.Time			`bson:"created_date" json:"created_date"`
	ModifiedDate 	time.Time			`bson:"modified_date" json:"modified_date"`
}

type UpdateBookDto struct {
	Id 				string			 	`bson:"_id" json:"_id"`
	BookName 		string 				`bson:"book_name" json:"book_name"`
	BookCode 		string 				`bson:"book_code" json:"book_code"`
	Description 	string 				`bson:"description,omitempty" json:"description,omitempty"` //omitempty --> this is not mendatroy
	Price 			float32 			`bson:"price" json:"price"`
	Status 			bool 				`bson:"status" json:"status"`
	CreatedBy 		string			 	`bson:"created_by" json:"created_by"`
	ModifiedBy 		string			 	`bson:"modified_by" json:"modified_by"`
	ModifiedDate 	time.Time			`bson:"modified_date" json:"modified_date"`
}