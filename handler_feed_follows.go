package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/svenellobo/blog_aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}
	feedURL := cmd.args[0]
	now := time.Now()

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't get specific feed: %w", err)
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("feed follow failed: %w", err)
	}
	fmt.Printf("FeedName: %s; CurrentUser: %s", feedFollow.FeedName, user.Name)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	for _, feed := range feedFollows {
		fmt.Printf("Feed: %s\n", feed.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}
	feedURL := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	deleteParam := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), deleteParam)
	if err != nil {
		return fmt.Errorf("feed unfollow failed: %w", err)
	}

	return nil

}
