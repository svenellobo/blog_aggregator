package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/svenellobo/blog_aggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %+v", feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	name := cmd.args[0]
	url := cmd.args[1]
	now := time.Now()
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil

}

func handlerListFeeds (s *state, cmd command) error {
	feeds, err := s.db.GetFeed(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch list of feeds: %w", err)
	}
	for _, feed := range feeds {
		fmt.Printf("Feed name: %s; Feed URL: %s; Created by: %s\n", feed.FeedName, feed.Url, feed.UserName)
	}
	return nil
}