package tools_test

import (
	"git.agiletech.de/AgileRCM/support-tools/tools"
	"github.com/urfave/cli"
)

func ExampleGenerateFlagsList() {
	tools.GenerateFlagsList(cli.App{})
	// Output: MD File written successfully
}
