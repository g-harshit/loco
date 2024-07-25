package conf

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/jacobstr/confer"
)

var configuration *confer.Config

func init() {
	//LoadDefaultConfig()
	LoadConfigBasedEnv()
}

//LoadConfigBasedEnv : Load Config Based On Env
func LoadConfigBasedEnv() {
	confFile := "dev.yaml"

	configuration = confer.NewConfig()

	err := configuration.ReadPaths(confFile)

	if err == nil {
		abs, _ := filepath.Abs(confFile)
		fmt.Println("Configuration loaded:", abs)
	}

	if err != nil {
		log.Println("No configuration file found")
	}
}

//String will return string value of the given key
func String(key string, defValue string) string {
	if Exists(key) {
		return configuration.GetString(key)
	}
	return defValue
}

//Exists will check if key exists in conf
func Exists(key string) bool {
	return configuration.IsSet(key)
}
