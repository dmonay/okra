package middleware

// import (
// 	"github.com/gin-gonic/gin"
// 	"gopkg.in/mgo.v2"
// 	// "gopkg.in/mgo.v2/bson"
// 	"fmt"
// 	"log"
// )

// type Person struct {
// 	Name  string
// 	Phone string
// }

// func DB() gin.HandlerFunc {
// 	session, err := mgo.Dial("localhost:27017")
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Clean up our session.
// 	defer session.Close()

// 	return func(c *gin.Context) {
// 		// Clone the session.
// 		fmt.Println("I ran once")
// 		s := session.Clone()

// 		// // Map a reference to the database to the request context
// 		// c.Map(s.DB(databaseName))

// 		db := s.DB("testing").C("testData")
// 		err = db.Insert(&Person{"Ale", "+55 53 8116 9639"},
// 			&Person{"Cla", "+55 53 8402 8510"})
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("I ran twice")
// 		// Pass control to next handler
// 		c.Next()
// 		fmt.Println("I ran a third time")
// 	}
// }
