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
package add

import (
	"centctl/colorMessage"
	"centctl/request"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// centreonProxyCmd represents the centreonProxy command
var centreonProxyCmd = &cobra.Command{
	Use:   "centreonProxy",
	Short: "Add a Centreon proxy",
	Long:  `Add a Centreon proxy into the Centreon server specifified by the flag --server`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")
		port, _ := cmd.Flags().GetInt("port")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		err := AddCentreonProxy(url, login, password, port, debugV, false)
		if err != nil {
			fmt.Println(err)
		}
	},
}

//AddCentreonProxy permits to add a centreon Proxy in the centreon server
func AddCentreonProxy(url string, login string, password string, port int, debugV bool, isImport bool) error {
	colorGreen := colorMessage.GetColorGreen()
	//Creation of the request body
	requestBody, err := json.Marshal(map[string]string{
		"url":      url,
		"port":     strconv.Itoa(port),
		"user":     login,
		"password": password,
	})
	if err != nil {
		return err
	}
	urlCentreon := "/configuration/proxy"
	err = request.GeneriqueCommandV2Put(urlCentreon, requestBody, "add centreonProxy", debugV)
	if err != nil {
		return err
	}
	fmt.Printf(colorGreen, "INFO: ")
	fmt.Printf("The centreonProxy %s is updated\n", url)
	return nil
}

func init() {
	centreonProxyCmd.Flags().StringP("url", "u", "", "To define the url of the configuration of the Centreon proxy")
	centreonProxyCmd.MarkFlagRequired("url")
	centreonProxyCmd.Flags().Int("port", -1, "To define the port of the proxy")
	centreonProxyCmd.MarkFlagRequired("port")
	centreonProxyCmd.Flags().StringP("login", "l", "", "To define the login used to connect to proxy")
	centreonProxyCmd.MarkFlagRequired("login")
	centreonProxyCmd.Flags().StringP("password", "p", "", "To define the password used to connect to proxy")
	centreonProxyCmd.MarkFlagRequired("password")
}
