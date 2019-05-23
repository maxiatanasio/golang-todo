package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"log"
	"noteloolvides/controllers"
)

func main() {
	db := ConnectToMongo()
	r := SetupRouter(db)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func SetupRouter(db *mgo.Database) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", controllers.HandlePing)
	r.POST("/note", controllers.HandleNoteCreate(db))
	r.GET("/note", controllers.HandleNoteGetAll(db))
	r.POST("/note/:id/done", controllers.HandleNoteDone(db))
	r.POST("/note/:id/undone", controllers.HandleNoteUndone(db))
	return r
}

func ConnectToMongo() *mgo.Database {
	session, err := mgo.Dial("mongodb:27017")
	if err != nil {
		log.Fatal(err)
	}
	return session.DB("noteloolvidesdev")
}
