package command

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/bugbounce/totp-cli/security"
	s "github.com/bugbounce/totp-cli/storage"
	"github.com/yitsushi/go-commander"
)

// List structure is the representation of the list command.
type SearchList struct {
}

// Execute is the main function. It will be called on list command.
func (c *SearchList) Execute(opts *commander.CommandHelper) {
	storage := s.PrepareStorage()
	accounts := []*s.Account{}
	namespaces := []string{}

	var otps []string
	now := time.Now()
	timer := uint64(math.Floor(float64(now.Unix()) / float64(30)))
	secondsUntilInvalid := ((timer+1)*30 - uint64(now.Unix()))

	for _, ns := range storage.Namespaces {
		for _, ac := range ns.Accounts {
			accounts = append(accounts, ac)
			namespaces = append(namespaces, ns.Name)
			generatedOtp := string(security.GenerateOTPCode(ac.Token, now))
			otps = append(otps, generatedOtp)

		}
	}
	out := ""
	for i := range accounts {
		o := namespaces[i] + "." + accounts[i].Name + strings.Repeat(" ", (20-len(accounts[i].Name)-len(namespaces[i]))) + "  |  " + otps[i] + "  |  " + strconv.Itoa(int(secondsUntilInvalid))
		out = out + o + "\n"

	}

	fmt.Print(out)
}

// NewList creates a new List command.
func NewSearchList(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &SearchList{},
		Help: &commander.CommandDescriptor{
			Name:             "search",
			ShortDescription: "Search all available accounts under all namespaces",
			Arguments:        "",
			Examples: []string{
				"",
			},
		},
	}
}
