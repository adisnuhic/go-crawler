package db

import (
	"context"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

// SetupArrangoDB create db and collection if doesn't exist
func SetupArrangoDB() driver.Database {

	/*
	 *	Create Connection to ArrangoDB
	 */
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		fmt.Println("NewConnection Error: ", err)
	}
	c, e := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "root"), // DB user & password
	})
	if e != nil {
		fmt.Println("NewClient Error: ", e)
	}

	/*
	 *	Open "crawler" Databse, note you need to make DB name "crawler" in ArrangoDB
	 */
	ctx := context.Background()
	db, err2 := c.Database(ctx, "crawler")
	if err2 != nil {
		fmt.Println("Error opening DB: ", err2)
	}

	/*
	 * Check if collection exist if not create a new one
	 */
	found, err3 := db.CollectionExists(ctx, "posts")
	if err3 != nil {
		fmt.Println("Collection error: ", err3)
	}

	if !found {
		options := &driver.CreateCollectionOptions{}
		_, err3 := db.CreateCollection(ctx, "posts", options)
		if err3 != nil {
			fmt.Println("Creating collection error: ", err3)
		}
	}

	return db // Return DB instance

}

// CreateDocument it will create document with the data and insert it into DB
func CreateDocument(doc map[string]interface{}, instance driver.Database) {
	/*
	 * Open DB collection
	 */
	ctx := context.Background()
	col, err := instance.Collection(ctx, "posts")
	if err != nil {
		fmt.Println("Error opening collection ", err)
	}

	/*
	 * Create document for that collection
	 */
	_, err2 := col.CreateDocument(ctx, doc)
	if err2 != nil {
		fmt.Println("Error creating document", err2)
	}
}
