package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	VerifyToken string `yaml:"verify_token"`
	AccessToken string `yaml:"access_token"`
	AppSecret   string `yaml:"app_secret"`
}

func parseContentFile() string {
	contentFile, err := ioutil.ReadFile("content.yml")

	if err != nil {
		log.Printf("Error open content file: %s\n\n", err)
	}
	er, _ := yaml.Marshal(contentFile)

	if er != nil {
		log.Printf("Couldnt marshal content file: %s\n\n", er)
	}

	return string(er)
}

func (c *Config) readYml() *Config {
	yamlFile, err := ioutil.ReadFile("bot.config.yml")
	if err != nil {
		log.Printf("Error opening file: %s\n\n", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("Couldn't marsal config file: %s\n\n", err)
	}

	fmt.Printf("Here is the parsed content.yml: %s\n\n", c)
	return c
}

func getToken() string {
	var c Config
	c.readYml()
	v, err := json.Marshal(c)
	if err != nil {
		log.Printf("Error marsalling our json file: %s\n\n", err)
	}
	return string(v)
}
