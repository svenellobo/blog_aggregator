package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/svenellobo/blog_aggregator/internal/database"
)


func handlerFollow(s* state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}
	feedURL := cmd.args[0]
	now := time.Now()
	userName := s.cfg.CurrentUserName

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feed, errr := s.db.GetFeedByURL(context.Background(), feedURL)
	if errr != nil {
		return fmt.Errorf("couldn't get specific feed: %w", errr)
	}

	newFeedFollow := database.CreateFeedFollowParams {
		ID: uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("feed follow failed: %w, err")
	}
	fmt.Printf("FeedName: %s; CurrentUser: %s", feedFollow.FeedName, user.Name)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	userName := s.cfg.CurrentUserName

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't fetch any feed follows: %w", err)
	}
	for _, feed := range feedFollows {
		fmt.Printf("Feed: %s\n", feed.FeedName)
	}

	return nil
}