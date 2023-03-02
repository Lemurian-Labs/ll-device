package main

import (
	"github.com/lemurian-labs/lld/cmd"
)

// semver and commit are set by the build environment
// GOOS={{default OS .GOOS}} GOARCH={{default ARCH .GOARCH}} go build -ldflags "-X main.semver={{.SEMVER}} -X main.commit={{.COMMIT}}"
// with .GOOS, .GOARCH, .SEMVER, and .COMMIT being environment variables
// For example in a Taskfile.yml

var (
	semver    string
	commit    string
	buildDate string
)

func main() {
	cmd.SetVersion(semver, commit, buildDate)

	cmd.Execute()
}
