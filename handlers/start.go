package handlers

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-cors"
	"gopkg.in/mgo.v2"
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

	// initialize mongo
	mongodb, err := InitMongo()
	if err != nil {
		colorMsg := "\x1b[31;1mMongoDB failed to initialize\x1b[0m"
		log.Fatalln(colorMsg, err)
		return err
	}
	// defer mongodb.Close()

	doWorkResource := &DoWorkResource{mongo: mongodb}

	r := gin.New()

	// middlewares
	r.Use(cors.Middleware(cors.Options{}))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// user
	r.POST("/register", doWorkResource.Register)
	r.POST("/login", doWorkResource.Login)
	r.POST("/logout", doWorkResource.Logout)

	// orgs
	r.POST("/create/organization", doWorkResource.CreateOrg)
	r.GET("/get/orgs/:userid", doWorkResource.GetAllOrgs)

	// trees
	r.POST("/create/tree/:organization", doWorkResource.CreateTree)
	r.POST("/update/tree/name/:organization/:treeid", doWorkResource.UpdateTreeName)
	r.GET("/get/trees/:organization/:treeid", doWorkResource.GetTree)
	r.GET("/get/trees/:organization", doWorkResource.GetAllTrees)

	// mission
	r.POST("/update/mission/:organization", doWorkResource.UpdateMission)

	// members
	r.POST("/update/members/:organization", doWorkResource.AddMembers)
	r.DELETE("/update/members/:organization", doWorkResource.DeleteMembers)

	// objectives and key results
	r.POST("/create/objective/:organization", doWorkResource.CreateObjective)
	r.POST("/update/objective/properties/:organization/:treeid/:objective", doWorkResource.UpdateObjProperties)
	r.POST("/create/kr/:organization/:objective", doWorkResource.CreateKeyResult)
	r.POST("/update/objective/properties/:organization/:treeid/:objective/:kr", doWorkResource.UpdateKrProperties)

	// tasks
	r.POST("/create/kr/:organization/:objective/:kr", doWorkResource.CreateTask)

	r.Run(cfg.SvcHost)

	return nil
}

func InitMongo() (*mgo.Database, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		colorMsg := "\x1b[31;1mMongo connection failed\x1b[0m"
		log.Fatalln(colorMsg, err)
	}

	db := session.DB("testing")

	return db, nil
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
			fmt.Println("\x1b[32;1mYou've started the server. Rejoice!\x1b[0m")

			if err = Run(cfg); err != nil {
				log.Fatal(err)
			}
		},
	},
}
