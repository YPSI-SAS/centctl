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
package template

import (
	"centctl/colorMessage"
	"centctl/request"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Modify template host",
	Long:  `Modify template host of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		parameter, _ := cmd.Flags().GetString("parameter")
		value, _ := cmd.Flags().GetString("value")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ModifyTemplateHost(name, parameter, value, debugV, false, true)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ModifyTemplateHost permits to modify a host in the centreon server
func ModifyTemplateHost(name string, parameter string, value string, debugV bool, isImport bool, detail bool) error {
	colorRed := colorMessage.GetColorRed()
	var action string
	var values string
	object := "HTPL"

	switch strings.ToLower(parameter) {
	case "template":
		action = "addtemplate"
		values = name + ";" + value
	case "linkedservice":
		action = "addhosttemplate"
		values = value + ";" + name
		object = "STPL"
	case "hostcategorie":
		action = "addmember"
		values = value + ";" + name
		object = "HC"
	case "contactgroup":
		action = "addcontactgroup"
		values = name + ";" + value
	case "contact":
		action = "addcontact"
		values = name + ";" + value
	case "macro":
		valueSplit := strings.Split(value, "|")
		if len(valueSplit) != 4 {
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println("The new value for macro must be of the form : macroName|macroValue|IsPassword(0 or 1)|macroDescription")
			os.Exit(1)
		}
		action = "setmacro"
		values = name + ";" + valueSplit[0] + ";" + valueSplit[1] + ";" + valueSplit[2] + ";" + valueSplit[3]
	default:
		action = "setparam"
		values = name + ";" + parameter + ";" + value

	}

	err := request.Modify(action, object, values, "modify template host", name, parameter, detail, debugV, false, "", isImport)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	hostCmd.Flags().StringP("name", "n", "", "To define the name of the host template to be modified")
	hostCmd.MarkFlagRequired("name")
	hostCmd.Flags().StringP("parameter", "p", "", "To define the parameter set in setparam section of centreon documentation or in this list: template,contactgroup,contact,linkedservice,hostcategorie,macro")
	hostCmd.MarkFlagRequired("parameter")
	hostCmd.Flags().StringP("value", "v", "", "To define the new value of the parameter to be modified. If parameter is MACRO the value must be of the form : macroName|macroValue|IsPassword(0 or 1)|macroDescription")
	hostCmd.MarkFlagRequired("value")
}
