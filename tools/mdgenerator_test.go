package tools_test

import (
	"github.com/agile-rcm/support-tools/tools"
	"github.com/urfave/cli"
)

func ExampleGenerateFlagsList() {
	tools.GenerateFlagsList(cli.App{})
	// Output: MD File written successfully
}
