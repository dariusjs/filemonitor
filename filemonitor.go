package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)

type Directory struct {
	Directory string `json:"directory"`
}

type Config struct {
	Directories []Directory `json:"directories"`
}
	
func LoadConfiguration(filename string) (Config, error) {
    var config Config
    configFile, err := os.Open(filename)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config, err
}

func CountFiles() {
	 
}

func main() {
	config, _ := LoadConfiguration("config.json")
	fmt.Println(config)
	fmt.Println(config.Directories)

	for _, dir := range config.Directories {
		fmt.Println(dir.Directory)

		files, err := ioutil.ReadDir(dir.Directory)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}

	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }
}
