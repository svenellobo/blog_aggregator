package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/svenellobo/blog_aggregator/internal/config"
	"github.com/svenellobo/blog_aggregator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)

	currentState := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	

	commandsCollection := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	
	commandsCollection.register("login", handlerLogin)
	commandsCollection.register("register", handlerRegister)
	commandsCollection.register("reset", handlerReset)

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
		log.Fatal(err)
	}
	
	fmt.Println(cfg)

}
