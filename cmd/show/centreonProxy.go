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
	"centctl/resources/centreonProxy"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// centreonProxyCmd represents the centreonProxy command
var centreonProxyCmd = &cobra.Command{
	Use:   "centreonProxy",
	Short: "Show centreon proxy's details",
	Long:  `Show centreon proxy's details of the Centreon server`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		output, _ := cmd.Flags().GetString("output")
		err := ShowCentreonProxy(name, debugV, output)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//ShowCentreonProxy permits to display the details of one centreonProxy
func ShowCentreonProxy(name string, debugV bool, output string) error {
	output = strings.ToLower(output)

	urlCentreon := os.Getenv("URL") + "/api/beta/configuration/proxy"
	err, body := request.GeneriqueCommandV2Get(urlCentreon, "show centreonProxy", debugV)
	if err != nil {
		return err
	}

	//Permits to recover the BVs contain into the response body
	centreonProxyRes := centreonProxy.CentreonProxy{}
	json.Unmarshal(body, &centreonProxyRes)

	var server centreonProxy.Server
	if centreonProxyRes.URL != "" {
		server = centreonProxy.Server{
			Server: centreonProxy.Informations{
				Name:          os.Getenv("SERVER"),
				CentreonProxy: &centreonProxyRes,
			},
		}
	} else {
		server = centreonProxy.Server{
			Server: centreonProxy.Informations{
				Name:          os.Getenv("SERVER"),
				CentreonProxy: nil,
			},
		}
	}

	//Display details of the centreonProxy
	displayBV, err := display.CentreonProxy(output, server)
	if err != nil {
		return err
	}
	fmt.Println(displayBV)
	return nil

}

func init() {
}
