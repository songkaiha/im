package models

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Redis struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Password    string `yaml:"password"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"max_idle"`
	MaxActive   int    `yaml:"max_active"`
	IdleTimeout int    `yaml:"idle_timeout"`
}

type Base struct {
	Url string `yaml:"url"`
}

type Config struct {
	Redis `yaml:"Redis"`
	Base  `yaml:"Base"`
	Logs  string `yaml:"Logs"`
}

func (c Config) Print() {
	buffer, _ := yaml.Marshal(c)
	log.Println("[INFO] Current configuration: ^_^ \n\n", string(buffer))
}

var Conf *Config

func init() {
	if buffer, err := ioutil.ReadFile(fmt.Sprintf("configs/%s.yml", os.Getenv("ENV"))); err != nil {
		log.Fatalln("[ERROR] Models -> Config -> init: ", err.Error())
	} else {
		if err = yaml.Unmarshal(buffer, &Conf); err != nil {
			log.Fatalln("[ERROR] Models -> Config -> init -> format: ", err.Error())
		} else {
			path, _ := os.Getwd()
			Conf.Logs = filepath.Join(path, "logs")
			_ = os.MkdirAll(Conf.Logs, os.ModePerm)
		}
	}
}
