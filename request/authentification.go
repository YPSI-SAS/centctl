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

package request

import (
	"bytes"
	"centctl/colorMessage"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//Token represents the generate token by the authentification
type Token struct {
	Token string `json:"authToken"`
}

//AuthentificationV1 allow the authentification at the server specified with API v1
func AuthentificationV1(urlServer string, login string, password string, insecure bool) (string, error) {
	colorRed := colorMessage.GetColorRed()
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
func AuthentificationV2(urlServer string, login string, password string, insecure bool, versionAPI string) (string, string, error) {
	colorRed := colorMessage.GetColorRed()
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
		return "", "", err
	}

REQUEST:
	urlCentreon := urlServer + "/api" + versionAPI + "/login"

	if insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	resp, err := http.Post(urlCentreon, "application/json",
		bytes.NewBuffer(requestBody))

	if err != nil {
		return "", "", err
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
		return "", "", err
	}
	_, ok := raw["code"]
	if ok {
		message, _ := raw["message"]
		if strings.Contains(message.(string), "No route found for") {
			versionAPI = "/latest"
			goto REQUEST
		}
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(message)
		os.Exit(1)
	}
	token, _ := raw["security"]
	token = token.(interface{}).(map[string]interface{})["token"]
	tokenVal := fmt.Sprintf("%v", token)

	return tokenVal, versionAPI, err
}
