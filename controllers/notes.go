package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type Note struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Title     string        `bson:"title" json:"title"`
	Done      bool          `bson:"done" json:"done"`
	CreatedAt time.Time     `bson:"created_at json:"created_at"`
}

const (
	ERROR_INSERTING_NOTE = "There was a problem inserting the note"
	ERROR_GETTING_NOTES  = "There was a problem getting all notes"
	ERROR_UPDATING_NOTE  = "There was a problem updating the note"
)

func HandleNoteCreate(db *mgo.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		var note Note
		if err := c.Bind(&note); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "invalid request",
			})
		}
		note.ID = bson.NewObjectId()
		note.Done = false
		note.CreatedAt = time.Now()
		if err := db.C("notes").Insert(&note); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ERROR_INSERTING_NOTE,
			})
		}
		c.JSON(http.StatusOK, note)
	}
}

func HandleNoteGetAll(db *mgo.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		var notes []Note
		if err := db.C("notes").Find(bson.M{}).All(&notes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ERROR_GETTING_NOTES,
			})
		}
		c.JSON(http.StatusOK, notes)

	}
}

func HandleNoteDone(db *mgo.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		var note Note
		note.ID = bson.ObjectIdHex(c.Param("id"))

		if err := db.C("notes").FindId(note.ID).One(&note); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ERROR_UPDATING_NOTE,
			})
			return
		}

		note.Done = true

		if err := db.C("notes").UpdateId(note.ID, &note); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ERROR_UPDATING_NOTE,
			})
			return
		}
		c.JSON(http.StatusOK, note)
	}
}

func HandleNoteUndone(db *mgo.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		var note Note
		note.ID = bson.ObjectIdHex(c.Param("id"))

		if err := db.C("notes").FindId(note.ID).One(&note); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ERROR_UPDATING_NOTE,
			})
			return
		}

		note.Done = false

		if err := db.C("notes").UpdateId(note.ID, &note); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ERROR_UPDATING_NOTE,
			})
		}
		c.JSON(http.StatusOK, note)
	}
}
