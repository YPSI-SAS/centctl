/*MIT License

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
package show

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/command"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Show one command's details",
	Long:  `Show one command's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowCommand(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowCommand permits to display the details of one command
func ShowCommand(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "CMD", name, "show command", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the commands contain into the response body
	commands := command.DetailResult{}
	json.Unmarshal(body, &commands)
	for i := range commands.Commands {
		cmd := &commands.Commands[i]
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

	//Permits to find the good command in the array
	var CommandFind command.DetailCommand
	for _, v := range commands.Commands {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			CommandFind = v
		}
	}

	var server command.DetailServer
	if CommandFind.Name != "" {
		//Organization of data
		server = command.DetailServer{
			Server: command.DetailInformations{
				Name:    os.Getenv("SERVER"),
				Command: &CommandFind,
			},
		}
	} else {
		server = command.DetailServer{
			Server: command.DetailInformations{
				Name:    os.Getenv("SERVER"),
				Command: nil,
			},
		}
	}

	//Display details of the command
	displayCommand, err := display.DetailCommand(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayCommand)
	return nil
}

func init() {
	commandCmd.Flags().StringP("name", "n", "", "To define the name of the command")
	commandCmd.MarkFlagRequired("name")
	commandCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetCommandNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
