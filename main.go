package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type Person struct{
	Name	string
	age		int	
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
		
	router.GET("/mark", func(c *gin.Context) {
		sess := getSession()
		
		defer sess.Close()
		sess.setSafe(&mgo.Safe{})
		err = collection.Insert(&Person{"Stefan Klaste", 2},
	                        &Person{"Nishant Modak", 5},
	                        &Person{"Prathamesh Sonpatki", 5},
	                        &Person{"murtuza kutub", 5},
	                        &Person{"aniket joshi", 6},
	                        &Person{"Michael de Silva", 8},
	                        &Person{"Alejandro Cespedes Vicente", 5})
        if err != nil {
                log.Fatal("Problem inserting data: ", err)
                return
        }
		result := Person{}
        err = collection.Find(bson.M{"name": "Prathamesh Sonpatki"}).One(&result)
        if err != nil {
                log.Fatal("Error finding record: ", err)
                return
        }
		c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("fsdf",result.Name))))
		 
	})

	router.Run(":" + port)
}
func getSession(){
	// Connect to our local mongo
	s, err := mgo.Dial(" mongodb://<heroku_1pkmmgpc>:<mongopassword1>@ds041556.mlab.com:41556/heroku_1pkmmgpc")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
