package command

import (
	"strings"
	"github.com/mitchellh/cli"
)

func mockSSHCommand() (*cli.MockUi, *SSHCommand) {
	ui := cli.NewMockUi()
	return ui, &SSHCommand{
		BaseCommand: &BaseCommand{
			UI: ui,
		},
	}
}

func FuzzerEntrypoint(Data []byte) int {
	_, cmd := mockSSHCommand()

	s := string( Data )
	sList := strings.Split( s, "\n" )
	// hostname, username, port, err := cmd.parseSSHCommand( sList )
	cmd.parseSSHCommand( sList )
	return 0;
}
