# MongoDB Tools

This package contains all necessary tools to establish connection and organize CRUD operations using MongoDB.

#### Example:
Connect to mongoDB, register gin endpoint, and create a gin middleware to create new item from POST request
```package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/michaelknyazev/go-tools/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemModel struct {
	ID    primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Label string             `json:"label,omitempty" bson:"label,omitempty"`
}

func getTestCollection() *mongo.Collection {
	return mongodb.GetCollection("test")
}

func main() {
	uri := os.Getenv("MONGO_URI")
	
	if err := mongodb.Init(uri); err != nil {
		panic(err)
	}

	if err := mongodb.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := mongodb.Disconnect(); err != nil {
			panic(err)
		}
	}()

	if err := mongodb.CreateTextIndex(getTestCollection(), []string{"label"}, false, "en"); err != nil {
		panic(err)
	}

	router := gin.Default()

	router.POST("/api/v1/create", createItem)
	router.GET("/api/v1/search", searchItems)
	router.Run(":8080")
}

func createItem(c *gin.Context) {
	var item ItemModel

	if err := c.BindJSON(&item); err != nil {
		c.JSON(400, map[string]interface{}{
			"result": "Failed to bind body to model",
		})
		c.Abort()
	}

	item.ID = primitive.NewObjectID()

	_, err := mongodb.CreateItem(getTestCollection(), item)

	if err != nil {
		c.JSON(500, map[string]interface{}{
			"result": "Failed to create item in Database",
		})
		c.Abort()
	}

	c.JSON(200, map[string]interface{}{
		"result": item,
	})
	c.Abort()
}

func searchItems(c *gin.Context) {
	query := c.DefaultQuery("query", "")

	result, err := mongodb.SearchTextIndex(getTestCollection(), query)

	if err != nil {
		c.JSON(500, map[string]interface{}{
			"result": "Failed to perform search operation"
		});
		c.Abort()
	}

	c.JSON(200, map[string]interface{}{
		"result": result
	})
	c.Abort()
}
```