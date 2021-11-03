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
package show

import (
	"centctl/display"
	"centctl/request"
	"centctl/resources/trap"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// trapCmd represents the trap command
var trapCmd = &cobra.Command{
	Use:   "trap",
	Short: "Show one trap's details",
	Long:  `Show one trap's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowTrap(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowTrap permits to display the details of one cooleanrule
func ShowTrap(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	err, body := request.GeneriqueCommandV1Post("show", "trap", name, "show trap", debugV, false, "")
	if err != nil {
		return err
	}

	//Permits to recover the booleanrules contain into the response body
	traps := trap.DetailResult{}
	json.Unmarshal(body, &traps)

	//Permits to find the good trap in the array
	var TrapFing trap.DetailTrap
	for _, v := range traps.Traps {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			TrapFing = v
		}
	}

	var server trap.DetailServer
	if TrapFing.Name != "" {
		//Organization of data
		server = trap.DetailServer{
			Server: trap.DetailInformations{
				Name: os.Getenv("SERVER"),
				Trap: &TrapFing,
			},
		}
	} else {
		server = trap.DetailServer{
			Server: trap.DetailInformations{
				Name: os.Getenv("SERVER"),
				Trap: nil,
			},
		}
	}

	//Display details of the trap
	displayTrap, err := display.DetailTrap(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayTrap)
	return nil
}

func init() {
	trapCmd.Flags().StringP("name", "n", "", "To define the name of the trap")
	trapCmd.MarkFlagRequired("name")
	trapCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetTrapNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
