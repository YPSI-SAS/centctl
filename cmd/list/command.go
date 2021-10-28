/*
MIT License

Copyright (c)  2020-2021 YPSI SAS
Centctl is developped by : Mélissa Bertin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package list

import (
	"centctl/colorMessage"
	"centctl/display"
	"centctl/request"
	"centctl/resources/command"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "List the commands",
	Long:  `List the commands of the Centreon Server`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		regex, _ := cmd.Flags().GetString("regex")
		typeCmd, _ := cmd.Flags().GetString("type")
		err := ListCommand(output, regex, typeCmd, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ListCommand permits to display the array of command return by the API
func ListCommand(output string, regex string, typeCmd string, debugV bool) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "CMD", "", "list command", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the commands contain into the response body
	commands := command.Result{}
	json.Unmarshal(body, &commands)
	finalCommands := commands.Commands
	if regex != "" {
		finalCommands = deleteCommand(finalCommands, regex)
	}
	if typeCmd != "all" {
		index := 0
		for _, s := range finalCommands {
			if strings.ToLower(s.Type) == strings.ToLower(typeCmd) {
				finalCommands[index] = s
				index++
			}
		}
		finalCommands = finalCommands[:index]
	}

	for i := range finalCommands {
		cmd := &finalCommands[i]
		if len(cmd.Line) > 0 {
			switch cmd.Line[0] {
			case '"':
				if err := json.Unmarshal(cmd.Line, &cmd.CmdLine); err != nil {
					return err
				}
			case '[':
				var s []string
				if err := json.Unmarshal(cmd.Line, &s); err != nil {
					return err
				}
				cmd.CmdLine = strings.Join(s, "|")
			}
		}
		cmd.Line = json.RawMessage{}
	}
	sort.SliceStable(finalCommands, func(i, j int) bool {
		return strings.ToLower(finalCommands[i].Name) < strings.ToLower(finalCommands[j].Name)
	})

	//Organization of data
	server := command.Server{
		Server: command.Informations{
			Name:     os.Getenv("SERVER"),
			Commands: finalCommands,
		},
	}

	//Display all commands
	displayCommand, err := display.Command(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayCommand)

	return nil
}

func deleteCommand(commands []command.Command, regex string) []command.Command {
	colorRed := colorMessage.GetColorRed()
	index := 0
	for _, s := range commands {
		matched, err := regexp.MatchString(regex, s.Name)
		if err != nil {
			fmt.Printf(colorRed, "ERROR:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if matched {
			commands[index] = s
			index++
		}
	}
	return commands[:index]
}

func init() {
	commandCmd.Flags().StringP("regex", "r", "", "The regex to apply on the command's name")
	commandCmd.Flags().StringP("type", "t", "all", "To define the type of command (all, notif, check, misc, discovery)")

}
