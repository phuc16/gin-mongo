package main

import (
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	models "gin-mongo/src/models"
	routes "gin-mongo/src/routes"

	middlewares "gin-mongo/src/middlewares"
)

var Global = "aaa"

func main() {
	log := logrus.New()
	gin.SetMode(gin.ReleaseMode)

	if err := os.MkdirAll("log", os.ModePerm); err != nil {
		log.Error(err)
	}

	f1, err := os.OpenFile("log/activity1.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	f2, err := os.OpenFile("log/activity2.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Println("error opening")
	}

	defer f2.Close()
	defer f1.Close()

	log.SetOutput(f1)

	yamlFile, _ := os.Open("config.yaml")
	defer yamlFile.Close()

	byteValue, _ := ioutil.ReadAll(yamlFile)

	var config models.Config
	yaml.Unmarshal(byteValue, &config)

	router := gin.Default()

	router.Use(gin.LoggerWithWriter(f2), gin.Recovery()) //middleware of gin-gonic
	router.Use(middlewares.Logger(log), gin.Recovery())

	routes.UserRoutes(router)

	router.Run(":" + config.Port)
}
