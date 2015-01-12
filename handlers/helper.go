package handlers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"log"
)

type ErrorMsg struct {
	Error string "json:'Error'"
}

type SuccessMsg struct {
	Success interface{} "json:'Success'"
}

type DoWorkResource struct {
	mongo *mgo.Database
}

func CheckErr(err error, msg string, c *gin.Context) {
	colorMsg := "\x1b[31;1m" + msg + "\x1b[0m"
	log.Println(colorMsg, err)
	c.JSON(400, ErrorMsg{msg})
}
