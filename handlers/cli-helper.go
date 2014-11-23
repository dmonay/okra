package handlers

import (
	"errors"
	// "fmt"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	// "github.com/coopernurse/gorp"
	"github.com/dmonay/do-work-api/common"
	// "github.com/dmonay/do-work-api/handlers"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	// "net/http"
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
	dbmap, err := InitDb(cfg)
	if err != nil {
		return err
	}
	// defer dbmap.Db.Close()

	doWorkResource := &DoWorkResource{db: dbmap}

	r := gin.Default()

	r.POST("/register", doWorkResource.Register)
	// r.POST("/login", handlers.Login)
	// r.POST("/logout", handlers.Logout)

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

			if err = Run(cfg); err != nil {
				log.Fatal(err)
			}
		},
	},
	// {
	// 	Name:  "migratedb",
	// 	Usage: "Perform database migrations",
	// 	Action: func(c *cli.Context) {
	// 		cfg, err := getConfig(c)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 			return
	// 		}

	// 		svc := service.TodoService{}

	// 		if err = svc.Migrate(cfg); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	},
	// },
}
