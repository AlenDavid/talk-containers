package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Bread struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

var database = "mydb1"
var client *mongo.Client

func Server() *gin.Engine {
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	api := engine.Group("api")
	{

		api.GET("/breads", func(c *gin.Context) {
			var breads []Bread = make([]Bread, 0)
			breadsColl := client.Database(database).Collection("breads")

			cursor, err := breadsColl.Find(c, bson.M{})
			if err != nil {
				log.Println(err)
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			if err := cursor.All(c, &breads); err != nil {
				log.Println(err)
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": breads})
		})

		api.POST("/breads", func(c *gin.Context) {
			var bread Bread

			if err := c.BindJSON(&bread); err != nil {
				c.AbortWithError(http.StatusUnprocessableEntity, errors.New("unprocessable entity"))
				return
			}

			// self-generate id
			bread.Id = primitive.NewObjectID()

			breadsColl := client.Database(database).Collection("breads")
			if _, err := breadsColl.InsertOne(c, bread); err != nil {
				log.Println(err)
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			c.JSON(http.StatusCreated, gin.H{"data": bread})
		})
	}

	return engine
}

func getenv(key, fallback string) string {
	if found := os.Getenv(key); found != "" {
		return found
	}

	return fallback
}

func main() {
	ctx := context.Background()
	var err error

	clientOpts := options.Client().ApplyURI(getenv("MONGODB_URL", "mongodb://localhost:27017"))
	client, err = mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	defer func() { _ = client.Disconnect(ctx) }()

	if err := Server().Run(); err != nil {
		log.Fatal(err)
	}
}
