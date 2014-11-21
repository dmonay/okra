package database

import (
	"errors"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/dmonay/do-work-api/handlers"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"os"
)

func getConfig(c *cli.Context) (Config, error) {
	yamlPath := c.GlobalString("config")
	config := Config{}

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

func Run(cfg Config) error {

	//initialize mysql
	dbmap, err := InitDb(cfg)
	if err != nil {
		return err
	}
	defer dbmap.Db.Close()

	m := martini.Classic()

	//define the endpoints
	m.Post("/register", binding.Json(handlers.Credentials{}), handlers.Register)
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
