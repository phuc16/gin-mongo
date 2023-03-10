package utils

import (
	model "gin-mongo/src/models"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var ResCode = InitResCode()

func InitResCode() model.Code {
	yamlFile, _ := os.Open("config.yaml")
	defer yamlFile.Close()

	byteValue, _ := ioutil.ReadAll(yamlFile)

	var resCode model.ResponseCode
	yaml.Unmarshal(byteValue, &resCode)

	// log.Println(resCode)

	return resCode.ResCode
}
