package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
	"github.com/thoas/go-funk"
)

// Discovery TBD
type Discovery struct {
	config *Config
}

// ShaUrl TBD
func shaURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		return "", err
	}
	defer resp.Body.Close()
	digest := sha256.New()
	io.Copy(digest, resp.Body)
	return fmt.Sprintf("%x", digest.Sum(nil)), nil
}
func getVersion(r *github.ReleaseAsset) string {
	versionRegex := regexp.MustCompile(".*/download/v(\\d+\\.\\d+\\..*?)/.*")
	captures := versionRegex.FindAllStringSubmatch(*r.BrowserDownloadURL, -1)
	if captures != nil {
		return captures[0][1]
	}
	return ""
}
func findOSXRelease(release *github.RepositoryRelease) (*github.ReleaseAsset, error) {
	osxRegex := regexp.MustCompile(".*x86_64-apple-darwin\\.tar\\.gz$")
	r := funk.Find(release.Assets, func(asset github.ReleaseAsset) bool {
		return osxRegex.MatchString(*asset.BrowserDownloadURL)
	})
	if r != nil {
		a := r.(github.ReleaseAsset)
		return &a, nil
	}

	return nil, errors.New("Cannot find latest release")
}

// NewDiscovery TBD
func NewDiscovery(config *Config) *Discovery {
	return &Discovery{config: config}
}

// Discover TBD
func (d *Discovery) Discover() (*FormulaInfo, error) {
	c := d.config
	release, _, _ := c.Github.Repositories.GetLatestRelease(c.Context, c.User, c.Project)
	repo, _, _ := c.Github.Repositories.Get(c.Context, c.User, c.Project)

	osxRelease, _ := findOSXRelease(release)
	sha, _ := shaURL(*osxRelease.BrowserDownloadURL)
	return &FormulaInfo{
		URL:         *osxRelease.BrowserDownloadURL,
		Version:     getVersion(osxRelease),
		Description: repo.GetDescription(),
		Digest:      sha,
		Homepage:    *repo.HTMLURL,
		Name:        strings.Title(*repo.Name),
		Bin:         *repo.Name,
		File:        fmt.Sprintf("%s.rb", *repo.Name),
	}, nil

}
