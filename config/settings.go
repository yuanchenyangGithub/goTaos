package configs

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	config      *Config
	configMutex sync.Mutex
)

type Config struct {
	TaosDb   taosDB   `yaml:"taosDB"`
	KafkaSet kafkaSet `yaml:"kafkaSet"`
}

type taosDB struct {
	HostName   string `yaml:"hostName"`
	ServerPort string `yaml:"serverPort"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DbName     string `yaml:"dbName"`
}

type kafkaSet struct {
	KafkaHost          []string `yaml:"kafkaHost"`
	KafkaPort          string   `yaml:"kafkaPort"`
	KafkaTopicHBase    []string `yaml:"kafkaTopicHBase"`
	KafkaTopicLocation []string `yaml:"kafkaTopicLocation"`
}

func GetConfig() *Config {
	fmt.Println("======two======")
	filename, _ := filepath.Abs("config/settings-prod.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var config *Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	//可以从config.Env中访问env
	fmt.Printf("--- config:%v", config)
	return config
}

// func GetConfig() *Config {
// 	if config != nil {
// 		return config
// 	}
// 	configMutex.Lock()
// 	defer configMutex.Unlock()

// 	// double check
// 	if config != nil {
// 		return config
// 	}

// 	config = &Config{}
// 	envconfig.Process("appname", config)
// 	return config
// }
