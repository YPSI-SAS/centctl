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
	"bytes"
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
	"crypto/tls"

	"centctl/colorMessage"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
		Default  string `yaml:"default,omitempty"`
		Insecure string `yaml:"insecure,omitempty"`
		Proxy    []struct {
			HttpURL  string `yaml:"httpURL"`
			HttpsURL string `yaml:"httpsURL"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"proxy,omitempty"`
	} `yaml:"servers"`
}

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

//AuthentificationV1 allow the authentification at the server specified with API v1
func AuthentificationV1(urlServer string, login string, password string, insecure bool) (string, error) {
	urlCentreon := urlServer + "/api/index.php?action=authenticate"
	if insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	resp, err := http.PostForm(urlCentreon,
		url.Values{"username": {login}, "password": {password}})
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 401 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(resp.Status)
		os.Exit(1)
	} else if resp.StatusCode == 407 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(resp.Status)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if string(body) == "\"Bad credentials\"" {
		fmt.Printf(colorRed, "ERROR: ")
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

//AuthentificationV2 allow the authentification at the server specified with API v2
func AuthentificationV2(urlServer string, login string, password string, insecure bool) (string, error) {
	request := make(map[string]interface{})
	request["security"] = map[string]interface{}{
		"credentials": map[string]interface{}{
			"login":    login,
			"password": password,
		},
	}

	// Marshal the map into a JSON string.
	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	urlCentreon := urlServer + "/api/beta/login"

	if insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	resp, err := http.Post(urlCentreon, "application/json",
		bytes.NewBuffer(requestBody))

	if err != nil {
		return "", err
	}

	if resp.StatusCode == 401 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(resp.Status)
		os.Exit(1)
	} else if resp.StatusCode == 407 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(resp.Status)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	var raw map[string]interface{}
	err = json.Unmarshal(body, &raw)
	if err != nil {
		return "", err
	}
	_, ok := raw["code"]
	if ok {
		message, _ := raw["message"]
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(message)
		os.Exit(1)
	}
	token, _ := raw["security"]
	token = token.(interface{}).(map[string]interface{})["token"]
	tokenVal := fmt.Sprintf("%v", token)

	return tokenVal, err
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
	if os.Args[1] != "version" && os.Args[1] != "completion" {
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

		var token string
		var err error
		if version == "v1" {
			token, err = AuthentificationV1(url, login, password, insecure)
		} else {
			token, err = AuthentificationV2(url, login, password, insecure)
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
func getValueInFile() (string, string, string, string, string) {
	//Recover the servers from the config file
	servers := &ServerList{}
	proxy := make(map[string]string, 0)
	err := viper.Unmarshal(servers)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	_ = viper.UnmarshalKey("proxy", &proxy)

	//Search if the server(s) exist(s) in the list of servers
	index := -1
	for i, v := range servers.Servers {
		if serverName == v.Server {
			index = i
		}
	}
	if index == -1 && serverName == "" {
		for i, v := range servers.Servers {
			if "1" == v.Default {
				index = i
			}
		}
		if index == -1 {
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println(errors.New("No default server in yaml file"))
		}
	}

	//If it exists => made authentification and create token and url as environments variables
	if index >= 0 {
		name := servers.Servers[index].Server
		url := servers.Servers[index].Url
		login := servers.Servers[index].Login
		password := servers.Servers[index].Password
		version := servers.Servers[index].Version
		if os.Getenv("http_proxy") == "" {
			if proxy["httpURL"] != "" {
				if proxy["password"] != "" {
					os.Setenv("http_proxy", "http://"+proxy["user"]+":"+proxy["password"]+"@"+proxy["httpURL"])
				} else {
					os.Setenv("http_proxy", "http://"+proxy["httpURL"])
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
			if proxy["httpsURL"] != "" {
				if proxy["password"] != "" {
					os.Setenv("https_proxy", "http://"+proxy["user"]+":"+proxy["password"]+"@"+proxy["httpsURL"])
				} else {
					os.Setenv("https_proxy", "http://"+proxy["httpsURL"])
				}
			} else if len(servers.Servers[index].Proxy) != 0 && servers.Servers[index].Proxy[1].HttpsURL != "" {
				if servers.Servers[index].Proxy[3].Password != "" {
					os.Setenv("https_proxy", "http://"+servers.Servers[index].Proxy[2].User+":"+servers.Servers[index].Proxy[3].Password+"@"+servers.Servers[index].Proxy[1].HttpsURL)
				} else {
					os.Setenv("https_proxy", "http://"+servers.Servers[index].Proxy[1].HttpsURL)
				}
			}
		}
		if servers.Servers[index].Insecure != "" {
			if servers.Servers[index].Insecure == "1" {
				insecure = true
			} else {
				insecure = false
			}

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
