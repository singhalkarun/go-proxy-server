package main

import (
	"fmt"
	"io/ioutil"

	"singhalkarun.github.io/proxy-server/utils"
)

func main() {
	confDirectoryLocation := "conf"

	var configs []utils.Config

	files, err := ioutil.ReadDir(confDirectoryLocation)

	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			config := utils.ConfigParser(confDirectoryLocation + "/" + file.Name())

			configs = append(configs, config)
		}
	}

	for i := 0; i < len(configs); i++ {
		fmt.Println("Proxy Block: " + fmt.Sprint(i))

		fmt.Println("Location: " + configs[i].Location)
		fmt.Println("Proxy URL " + configs[i].ProxyUrl)

		fmt.Println("")
	}
}
