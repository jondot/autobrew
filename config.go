package main

import (
	"context"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Config struct {
	Context context.Context
	Github  *github.Client
	Tap     string
	TapUser string
	User    string
	Project string
}

func NewConfig(token string, user string, project string, tap string, tapUser string) *Config {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv(token)},
	)
	tc := oauth2.NewClient(ctx, ts)
	c := github.NewClient(tc)
	return &Config{
		Context: ctx,
		Github:  c,
		Tap:     tap,
		TapUser: tapUser,
		User:    user,
		Project: project,
	}
}
