package repository

import (
	_struct "github.com/PutskouDzmitry/golang-training-Library/pkg/struct"
	"go.mongodb.org/mongo-driver/mongo"
)


//ReadAll output all data with table books
func (b BookData) ReadAll() ([]_struct.Book, error){
	//var books []Book
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//cursor, err := B.collection.Find(ctx, bson.M{})
	//if err != nil {
	//	return nil, err
	//}
	//defer cursor.Close(ctx)
	//for cursor.Next(ctx) {
	//	var book Book
	//	err = cursor.Decode(&book)
	//	if err != nil {
	//		return nil, err
	//	}
	//	logrus.Debug(book)
	//	books = append(books, book)
	//}
	//if err = cursor.Err(); err != nil {
	//	return nil, err
	//}
	//if len(books) == 0 {
	//	return nil, mongo.ErrNoDocuments
	//}
	//return books, nil
	return []_struct.Book{}, nil
}

//Read read data in db
func (b BookData) Read(id string) (_struct.Book, error){
	//var book Book
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//bookId, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	return book, err
	//}
	//cur, err := B.collection.Find(ctx, bson.M{"_id": bookId})
	//defer cur.Close(ctx)
	//cur.Next(ctx)
	//if err = cur.Decode(&book); err != nil {
	//	return book, err
	//}
	return _struct.Book{}, nil
}

//Add add data in db
func (B BookData) Add(book _struct.Book) error {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//result, err := B.collection.InsertOne(ctx, book)
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//if result == nil {
	//	return fmt.Errorf("Error with add data")
	//}
	return nil
}

//Update update number of books by the id
func (B BookData) Update(id string, value int) error {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//idObj, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	return err
	//}
	//filter := bson.D{{"_id", idObj}}
	//update := bson.D{
	//	{"$set", bson.D{
	//		{"number", value},
	//	}},
	//}
	//_, err = B.collection.UpdateOne(ctx, filter, update)
	//if err != nil {
	//	return fmt.Errorf("error with update data", err)
	//}
	return nil
}

func (B BookData) Delete(id string) error {
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//idObj, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	return err
	//}
	//_ , err = B.collection.DeleteOne(ctx, bson.M{"_id": idObj})
	//if err != nil {
	//	return err
	//}
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
