package handlers

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/dmonay/do-work-api/common"
	// "github.com/dmonay/do-work-api/middleware"
	"github.com/gin-gonic/gin"
	// "gopkg.in/mgo.v2"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"os"
)

func getConfig(c *cli.Context) (common.Config, error) {
	yamlPath := c.GlobalString("config")
	config := common.Config{}

	if _, err := os.Stat(yamlPath); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(ymlData), &config)
	return config, err
}

func Run(cfg common.Config) error {

	//initialize mysql
	dbmap, err := InitSqlDb(cfg)
	if err != nil {
		return err
	}
	// defer dbmap.Db.Close()

	dbmap.SingularTable(true)

	// databaseUrl := "localhost:27017"
	// databseName := "testing"

	// initialize mongo
	mongodb, err := InitMongo()
	if err != nil {
		return err
	}

	// defer mongodb.Close()

	doWorkResource := &DoWorkResource{db: dbmap, mongo: mongodb}

	r := gin.New()

	// r.Use(middleware.DB(*databaseUrl, *databaseName))
	// r.Use(middleware.DB())

	// session, err := mgo.Dial("localhost:27017")
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()
	// db := session.DB("testing").C("testData")
	// err = db.Insert(&Person{"Ale", "+55 53 8116 9639"},
	// 	&Person{"Cla", "+55 53 8402 8510"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	r.POST("/register", doWorkResource.Register)
	r.POST("/login", doWorkResource.Login)
	// r.POST("/logout", handlers.Logout)
	r.POST("/create", doWorkResource.Create)

	r.Run(cfg.SvcHost)

	return nil
}

func Migrate(cfg common.Config) error {
	db, err := InitSqlDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	db.CreateTable(&common.User{})

	db.AutoMigrate(common.User{})
	return nil
}

var Commands = []cli.Command{
	{
		Name:  "server",
		Usage: "Run the http server",
		Action: func(c *cli.Context) {
			cfg, err := getConfig(c)
			if err != nil {
				log.Fatal(err)
				return
			}

			if err = Run(cfg); err != nil {
				log.Fatal(err)
			}
		},
	},
	{
		Name:  "migratedb",
		Usage: "Perform database migrations",
		Action: func(c *cli.Context) {
			cfg, err := getConfig(c)
			if err != nil {
				log.Fatal(err)
				return
			}

			fmt.Println("\x1b[32;1mYou've created the table 'user'!\x1b[0m")

			if err = Migrate(cfg); err != nil {
				log.Fatal(err)
			}
		},
	},
}