package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Batch 		int 	`yaml:"batch"`
	FilePath 	string 	`yaml:"filePath"`
	Db 		struct {
		Host 		string `yaml:"host"`
		Port 		int		`yaml:"port"`
		User 		string  `yaml:"user"`
		Password 	string  `yaml:"password"`
		Dbname 		string  `yaml:"dbname"`
	}
}
func (c *Conf) GetConf() *Conf {

	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}