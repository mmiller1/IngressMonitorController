package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Providers             []Provider `yaml:"providers"`
	EnableMonitorDeletion bool       `yaml:"enableMonitorDeletion"`
	MonitorNameTemplate   string     `yaml:"monitorNameTemplate"`
}

type Provider struct {
	Name          string `yaml:"name"`
	ApiKey        string `yaml:"apiKey"`
	ApiURL        string `yaml:"apiURL"`
	AlertContacts string `yaml:"alertContacts"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
}

func (p *Provider) createMonitorService() MonitorServiceProxy {
	monitorService := (&MonitorServiceProxy{}).OfType(p.Name)
	monitorService.Setup(p.ApiKey, p.ApiURL, p.AlertContacts, p.Username, p.Password)
	return monitorService
}

func ReadConfig(filePath string) Config {
	var config Config
	// Read YML
	log.Println("reading yaml configuration")
	source, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}

	// Unmarshall
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Panic(err)
	}

	return config
}
