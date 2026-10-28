package main

import (
	"context"
	"fmt"
	
	"github.com/svenellobo/blog_aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	
	return func(s *state, cmd command) error{
	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	err = handler(s, cmd, user)
	if err != nil {
		return err
	}

	return nil
	
	}
}