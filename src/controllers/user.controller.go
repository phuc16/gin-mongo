package controller

import (
	models "gin-mongo/src/models"
	mongoDb "gin-mongo/src/mongoDb"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v3"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.Bind(&user); err != nil {
		log.Fatal(err)
	}

	err := mongoDb.CreateUser(user)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success"})
}

func GetAllUsers(c *gin.Context) {
	res, err := mongoDb.GetAllUsers()

	if err != nil {
		log.Fatal(err)
	}

	res2 := gin.H{"status": http.StatusOK, "message": "Success", "data": res}

	data, err := yaml.Marshal(&res2)

	err2 := ioutil.WriteFile("result.yaml", data, 0)

	if err2 != nil {
		log.Fatal(err2)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success", "data": res})
}

func GetUserById(c *gin.Context) {
	userId := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(userId)

	res, err := mongoDb.GetUserById(objId)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found"})
		return
	} else if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Success", "data": res})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(userId)

	var user models.User
	c.Bind(&user)

	res, err := mongoDb.UpdateUser(objId, user)

	if err != nil {
		log.Fatal(err)
	}

	if res < 1 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success"})
}

func DeleteUserById(c *gin.Context) {
	userId := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(userId)

	res, err := mongoDb.DeleteUserById(objId)

	if err != nil {
		log.Fatal(err)
	}

	if res < 1 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "success"})
}
