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

package cmd

import (
	"centctl/colorMessage"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show the config file",
	Long:  `Show the config file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := readConfigFile()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func readConfigFile() error {
	colorRed := colorMessage.GetColorRed()

	cfgFile := os.Getenv("CENTCTL_CONF")
	if cfgFile == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("The environment variable CENTCTL_CONF is not set")
		os.Exit(1)
	}
	f, err := os.Open(cfgFile)
	if err != nil {
		return err
	}
	defer f.Close()
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	servers := ServerList{}
	err = yaml.Unmarshal(byteValue, &servers)
	if err != nil {
		return nil
	}

	out, _ := yaml.Marshal(servers)
	fmt.Println(string(out))
	return nil
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.ResetCommands()
}
