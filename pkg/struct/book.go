package _struct

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	BookId            primitive.ObjectID `bson:"_id"`
	AuthorId          int                `bson:"author_id"`
	BookVolume        int                `bson:"book_volume"`
	NameOfBook        string             `bson:"name_of_book"`
	Number            int                `bson:"number"`
	PublisherId       int                `bson:"publisher_id"`
	YearOfPublication string             `bson:"year_of_publication"`
}