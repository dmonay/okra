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

	doWorkResource := &DoWorkResource{
		mongo: mongodb,
	}

	r := gin.New()

	// middlewares
	r.Use(cors.Middleware(cors.Options{}))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// user
	r.POST("/register", doWorkResource.Register)
	r.POST("/login", doWorkResource.Login)
	r.POST("/logout", doWorkResource.Logout)

	// users
	r.GET("/get/users/all/:user", doWorkResource.GetAllUsers)

	// orgs
	r.POST("/create/organization", doWorkResource.CreateOrg)
	r.POST("/update/organization", doWorkResource.UpdateOrgName)
	r.GET("/get/orgs/all/:userid", doWorkResource.GetAllOrgs)
	r.GET("/get/orgs/members/:org", doWorkResource.GetMembers)

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
	r.POST("/update/kr/properties/:organization/:treeid/:objective/:kr", doWorkResource.UpdateKrProperties)

	// tasks
	r.POST("/create/task/:organization/:objective/:kr", doWorkResource.CreateTask)
	r.POST("/update/task/properties/:organization/:treeid/:objective/:kr/:task", doWorkResource.UpdateTaskProperties)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.SvcHost
	}
	fmt.Println("port: ", port)
	r.Run(":" + port)

	return nil
}

func InitMongo() (*mgo.Database, error) {

	uri := os.Getenv("MONGOHQ_URL")
	fmt.Println("uri: ", uri)
	if uri == "" {
		fmt.Println("\x1b[31;1mno connection string provided\x1b[0m")
		os.Exit(1)
	}
	db_name := os.Getenv("MONGOHQ_DB")
	fmt.Println("db: ", db_name)
	if db_name == "" {
		fmt.Println("\x1b[31;1mno db name provided\x1b[0m")
		os.Exit(1)
	}
	url := uri + "/" + db_name
	fmt.Println("uri: ", url)
	session, err := mgo.Dial(url)
	if err != nil {
		colorMsg := "\x1b[31;1mMongo connection failed\x1b[0m"
		log.Fatalln(colorMsg, err)
	}
	defer session.Close()

	db := session.DB(db_name)

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
