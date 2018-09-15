package main

import "github.com/google/go-github/github"

// CommitPublisher TBD
type CommitPublisher struct {
	config *Config
}

// NewCommitPublisher TBD
func NewCommitPublisher(config *Config) *CommitPublisher {
	return &CommitPublisher{config: config}
}

// Publish TBD
func (cp *CommitPublisher) Publish(info *FormulaInfo) {
	c := cp.config
	opts := &github.RepositoryContentFileOptions{
		Message: github.String("Autobrew: update"),
		Content: info.AsFileContent(),
	}
	res, _, _, err := c.Github.Repositories.GetContents(c.Context, c.TapUser, c.Tap, info.File, nil)
	if res != nil {
		opts.SHA = res.SHA
		_, _, err = c.Github.Repositories.UpdateFile(c.Context, c.TapUser, c.Tap, info.File, opts)
		if err != nil {
			panic(err)
		}
	} else {
		_, _, err := c.Github.Repositories.CreateFile(c.Context, c.TapUser, c.Tap, info.File, opts)
		if err != nil {
			panic(err)
		}
	}
}
