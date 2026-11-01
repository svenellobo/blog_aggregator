package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/svenellobo/blog_aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		n, err := strconv.Atoi(cmd.args[0])
		if err != nil || n <= 0 || n > 50 {
			return fmt.Errorf("limit must be a number between 1 and 50")
		}
		limit = n
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s [limit]", cmd.name)
	}

	userParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("couldn't retrieve posts from database :%w", err)
	}

	for _, post := range posts {
		fmt.Printf("%s\n", post.Title.String)
		if post.PublishedAt.Valid {
			fmt.Printf("  Published: %s\n", post.PublishedAt.Time.Format("2006-01-02 15:04"))
		}
		fmt.Printf("  %s\n", strings.TrimSpace(post.Description.String))
		fmt.Printf("  %s\n\n", post.Url)
	}

	return nil

}
