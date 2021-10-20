/*MIT License

Copyright (c)  2020-2021 YPSI SAS


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
package modify

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// timePeriodCmd represents the timePeriod command
var timePeriodCmd = &cobra.Command{
	Use:   "timePeriod",
	Short: "Modfiy a timePeriod",
	Long:  `Modfiy a timePeriod into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyTimePeriod(name, parameter, value, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyTimePeriod permits to modify a timePeriod in the centreon server
func ModifyTimePeriod(name string, parameter string, value string, debugV bool, isImport bool, detail bool) error {
	var values string
	var action string

	switch parameter {
	case "exception":
		action = "setexception"
		values = name + ";" + value
	default:
		action = "setparam"
		values = name + ";" + parameter + ";" + value
	}

	err := request.Modify(action, "TP", values, "modify timePeriod", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	timePeriodCmd.Flags().StringP("name", "n", "", "To define the name of the timePeriod to be modified")
	timePeriodCmd.MarkFlagRequired("name")
	timePeriodCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: exception")
	timePeriodCmd.MarkFlagRequired("parameter")
	timePeriodCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. If parameter is exception the value must be of the form : day;timeRange")
	timePeriodCmd.MarkFlagRequired("value")
}
