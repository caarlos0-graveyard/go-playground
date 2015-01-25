package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config map[string]interface{}

func (c Config) Name() string {
	return c["name"].(string)
}

func (c Config) Version() float64 {
	return c["version"].(float64)
}

func (c Config) Free() bool {
	return c["free"].(bool)
}

func LoadConfig(path string) (Config, error) {
	var m Config
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return m, err
	}

	err = json.Unmarshal(data, &m)
	return m, err
}

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	fmt.Println(config.Name())
	fmt.Println(config.Free())
	fmt.Println(config.Version())
}
