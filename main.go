package main

import (
	"flag"
	"fmt"
	"log"

	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
)

var (
	uriFlag  string
	userFlag string
	passFlag string
)

func init() {
	flag.StringVar(&uriFlag, "host", "http://localhost:15672", "rabbitmq hostname")
	flag.StringVar(&userFlag, "user", "guest", "rabbitmq user")
	flag.StringVar(&passFlag, "pass", "guest", "rabbitmq user password")
}

func main() {
	flag.Parse()

	rmqc, err := rabbithole.NewClient(uriFlag, userFlag, passFlag)
	if err != nil {
		log.Printf("[ main() -- 25 ] rmq client was created with error: %s", err.Error())
	}

	nodes, err := rmqc.ListNodes()
	if err != nil {
		log.Printf("[ main() -- 30 ] ListNodes request return error: %s", err.Error())
	}

	for _, node := range nodes {
		fmt.Println(node)
	}
}
