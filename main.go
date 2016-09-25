package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"Heroku-Go_FirstDeployedApp\models"
)
UserController struct {  
    session *mgo.Session
}
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	
	uc := controllers.NewUserController(getSession())

	router.GET("/mark", func(c *gin.Context) {
		u := models.User{
	          Name:   "Bob Smith",
	          Gender: "male",
	          Age:    50,
	          Id:     p.ByName("id"),
        }
		c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**hi!**"))))
	})

	router.Run(":" + port)
}
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial(" mongodb://<heroku_1pkmmgpc>:<mongopassword1>@ds041556.mlab.com:41556/heroku_1pkmmgpc")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
