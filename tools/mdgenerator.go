package tools

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

// This is a generator for automatically providing a list of all commands, subcommands and flags of the provided github.com/urfave/cli CLI.
// It generates a simple markdown file in the current directory.
func GenerateFlagsList(app cli.App) error {
	file, err := os.Create("Commandlist.MD")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i := range app.Commands {
		fmt.Println(fmt.Sprintf("Command %s", app.Commands[i].Name))

		_, err := file.WriteString(fmt.Sprintf("# Command %s \n", app.Commands[i].Name))
		if err != nil {
			fmt.Println(err)
			file.Close()
			return err
		}

		for s := range app.Commands[i].Subcommands {
			fmt.Println(fmt.Sprintf("Subcommand %s", app.Commands[i].Subcommands[s].Name))

			_, err := file.WriteString(fmt.Sprintf("## Subcommand %s \n", app.Commands[i].Subcommands[s].Name))
			if err != nil {
				fmt.Println(err)
				file.Close()
				return err
			}

			fmt.Println(fmt.Sprintf("Aliases %s", app.Commands[i].Subcommands[s].Aliases))

			_, err = file.WriteString(fmt.Sprintf("### Alias: \n %s \n", app.Commands[i].Subcommands[s].Aliases))
			if err != nil {
				fmt.Println(err)
				file.Close()
				return err
			}

			_, err = file.WriteString(fmt.Sprintf("### Flag(s): \n"))
			if err != nil {
				fmt.Println(err)
				file.Close()
				return err
			}

			for f := range app.Commands[i].Subcommands[s].Flags {
				fmt.Println(fmt.Sprintf("%s", app.Commands[i].Subcommands[s].Flags[f].String()))

				_, err = file.WriteString(fmt.Sprintf("%s \n\n", app.Commands[i].Subcommands[s].Flags[f].String()))
				if err != nil {
					fmt.Println(err)
					file.Close()
					return err
				}
			}
		}
	}

	fmt.Println("MD File written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
