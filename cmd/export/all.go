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
	"centctl/cmd/export/acl"
	"centctl/cmd/export/category"
	"centctl/cmd/export/group"
	"centctl/cmd/export/template"
	"fmt"

	"github.com/spf13/cobra"
)

// allCmd represents the export/all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Export all centreon objects",
	Long:  `Export all centreon objects`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := ExportAll(file, debugV)
		if err != nil {
			fmt.Println(err)
		}

	},
}

//ExportAll permits to export all objects of the centreon server
func ExportAll(file string, debugV bool) error {
	err := ExportPoller([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = ExportEngineCFG([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = ExportResourceCFG([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = ExportTimePeriod([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = ExportCommand([]string{}, "", "all", file, true, debugV)
	if err != nil {
		return err
	}
	err = template.ExportTemplateContact([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = template.ExportTemplateHost([]string{}, "", file, true, false, debugV)
	if err != nil {
		return err
	}
	err = template.ExportTemplateService([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = ExportContact([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = group.ExportGroupContact([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = group.ExportGroupHost([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = group.ExportGroupService([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = acl.ExportACLAction([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = acl.ExportACLMenu([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = acl.ExportACLResource([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = acl.ExportACLGroup([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = ExportLDAP([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = ExportVendor([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = ExportTrap([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = ExportHost([]string{}, "", file, true, debugV, false)
	if err != nil {
		return err
	}
	err = ExportService([]string{}, file, []string{}, true, debugV)
	if err != nil {
		return err
	}
	err = category.ExportCategoryHost([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	err = category.ExportCategoryService([]string{}, "", file, true, debugV)
	if err != nil {
		return err
	}
	return nil
}

func init() {
}
