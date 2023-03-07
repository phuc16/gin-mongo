package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"

	models "gin-mongo/src/models"
	routes "gin-mongo/src/routes"
)

func main() {
	yamlFile, _ := os.Open("config.yaml")
	defer yamlFile.Close()

	byteValue, _ := ioutil.ReadAll(yamlFile)

	var config models.Config
	yaml.Unmarshal(byteValue, &config)
	log.Println(config)

	router := gin.Default()

	routes.UserRoutes(router)

	router.Run(":" + config.Port)

}
