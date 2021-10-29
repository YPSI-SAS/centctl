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
package export

import (
	"centctl/colorMessage"
	"centctl/request"
	"centctl/resources/command"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Export a command",
	Long:  `Export in a csv file a command`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		regex, _ := cmd.Flags().GetString("regex")
		typeCmd, _ := cmd.Flags().GetString("type")
		name, _ := cmd.Flags().GetStringSlice("name")
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportCommand(name, regex, typeCmd, file, all, debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ExportCommand permits to export a command of the centreon server
func ExportCommand(name []string, regex string, typeCmd string, file string, all bool, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	if !all && len(name) == 0 && regex == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("You must pass flag name or flag all or flag regex")
		os.Exit(1)
	}

	writeFile := false
	if file != "" {
		writeFile = true
	}

	if all {
		cmds := getAllCommand(debugV)
		for _, a := range cmds {
			if strings.ToLower(typeCmd) == "all" {
				name = append(name, a.Name)
			} else {
				if strings.ToLower(a.Type) == strings.ToLower(typeCmd) {
					name = append(name, a.Name)
				}
			}

		}
	} else if regex != "" {
		cmds := getAllCommand(debugV)
		for _, a := range cmds {
			matched, err := regexp.MatchString(regex, a.Name)
			if err != nil {
				fmt.Printf(colorRed, "ERROR:")
				fmt.Println(err.Error())
				os.Exit(1)
			}
			if matched {
				name = append(name, a.Name)
			}
		}
	}
	for _, n := range name {
		err, command := getCommandInfo(n, debugV)
		if err != nil {
			return err
		}
		if command.Name == "" {
			continue
		}

		request.WriteValues("\n", file, writeFile)
		request.WriteValues("add,command,\""+command.Name+"\",\""+command.Type+"\",\""+strings.ReplaceAll(command.Line, "\"", "\"\"")+"\"\n", file, writeFile)
		request.WriteValues("modify,command,\""+command.Name+"\",graph,"+command.Graph+"\n", file, writeFile)
		request.WriteValues("modify,command,\""+command.Name+"\",example,"+command.Example+"\n", file, writeFile)
		request.WriteValues("modify,command,\""+command.Name+"\",comment,\""+command.Comment+"\"\n", file, writeFile)
		request.WriteValues("modify,command,\""+command.Name+"\",activate,\""+command.Activate+"\"\n", file, writeFile)
		request.WriteValues("modify,command,\""+command.Name+"\",enable_shell,\""+command.EnableShell+"\"\n", file, writeFile)
	}
	return nil
}

//getCommandInfo permits to get all information about a command
func getCommandInfo(name string, debugV bool) (error, command.ExportCommand) {
	colorRed := colorMessage.GetColorRed()

	//Get the parameters of the command
	values := name + ";name|line|type|graph|example|comment|activate|enable_shell"
	err, body := request.GeneriqueCommandV1Post("getparam", "CMD", values, "export command", debugV, false, "")
	if err != nil {
		return err, command.ExportCommand{}
	}
	var result command.ExportResult
	json.Unmarshal(body, &result)

	//Check if the command is found
	if len(result.Commands) == 0 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Object not found: " + name)
		return nil, command.ExportCommand{}
	}
	if result.Commands[0].Type != "2" && result.Commands[0].Type != "4" && result.Commands[0].Type != "3" && result.Commands[0].Type != "1" {
		result.Commands[0].Type = "1"
	}

	return nil, result.Commands[0]
}

//getAllBV permits to find all command in the centreon server
func getAllCommand(debugV bool) []command.ExportCommand {
	//Get all command
	err, body := request.GeneriqueCommandV1Post("show", "CMD", "", "export command", debugV, false, "")
	if err != nil {
		return []command.ExportCommand{}
	}
	var resultCmd command.ExportResult
	json.Unmarshal(body, &resultCmd)

	return resultCmd.Commands
}

func init() {
	commandCmd.Flags().StringSliceP("name", "n", []string{}, "Command's name (separate by a comma the multiple values)")
	commandCmd.Flags().StringP("type", "t", "all", "To define the type of command (all, notif, check, misc, discovery)")
	commandCmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"check", "notif", "misc", "discovery", "all"}, cobra.ShellCompDirectiveDefault
	})
	commandCmd.Flags().StringP("regex", "r", "", "The regex to apply on the command's name")

}
