# Blog Aggregator

blog_aggregator is a small command line RSS feed aggregator written in Go. You add and follow the feeds you like, collect posts in the background, and  browse the posts from your terminal.

I built it as part of the Boot.dev Backend Developer path.

## What you need

- **Go** (a recent version). You only need it for the install. Once it's installed, the `blog_aggregator` binary runs on its own.
- **PostgreSQL**. The program stores users, feeds and posts in a Postgres database, so you need a running Postgres server.

## Installing

Install the CLI with `go install`:

```bash
go install github.com/svenellobo/blog_aggregator@latest
```


## Setup

### 1. Create the database


```bash
psql -c "CREATE DATABASE blog_aggregator;"
```

### 2. Run the migrations

The database schema lives in `sql/schema`. I used [goose](https://github.com/pressly/goose) for migrations:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
cd sql/schema
goose postgres "postgres://username:password@localhost:5432/blog_aggregator" up
```

Use your own Postgres username and password.

### 3. Create the config file

The program reads its settings from a file called `.gatorconfig.json` in your home directory. Create it with this content:

```json
{
"db_url": "postgres://username:password@localhost:5432/blog_aggregator?sslmode=disable"
}
```
User is filled automatically when you register or log in.


## Using it

Start by registering a user:

```bash
blog_aggregator register name
```

Add a feed:

```bash
blog_aggregator addfeed "TechCrunch" "https://techcrunch.com/feed/"
```

Start collecting posts. This runs in an interval of your choice so let it run in its own terminal windows:

```bash
blog_aggregator agg 1m
```

Browse the posts: 

```bash
blog_aggregator browse 10
```

### All commands

| Command | What it does |
| --- | --- |
| `register <name>` | Create a new user and log in as them |
| `login <name>` | Switch to an existing user |
| `users` | List all users and show who's logged in |
| `addfeed <name> <url>` | Add a new feed and follow it |
| `feeds` | List every feed in the database |
| `follow <url>` | Follow a feed someone already added |
| `following` | Show the feeds you follow |
| `unfollow <url>` | Stop following a feed |
| `agg <interval>` | Fetch posts on a loop, e.g. `30s`, `1m`, `1h` |
| `browse [limit]` | Show the latest posts from your feeds |
| `reset` | Delete all users (and everything tied to them). Careful with this one! |

