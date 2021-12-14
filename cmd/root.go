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
	"centctl/cmd/acknowledge"
	"centctl/cmd/add"
	"centctl/cmd/apply"
	"centctl/cmd/delete"
	"centctl/cmd/downtime"
	"centctl/cmd/export"
	"centctl/cmd/list"
	"centctl/cmd/modify"
	"centctl/cmd/show"
	"centctl/cmd/submit"
	"centctl/request"
	"io/ioutil"
	"strings"

	"centctl/colorMessage"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/withmandala/go-log"
	"gopkg.in/yaml.v2"

	"github.com/spf13/viper"
)

var serverName string
var insecure bool
var colorRed = colorMessage.GetColorRed()
var serverLogin string
var serverPassword string
var serverUrl string
var serverVersion string
var serverHttpProxyURL string
var serverHttpsProxyURL string
var serverUserProxy string
var serverPasswordProxy string

type ServerList struct {
	Servers []struct {
		Server   string `yaml:"server"`
		Url      string `yaml:"url"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Version  string `yaml:"version"`
		Default  bool   `yaml:"default,omitempty"`
		Insecure bool   `yaml:"insecure,omitempty"`
		Proxy    []struct {
			HttpURL  string `yaml:"httpURL,omitempty"`
			HttpsURL string `yaml:"httpsURL,omitempty"`
			User     string `yaml:"user,omitempty"`
			Password string `yaml:"password,omitempty"`
		} `yaml:"proxy,omitempty"`
	} `yaml:"servers"`
}

type ProxyStruct struct {
	Proxy []struct {
		HttpURL  string `yaml:"httpURL,omitempty"`
		HttpsURL string `yaml:"httpsURL,omitempty"`
		User     string `yaml:"user,omitempty"`
		Password string `yaml:"password,omitempty"`
	} `yaml:"proxy,omitempty"`
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
	rootCmd.PersistentFlags().StringVar(&serverName, "server", "", "server name")
	rootCmd.PersistentFlags().BoolVar(&insecure, "insecure", false, "To forced connection https")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "helping")
	rootCmd.PersistentFlags().Bool("DEBUG", false, "debugging")
	rootCmd.PersistentFlags().StringVar(&serverLogin, "login", "", "Server Login")
	rootCmd.PersistentFlags().StringVar(&serverPassword, "password", "", "Server Password")
	rootCmd.PersistentFlags().StringVar(&serverUrl, "url", "", "Server URL")
	rootCmd.PersistentFlags().StringVar(&serverVersion, "version", "", "Server version")
	rootCmd.PersistentFlags().StringVar(&serverHttpProxyURL, "proxyHTTP", "", "URL proxy HTTP")
	rootCmd.PersistentFlags().StringVar(&serverHttpsProxyURL, "proxyHTTPs", "", "URL proxy HTTPs")
	rootCmd.PersistentFlags().StringVar(&serverPasswordProxy, "proxyPassword", "", "Proxy password")
	rootCmd.PersistentFlags().StringVar(&serverUserProxy, "proxyUser", "", "Proxy user")
	rootCmd.RegisterFlagCompletionFunc("server", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var serversList []string
		if os.Getenv("CENTCTL_CONF") != "" {
			servers := &ServerList{}
			yamlFile, _ := ioutil.ReadFile(os.Getenv("CENTCTL_CONF"))
			_ = yaml.Unmarshal(yamlFile, servers)
			for _, server := range servers.Servers {
				serversList = append(serversList, server.Server)
			}
		}
		return serversList, cobra.ShellCompDirectiveDefault
	})

	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(export.Cmd)
	rootCmd.AddCommand(downtime.Cmd)
	rootCmd.AddCommand(list.Cmd)
	rootCmd.AddCommand(acknowledge.Cmd)
	rootCmd.AddCommand(delete.Cmd)
	rootCmd.AddCommand(show.Cmd)
	rootCmd.AddCommand(apply.Cmd)
	rootCmd.AddCommand(add.Cmd)
	rootCmd.AddCommand(modify.Cmd)
	rootCmd.AddCommand(submit.Cmd)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if os.Args[1] != "version" && os.Args[1] != "completion" && os.Args[1] != "encrypt" {
		// colorRed := colorMessage.GetColorRed()
		cfgFile := os.Getenv("CENTCTL_CONF")
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)

		viper.AutomaticEnv() // read in environment variables that match

		var name string
		var login string
		var password string
		var url string
		var version string
		// If a config file is not found.
		if err := viper.ReadInConfig(); err != nil {
			if serverLogin != "" && serverPassword != "" && serverUrl != "" && serverVersion != "" && serverName != "" {
				name = serverName
				login = serverLogin
				password = serverPassword
				version = serverVersion
				url = serverUrl
				if serverHttpProxyURL != "" && serverPasswordProxy != "" && serverUserProxy != "" {
					os.Setenv("http_proxy", "http://"+serverUserProxy+":"+serverPasswordProxy+"@"+serverHttpProxyURL)
				} else if serverHttpProxyURL != "" {
					os.Setenv("http_proxy", "http://"+serverHttpProxyURL)
				}
				if serverHttpsProxyURL != "" && serverPasswordProxy != "" && serverUserProxy != "" {
					os.Setenv("https_proxy", "http://"+serverUserProxy+":"+serverPasswordProxy+"@"+serverHttpsProxyURL)
				} else if serverHttpsProxyURL != "" {
					os.Setenv("https_proxy", "http://"+serverHttpsProxyURL)
				}
			} else {
				fmt.Printf(colorRed, "ERROR: ")
				fmt.Println("Fill in the server either in the yaml configuration file or with the following flags: --server --password --login --url --version")
				os.Exit(1)
			}
		} else {
			name, login, password, url, version = getValueInFile()
		}
		if password == "" && url != "" && name != "" {
			password = getPasswordStdin(name)
		}
		var token string
		var err error
		if version == "v1" {
			token, err = request.AuthentificationV1(url, login, password, insecure)
		} else if version == "v2" {
			token, err = request.AuthentificationV2(url, login, password, insecure)
		}
		logger := log.New(os.Stdout).WithColor()
		if err != nil {
			logger.Error("centctl authentification - " + err.Error())
			os.Exit(1)
		}
		os.Setenv("SERVER", name)
		os.Setenv("LOGIN", login)
		os.Setenv("PASSWORD", password)
		os.Setenv("TOKEN", token)
		os.Setenv("URL", url)
		os.Setenv("VERSION", version)
	}

}

func getPasswordStdin(name string) string {
	fmt.Print("Enter the password for the server \"" + name + "\": ")
	var input string
	fmt.Scanln(&input)
	return input
}

func getValueInFile() (string, string, string, string, string) {
	//Recover the servers from the config file
	servers := &ServerList{}
	proxy := &ProxyStruct{}
	err := viper.Unmarshal(servers)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	_ = viper.Unmarshal(proxy)

	//Search if the server(s) exist(s) in the list of servers
	index := -1
	for i, v := range servers.Servers {
		if serverName == v.Server {
			index = i
		}
	}
	if index == -1 && serverName == "" {
		for i, v := range servers.Servers {
			if true == v.Default {
				index = i
			}
		}
		if index == -1 {
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println(errors.New("No default server"))
			os.Exit(1)
		}
	}

	//If it exists => made authentification and create token and url as environments variables
	if index >= 0 {
		name := servers.Servers[index].Server
		url := servers.Servers[index].Url
		login := servers.Servers[index].Login
		var password string
		if strings.HasPrefix(servers.Servers[index].Password, "CENTCRYPT_") {
			ciphertext := strings.TrimPrefix(servers.Servers[index].Password, "CENTCRYPT_")
			keyString := os.Getenv("CENTCTL_DECRYPT_KEY")
			if keyString == "" {
				fmt.Printf(colorRed, "ERROR: ")
				fmt.Println(errors.New("The env variable CENTCTL_DECRYPT_KEY is not set but your passwords are crypted"))
				os.Exit(1)
			}
			password = request.Decrypt(ciphertext, keyString)
		} else {
			password = servers.Servers[index].Password
		}

		version := servers.Servers[index].Version
		if os.Getenv("http_proxy") == "" {
			if len(proxy.Proxy) != 0 && proxy.Proxy[0].HttpURL != "" {
				if proxy.Proxy[3].Password != "" {
					os.Setenv("http_proxy", "http://"+proxy.Proxy[2].User+":"+proxy.Proxy[3].Password+"@"+proxy.Proxy[0].HttpURL)
				} else {
					os.Setenv("http_proxy", "http://"+proxy.Proxy[0].HttpURL)
				}
			} else if len(servers.Servers[index].Proxy) != 0 && servers.Servers[index].Proxy[0].HttpURL != "" {
				if servers.Servers[index].Proxy[3].Password != "" {
					os.Setenv("http_proxy", "http://"+servers.Servers[index].Proxy[2].User+":"+servers.Servers[index].Proxy[3].Password+"@"+servers.Servers[index].Proxy[0].HttpURL)
				} else {
					os.Setenv("http_proxy", "http://"+servers.Servers[index].Proxy[0].HttpURL)
				}
			}
		}

		if os.Getenv("https_proxy") == "" {
			if len(proxy.Proxy) != 0 && proxy.Proxy[1].HttpsURL != "" {
				if proxy.Proxy[3].Password != "" {
					os.Setenv("https_proxy", "http://"+proxy.Proxy[2].User+":"+proxy.Proxy[3].Password+"@"+proxy.Proxy[1].HttpsURL)
				} else {
					os.Setenv("https_proxy", "http://"+proxy.Proxy[1].HttpsURL)
				}
			} else if len(servers.Servers[index].Proxy) != 0 && servers.Servers[index].Proxy[1].HttpsURL != "" {
				if servers.Servers[index].Proxy[3].Password != "" {
					os.Setenv("https_proxy", "http://"+servers.Servers[index].Proxy[2].User+":"+servers.Servers[index].Proxy[3].Password+"@"+servers.Servers[index].Proxy[1].HttpsURL)
				} else {
					os.Setenv("https_proxy", "http://"+servers.Servers[index].Proxy[1].HttpsURL)
				}
			}
		}
		if servers.Servers[index].Insecure == true {
			insecure = true
		} else if servers.Servers[index].Insecure == false {
			insecure = false
		}

		return name, login, password, url, version
	} else if serverName != "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(errors.New("The server is not correct"))
		os.Exit(1)
		return "", "", "", "", ""
	}
	return "", "", "", "", ""
}
