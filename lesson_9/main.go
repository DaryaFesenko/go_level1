package main

import (
	"fmt"
	"./config"
	"flag"
)

func main(){
	configFileName := flag.String("c", "", "config file in json")
	flag.Parse()

	c, err := config.Configure(*configFileName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c)
}
