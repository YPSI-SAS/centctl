/*
MIT License

Copyright (c) 2020 YPSI SAS
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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"

	"github.com/spf13/viper"
)

var cfgFile string
var serverName string

type serverList map[string]string

//Token represents the generate token by the authentification
type Token struct {
	Token string `json:"authToken"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "centctl",
	Short: "centctl controls Centreon throught Centreon API",
	Long:  `centctl is a tool used to controls Centreon through Centreon API (hosts, services, contacts...)`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
}

//Authentification allow the authentification at the server specified
func Authentification(urlServer string, login string, password string) (string, error) {
	urlCentreon := urlServer + "/api/index.php?action=authenticate"
	resp, err := http.PostForm(urlCentreon,
		url.Values{"username": {login}, "password": {password}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if string(body) == "\"Bad credentials\"" {
		fmt.Println("Login or password incorrect")
		os.Exit(1)
	}
	t := &Token{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return "", err
	}
	return t.Token, err
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file")
	rootCmd.PersistentFlags().StringVar(&serverName, "server", "", "server name (required)")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "helping")
	rootCmd.PersistentFlags().Bool("DEBUG", false, "debugging")
	rootCmd.MarkPersistentFlagRequired("server")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Use config file from the flag.
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is not found.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Use config file:", viper.ConfigFileUsed())
	}

	//Recover the servers from the config file
	servers := make([]serverList, 0)
	err := viper.UnmarshalKey("servers", &servers)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	//Search if the server(s) exist(s) in the list of servers
	index := -1
	for i, v := range servers {
		if serverName == v["server"] {
			index = i
		}
	}

	//If it exists => made authentification and create token and url as environments variables
	if index >= 0 {
		url := servers[index]["url"]
		login := servers[index]["login"]
		password := servers[index]["password"]
		token, err := Authentification(url, login, password)
		logger := log.New(os.Stdout).WithColor()
		if err != nil {
			logger.Error("centctl authentification - " + err.Error())
			os.Exit(1)
		}
		os.Setenv("SERVER", serverName)
		os.Setenv("TOKEN", token)
		os.Setenv("URL", url)
	} else if serverName != "" {
		fmt.Println(errors.New("The server is not correct"))
	}
}
