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

package add

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Add a command",
	Long:  `Add a command into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		typeC, _ := cmd.Flags().GetString("type")
		line, _ := cmd.Flags().GetString("line")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddCommand(name, typeC, line, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddCommand permits to add a command in the centreon server
func AddCommand(name string, typeC string, line string, debugV bool, isImport bool) error {
	//Creation of the request body
	values := name + ";" + typeC + ";" + line
	err := request.Add("add", "CMD", values, "add command", name, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}
	return nil
}

func init() {
	commandCmd.Flags().StringP("name", "n", "", "To define the name of the command")
	commandCmd.MarkFlagRequired("name")
	commandCmd.Flags().StringP("type", "t", "", "To define the type of the command (check, notif, misc or discovery)")
	commandCmd.MarkFlagRequired("type")
	commandCmd.Flags().StringP("line", "l", "", "To define the line of the command")
	commandCmd.MarkFlagRequired("line")
}
