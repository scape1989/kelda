// +build !windows

package command

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/kelda/kelda/db"
	"github.com/kelda/kelda/minion"
	"github.com/kelda/kelda/util"
	"github.com/kelda/kelda/version"

	log "github.com/sirupsen/logrus"
)

// Minion contains the options for running the Kelda minion.
type Minion struct {
	role                            string
	inboundPubIntf, outboundPubIntf string

	connectionFlags
}

// NewMinionCommand creates a new Minion command instance.
func NewMinionCommand() *Minion {
	return &Minion{}
}

var minionCommands = "kelda minion [OPTIONS] ROLE"
var minionExplanation = `Run the kelda minion.

ROLE should be Worker or Master.`

// InstallFlags sets up parsing for command line flags.
func (mCmd *Minion) InstallFlags(flags *flag.FlagSet) {
	mCmd.connectionFlags.InstallFlags(flags)
	flags.StringVar(&mCmd.role, "role", "", "the role of this kelda minion (Worker"+
		" or Master)")
	flags.StringVar(&mCmd.inboundPubIntf, "inbound-pub-intf", "",
		"the interface on which to allow inbound traffic")
	flags.StringVar(&mCmd.outboundPubIntf, "outbound-pub-intf", "",
		"the interface on which to allow outbound traffic")

	flags.Usage = func() {
		util.PrintUsageString(minionCommands, minionExplanation, flags)
	}
}

// Parse parses the command line arguments for the minion command.
func (mCmd *Minion) Parse(args []string) error {
	if len(args) > 0 {
		mCmd.role = args[0]
	}
	return nil
}

// BeforeRun makes any necessary post-parsing transformations.
func (mCmd *Minion) BeforeRun() error {
	return nil
}

// AfterRun performs any necessary post-run cleanup.
func (mCmd *Minion) AfterRun() error {
	return nil
}

// Run starts the minion.
func (mCmd *Minion) Run() int {
	log.WithField("version", version.Version).Info("Starting Kelda minion")

	if err := mCmd.run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 1
	}
	return 0
}

func (mCmd *Minion) run() error {
	role, err := db.ParseRole(mCmd.role)
	if err != nil || role == db.None {
		return errors.New("no or improper role specified")
	}

	minion.Run(role, mCmd.inboundPubIntf, mCmd.outboundPubIntf)
	return nil
}
