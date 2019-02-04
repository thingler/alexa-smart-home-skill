package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config
type Config struct {
	Region      string `json:"region"`
	IOTEndpoint string `json:"iot_endpoint"`
	IOTTopic    string `json:"iot_topic"`
}

func (c *Config) Parse() error {
	cfgStr, err := ioutil.ReadFile("config.json")
	if err == nil {
		json.Unmarshal(cfgStr, c)
	}
	return err
}
