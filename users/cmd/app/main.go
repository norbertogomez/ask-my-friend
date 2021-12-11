package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"users/internal/config"
	"users/internal/server"
)

func main() {
	environment := flag.String("e", "development", "")
	log.Println(*environment)
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	err := config.Init(*environment)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Init()

	if err != nil {
		log.Fatal(err)
	}
}
