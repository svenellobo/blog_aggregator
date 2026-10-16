package main

import (
	"context"
	"errors"
	"fmt"	
	"time"

	"github.com/google/uuid"
	"github.com/svenellobo/blog_aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("There are no command arguments")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("The user has been set")
	return nil
}


func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("There are no command arguments")
	}
	name := cmd.args[0]
	now := time.Now()
	params := database.CreateUserParams {
		ID: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name: name,
	}
	
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(name)
	fmt.Printf("User %s was created\n", user.Name)
	fmt.Printf("User: %+v\n", user)

	return nil

}