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
	"centctl/request"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt password in config file centctl.yaml",
	Long:  `Encrypt password in config file centctl.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		encrypt()
	},
}

//modifyPassword permits to modify the password in server struct
func (s *ServerList) modifyPassword(index int, newPassword string) {
	s.Servers[index].Password = newPassword
}

//encrypt permits to encrypt the passwords in config file
func encrypt() {
	//Get or generage encryption key
	var keyString string
	if os.Getenv("CENTCTL_DECRYPT_KEY") == "" {
		keyString = request.GenerateRandomString(150)
	} else {
		keyString = os.Getenv("CENTCTL_DECRYPT_KEY")
	}

	//Get config file
	confFile := os.Getenv("CENTCTL_CONF")
	colorRed := colorMessage.GetColorRed()

	//Get elements in config file
	servers := &ServerList{}
	if confFile == "" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(errors.New("No config file"))
		os.Exit(1)
	}
	viper.SetConfigFile(confFile)
	proxy := &ProxyStruct{}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}

	err := viper.Unmarshal(servers)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	err = viper.Unmarshal(proxy)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	//Encrypt password not already encrypted
	for i, server := range servers.Servers {
		if !strings.HasPrefix(server.Password, "CENTCRYPT_") {
			ciphertext := request.Encrypt(server.Password, keyString)
			servers.modifyPassword(i, "CENTCRYPT_"+ciphertext)
		}
	}

	//Write in config file the new passwords
	file, err := os.OpenFile(confFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	defer file.Close()

	dataS, _ := yaml.Marshal(&servers)
	if len(proxy.Proxy) != 0 {
		proxyS, _ := yaml.Marshal(&proxy)
		file.Write(proxyS)
	}

	file.Write(dataS)

	fmt.Println("Your key is : " + keyString)
	fmt.Printf(colorRed, "WARNING: ")
	fmt.Println(errors.New("Save the key into env variable (CENTCTL_DECRYPT_KEY) !!"))
}

func init() {

}
