package repository

import (
	"context"
	"fmt"
	_struct "github.com/PutskouDzmitry/golang-training-Library/pkg/struct"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)


//ReadAll output all data with table books
func (b BookData) ReadAll() ([]_struct.Book, error){
	db := b.collection.Database("book")
	collection := db.Collection("book")
	var books []_struct.Book
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book _struct.Book
		err = cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		logrus.Debug(book)
		books = append(books, book)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, mongo.ErrNoDocuments
	}
	return books, nil
}

//Read read data in db
func (b BookData) Read(id string) (_struct.Book, error){
	db := b.collection.Database("book")
	collection := db.Collection("book")
	var book _struct.Book
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book, err
	}
	cur, err := collection.Find(ctx, bson.M{"_id": bookId})
	defer cur.Close(ctx)
	cur.Next(ctx)
	if err = cur.Decode(&book); err != nil {
		return book, err
	}
	return _struct.Book{}, nil
}

//Add add data in db
func (B BookData) Add(book _struct.Book) (string, error) {
	db := B.collection.Database("book")
	collection := db.Collection("book")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, book)
	if err != nil {
		logrus.Fatal(err)
	}
	if result == nil {
		return "", fmt.Errorf("Error with add data")
	}
	return book.BookId.String(), nil
}

//Update update number of books by the id
func (B BookData) Update(id string, value int) (string, error) {
	db := B.collection.Database("book")
	collection := db.Collection("book")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	filter := bson.D{{"_id", idObj}}
	update := bson.D{
		{"$set", bson.D{
			{"number", value},
		}},
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", fmt.Errorf("error with update data", err)
	}
	return "", nil
}

func (B BookData) Delete(id string) error {
	db := B.collection.Database("book")
	collection := db.Collection("book")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_ , err = collection.DeleteOne(ctx, bson.M{"_id": idObj})
	if err != nil {
		return err
	}
	return nil
}


//String output data in console
//func (B Book) String() string {
//	return fmt.Sprintln(B.BookId, B.AuthorId, B.PublisherId, strings.TrimSpace(B.NameOfBook), B.YearOfPublication, B.BookVolume, B.Number)
//}

//BookData create a new connection
type BookData struct {
	collection *mongo.Client
}

//NewBookData it's imitation constructor
func NewBookData(collection *mongo.Client) *BookData {
	return &BookData{collection: collection}
}
