package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	user    = kingpin.Flag("user", "Github user").Required().OverrideDefaultFromEnvar("AUTOBREW_USER").Short('u').String()
	project = kingpin.Flag("project", "Github project").Required().OverrideDefaultFromEnvar("AUTOBREW_PROJECT").Short('p').String()
	tap     = kingpin.Flag("tap", "Homebrew tap").Required().OverrideDefaultFromEnvar("AUTOBREW_TAP").Short('b').String()
	tapUser = kingpin.Flag("tap-user", "Homebrew tap user (default to user)").OverrideDefaultFromEnvar("AUTOBREW_TAP_USER").String()
	token   = kingpin.Flag("github-token", "Github token").Required().OverrideDefaultFromEnvar("AUTOBREW_GITHUB_TOKEN").Short('t').String()
)

func main() {
	kingpin.Version(fmt.Sprintf("%s (%s)", version, commit))
	kingpin.Parse()
	if *tapUser == "" {
		tapUser = user
	}

	config := NewConfig(*token, *user, *project, *tap, *tapUser)
	discovery := NewDiscovery(config)
	publisher := NewCommitPublisher(config)
	info, _ := discovery.Discover()

	fmt.Printf("Discovered from [%s/%s]:\n", config.User, config.Project)
	fmt.Printf("%#v\n\n", *info)

	fmt.Printf("Publishing [%s] to [%s/%s]:\n", info.File, config.TapUser, config.Tap)
	println(string(info.AsFileContent()))
	publisher.Publish(info)
}
