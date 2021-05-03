package main

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoImage struct {
	ID          bson.ObjectId `bson:"_id"`
	Author      string        `bson:"author"`
	Caption     string        `bson:"caption"`
	ContentType string        `bson:"contentType"`
	DateTime    string        `bson:"dateTime"`
	FileID      bson.ObjectId `bson:"fileID"`
	FileSize    int64         `bson:"fileSize"`
	Height      int           `bson:"height"`
	Name        string        `bson:"name"`
	Width       int           `bson:"width"`
}

var dbSession *mgo.Session
var db *mgo.Database

func main() {
	var err error

	/*
	 * Connect to the server and get a database handle
	 */
	if dbSession, err = mgo.Dial("localhost:27017"); err != nil {
		panic(err)
	}

	db = dbSession.DB("test")
	defer dbSession.Close()

	/*
	 * Create the Echo server and register handlers for our endpoints
	 */
	httpServer := echo.New()
	httpServer.GET("/image/:id", streamImage)

	/*
	 * Start the server in a goroutine
	 */
	go func() {
		err = httpServer.Start(":8080")

		if err != nil {
			if err != http.ErrServerClosed {
				panic("Error starting server! " + err.Error())
			} else {
				fmt.Printf("Shutting down...")
			}
		}
	}()

	/*
	 * Setup shutdown channel so we can wait for CTRL+C to shut
	 * down the HTTP server
	 */
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = httpServer.Shutdown(ctx); err != nil {
		panic("Server did not shut down before timeout: " + err.Error())
	} else {
		fmt.Printf("Server shutdown")
	}
}

func streamImage(ctx echo.Context) error {
	var err error
	var gridFile *mgo.GridFile
	var imageRecord *MongoImage

	/*
	 * Get the record by ID
	 */
	id := bson.ObjectIdHex(ctx.Param("id"))

	if err = db.C("images").FindId(id).One(&imageRecord); err != nil {
		fmt.Printf("Error getting image by ID: %s\n", err.Error())
		return ctx.String(http.StatusNotFound, "Image not found")
	}

	/*
	 * Get the file from GridFS
	 */
	fileID := imageRecord.FileID

	if gridFile, err = db.GridFS("imagefiles").OpenId(fileID); err != nil {
		fmt.Printf("Error getting file from GridFS: %s\n", err.Error())
		return ctx.String(http.StatusInternalServerError, "Error getting file from database")
	}

	defer gridFile.Close()

	/*
	 * Set headers to tell the caller how big the file is and stream it back
	 */
	ctx.Response().Header().Set("Content-Length", strconv.Itoa(int(gridFile.Size())))
	ctx.Response().Header().Set("Content-Disposition", "inline; filename=\""+gridFile.Name()+"\"")

	return ctx.Stream(http.StatusOK, mime.TypeByExtension(filepath.Ext(gridFile.Name())), gridFile)
}
