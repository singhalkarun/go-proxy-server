package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Route struct {
	Location string `json:"location"`
	ProxyUrl string `json:"proxy_url"`
}

type Config struct {
	Routes   []Route `json:"routes"`
	Hostname string  `json:"hostname"`
}

func ReadFile(location string) []byte {
	file, err := os.Open(location)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	return byteValue
}

func ConfigParser(location string) Config {
	byteValue := ReadFile(location)

	var config Config

	err := json.Unmarshal(byteValue, &config)

	if err != nil {
		fmt.Println(err)
	}

	return config
}
