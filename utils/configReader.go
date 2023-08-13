package utils

import (
	"fmt"
	"io/ioutil"
)

func ReadConfig() []Config {
	confDirectoryLocation := "conf"

	var configs []Config

	files, err := ioutil.ReadDir(confDirectoryLocation)

	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			config := ConfigParser(confDirectoryLocation + "/" + file.Name())

			configs = append(configs, config)
		}
	}

	return configs
}
