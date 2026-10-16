package main

import (
	"github.com/svenellobo/blog_aggregator/internal/config"
	"github.com/svenellobo/blog_aggregator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}
