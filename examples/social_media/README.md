# Social Media

A Twitter/X-inspired social media app built with Gox templates, Go standard library HTTP server, and SQLite.

## Features

- User registration and login (bcrypt + cookie sessions)
- Post create, read, delete
- Home feed (posts from followed users + own posts)
- Explore page (all posts + suggested users)
- Follow / unfollow users
- Like / unlike posts
- Comments on posts
- Threaded replies (post-as-reply creates conversation threads)
- User profiles with follower/following counts
- Twitter/X dark theme UI

## Quick Start

```bash
make start
```

Then open http://localhost:5000, register an account, and start posting.

## Development

```bash
make dev    # watch .gox changes + run server
```

## Build

```bash
make generate   # compile .gox → .go
make build      # generate + go build
make clean      # remove generated files + binary
```
