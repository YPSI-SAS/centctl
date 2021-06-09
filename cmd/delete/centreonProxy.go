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

package delete

import (
	"centctl/colorMessage"
	"centctl/request"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// centreonProxyCmd represents the centreonProxy command
var centreonProxyCmd = &cobra.Command{
	Use:   "centreonProxy",
	Short: "Delete a Centreon proxy",
	Long:  `Delete a Centreon proxy into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := DeleteCentreonProxy(debugV)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//DeleteCentreonProxy permits to delete a centreon Proxy in the centreon server
func DeleteCentreonProxy(debugV bool) error {
	colorGreen := colorMessage.GetColorGreen()
	//Creation of the request body
	requestBody, err := json.Marshal(map[string]string{
		"url":      "",
		"port":     "",
		"user":     "",
		"password": "",
	})
	if err != nil {
		return err
	}
	urlCentreon := os.Getenv("URL") + "/api/beta/configuration/proxy"
	err = request.GeneriqueCommandV2Put(urlCentreon, requestBody, "add centreonProxy", debugV)
	if err != nil {
		return err
	}
	fmt.Printf(colorGreen, "INFO: ")
	fmt.Printf("The centreonProxy is deleted\n")
	return nil
}

func init() {
}
