package database

import (
	"errors"
	// "fmt"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	// "github.com/coopernurse/gorp"
	"github.com/dmonay/do-work-api/common"
	"github.com/dmonay/do-work-api/handlers"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"net/http"
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
	defer dbmap.Db.Close()

	m := martini.Classic()

	m.Post("/register", binding.Json(handlers.Credentials{}), func(attr handlers.Credentials, erro binding.Errors) (int, string) {
		uname := attr.Username

		query := &common.User{0, uname, attr.Password}
		err = dbmap.Insert(query)

		// pwd := attr.Password
		return http.StatusOK, handlers.JsonString(handlers.SuccessMsg{"You have successfully registered, " + uname})
	})

	m.Post("/login", binding.Json(handlers.Credentials{}), handlers.Login)
	m.Post("/logout", binding.Json(handlers.Credentials{}), handlers.Logout)

	m.Run()

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
