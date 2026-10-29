package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/svenellobo/blog_aggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs> (e.g. 1m)", cmd.name)
	}
	
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't parse duration: %w", err)
	}

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
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

	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return fmt.Errorf("feed follow failed: %w", err)
	}

	return nil

}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch list of feeds: %w", err)
	}
	for _, feed := range feeds {
		fmt.Printf("Feed name: %s; Feed URL: %s; Created by: %s\n", feed.FeedName, feed.Url, feed.UserName)
	}
	return nil
}

func scrapeFeeds(s *state) {
	localFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("couldn't get the next feed: %v", err)
		return
	}

	err = s.db.MarkFeedFetched(context.Background(), localFeed.ID)
	if err != nil {
		log.Printf("couldn't mark feed as fetched: %v", err)
		return
	}
	feed, err := fetchFeed(context.Background(), localFeed.Url)
	if err != nil {
		log.Printf("couldn't fetch the feed: %v", err)
		return
	}

	for _, item := range feed.Channel.Item {
		fmt.Printf("Feed: %s --- Title: %s\n", feed.Channel.Title, item.Title)
	}
}
