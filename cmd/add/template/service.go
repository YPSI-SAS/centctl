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
package template

import (
	"centctl/request"
	"fmt"

	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Add template service",
	Long:  `Add template service of the centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		alias, _ := cmd.Flags().GetString("alias")
		template, _ := cmd.Flags().GetString("template")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddTemplateService(name, alias, template, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddTemplateService permits to add a service template in the centreon server
func AddTemplateService(name string, alias string, template string, debugV bool, isImport bool) error {
	//Creation of the request body
	var values string
	if template == "" {
		values = name + ";" + alias + ";"
	} else {
		values = name + ";" + alias + ";" + template
	}
	err := request.Add("add", "STPL", values, "add template service", name, debugV, isImport, false, "", "")
	if err != nil {
		return err
	}

	return nil
}

func init() {
	serviceCmd.Flags().StringP("name", "n", "", "The name of the service template")
	serviceCmd.MarkFlagRequired("name")
	serviceCmd.Flags().StringP("alias", "a", "", "The alias of the service template")
	serviceCmd.MarkFlagRequired("alias")
	serviceCmd.Flags().StringP("template", "t", "", "To define the template to wich the service template is attached")
	serviceCmd.RegisterFlagCompletionFunc("template", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var values []string
		if request.InitAuthentification(cmd) {
			values = request.GetTemplateServiceNames()
		}
		return values, cobra.ShellCompDirectiveDefault
	})
}
