package config

import (
	"encoding/json"
	"github.com/chuy2001/gin-es/models"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
)

type Config struct {
	ApiURL       string   `json:"apiURL"`
	ElasticURL   string   `json:"elasticURL"`
	Indices      []string `json:"indices"`
	DefaultIndex string   `json:"defaultIndex"`
	SniffCluster bool     `json:"sniffCluster, omitempty"`
	UseAMQP      bool     `json:"useAMQP"`
	AmqpURL      string   `json:"amqpURL"`
	AmqpQueues   []string `json:"amqpQueues"`
	GinDebug     bool     `json:"ginDebug, omitempty"`
}

//GetConfFromJSONFile reads application configuration from *filename* and maps it to a Config struct
func GetConfFromJSONFile(filename string) *Config {

	confContent, err := ioutil.ReadFile(filename)
	models.CheckFatalError(err)
	config := &Config{}
	err = json.Unmarshal(confContent, config)
	models.CheckFatalError(err)

	// If no DefaultIndex value is given, use the first Index value
	if config.DefaultIndex == "" {
		config.DefaultIndex = config.Indices[0]
	}

	log.Println("Configuration loaded:")
	spew.Config.Indent = "\t"
	spew.Dump(*config)
	return config
}
