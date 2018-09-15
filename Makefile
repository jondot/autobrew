setup:
	brew install goreleaser/tap/goreleaser
deps:
	go mod tidy && go mod vendor
release:
	goreleaser --rm-dist

.PHONY: deps setup release
