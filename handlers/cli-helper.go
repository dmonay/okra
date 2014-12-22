package handlers

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	// "gopkg.in/mgo.v2"
	"github.com/tommy351/gin-cors"
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
	CheckErr(err, "MongoDB failed to initialize")
	// defer mongodb.Close()

	doWorkResource := &DoWorkResource{mongo: mongodb}

	r := gin.New()

	r.Use(cors.Middleware(cors.Options{}))

	// Global middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/register", doWorkResource.Register)
	r.POST("/login", doWorkResource.Login)
	// r.POST("/logout", handlers.Logout)
	r.POST("/create/organization", doWorkResource.CreateOrg)
	r.POST("/create/tree/:organization", doWorkResource.CreateTree)
	r.POST("/update/mission/:organization", doWorkResource.UpdateMission)
	r.POST("/update/members/:organization", doWorkResource.AddMembers)
	r.DELETE("/update/members/:organization", doWorkResource.DeleteMembers)
	r.POST("/update/objective/:organization", doWorkResource.UpdateObjective)
	r.POST("/create/objective/:organization/:objective", doWorkResource.CreateKeyResult)

	// r.POST("/get/trees/", doWorkResource.GetTrees)

	r.Run(cfg.SvcHost)

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
			fmt.Println("\x1b[32;1mYou've started the server. Rejoice!\x1b[0m")

			if err = Run(cfg); err != nil {
				log.Fatal(err)
			}
		},
	},
}
