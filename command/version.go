package command

import (
	"fmt"
	"runtime"

	"github.com/bugbounce/totp-cli/info"
	"github.com/yitsushi/go-commander"
)

// Version structure is the representation of the Version command.
type Version struct {
}

// Execute is the main function. It will be called on version command.
func (c *Version) Execute(opts *commander.CommandHelper) {
	fmt.Printf("%s %s (%s/%s)\n", info.AppName, info.AppVersion, runtime.GOOS, runtime.GOARCH)
}

// NewVersion creates a new Version command.
func NewVersion(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Version{},
		Help: &commander.CommandDescriptor{
			Name:             "version",
			ShortDescription: "Print current version of this application",
		},
	}
}
