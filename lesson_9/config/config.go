package config

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

type Config struct{
	LineNumber int64 `json:"linenumber"`
	FileName string `json:"filename"`
	URL string `json:"url"`
}


func Configure(configFileName string) (Config, error) {
	f, err := ioutil.ReadFile(configFileName)
	if err != nil{
		log.Println(err)
		return Config{}, err
	}
	c := Config{}
	err = json.Unmarshal(f, &c)

	if err != nil{
		log.Println(err)
		return Config{}, err
	}

	return c, nil
}