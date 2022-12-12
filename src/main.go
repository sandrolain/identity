package main

import (
	"log"

	"github.com/sandrolain/identity/src/config"
)

func main() {
	cfg, err := config.GetConfiguration()
	if err != nil {
		panic(err)
	}

	log.Println(cfg)
	log.Println("Hello World!")
}
