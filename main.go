package main

import (
	"fmt"
	"log"
	"os"

	"github.com/svenellobo/blog_aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	currentState := &state{
		cfg: &cfg,
	}

	commandsCollection := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	
	commandsCollection.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	commandInstance := command{
		name: commandName,
		args: commandArgs,
	}

	err = commandsCollection.run(currentState, commandInstance)
	if err != nil {
		log.Fatal("username is required")
	}
	
	fmt.Println(cfg)

}
