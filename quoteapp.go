package main

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
)

type QuoteForm struct {
	Quote string `form:"quote"`
}

var quotesBucket = []byte("quotes")

func main() {
	db, err := bolt.Open("quotes.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Static("/css", "./assets/css")
	router.StaticFile("/favicon.png", "./assets/favicon.png")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(context *gin.Context) {
		var quotes []string
		err = db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(quotesBucket)
			if bucket == nil {
				return errors.New("Bucket is empty")
			}
			cursor := bucket.Cursor()
			for k, _ := cursor.First(); k != nil; k, _ = cursor.Next() {
				quotes = append(quotes, fmt.Sprintf("%s", k))
			}
			return nil
		})
		if err != nil || len(quotes) == 0 {
			context.Redirect(http.StatusFound, "/add")
		}
		quote := quotes[rand.Intn(len(quotes))]
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Quotes",
			"quote": quote,
		})
	})

	router.GET("/manage", func(context *gin.Context) {
		var quotes []string
		err = db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(quotesBucket)
			if bucket == nil {
				return errors.New("Bucket is empty")
			}
			cursor := bucket.Cursor()
			for k, _ := cursor.First(); k != nil; k, _ = cursor.Next() {
				quotes = append(quotes, fmt.Sprintf("%s", k))
			}
			return nil
		})
		if err != nil || len(quotes) == 0 {
			context.Redirect(http.StatusFound, "/add")
		}
		context.HTML(http.StatusOK, "manage.tmpl", gin.H{
			"title":  "Manage Quotes",
			"quotes": quotes,
		})
	})

	router.GET("/add", func(context *gin.Context) {
		context.HTML(http.StatusOK, "add.tmpl", gin.H{
			"title": "Add New Quote",
		})
	})

	router.POST("/add", func(context *gin.Context) {
		var form QuoteForm
		if context.Bind(&form) == nil {
			err = db.Update(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists(quotesBucket)
				if err != nil {
					return err
				}
				if form.Quote == "" {
					return nil
				}
				err = bucket.Put([]byte(form.Quote), nil)
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				log.Fatal(err)
			}
		}
		context.Redirect(http.StatusFound, "/")
	})

	router.POST("/delete", func(context *gin.Context) {
		var form QuoteForm
		if context.Bind(&form) == nil {
			err = db.Update(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists(quotesBucket)
				if err != nil {
					return err
				}
				err = bucket.Delete([]byte(form.Quote))
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				log.Fatal(err)
			}
		}
		context.Redirect(http.StatusFound, "/manage")
	})

	router.Run(":3030")
}
