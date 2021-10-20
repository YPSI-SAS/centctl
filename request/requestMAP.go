/*
MIT License

Copyright (c)  2020-2021 YPSI SAS


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
	"centctl/debug"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type authMap struct {
	StudioSession string `json:"studioSession"`
	Login         string `json:"login"`
}

type serverVersion struct {
	ServerVersion string `json:"server-version"`
}

type clientMAP struct {
	http.Client
	url string
}

//Client is the interface which implements the functions request
type ClientMAP interface {
	GetSessionToken([]byte, bool) error
	GetServerVersion(bool) error
	Post([]byte) (int, []byte, error)
	Get() (int, []byte, error)
	Put() (int, []byte, error)
	Delete() (int, []byte, error)
}

//NewClientMAP permits to create a new client associate with an url
func NewClientMAP(url string) ClientMAP {
	return &clientMAP{
		http.Client{
			Timeout: time.Duration(30) * time.Second,
		},
		url,
	}
}

//GetSessionToken permits to send a request to the API MAP for get the studio-token
func (c *clientMAP) GetSessionToken(requestBody []byte, debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	request, err := http.NewRequest("POST", c.url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		return err
	}
	resp, err := c.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//If flag debug, print informations about the request
	if debugV {
		debug.Show("authentification map", string(requestBody), c.url, resp.StatusCode, body)
	}
	if err != nil {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(err.Error())
	}
	if resp.StatusCode != 200 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(string(body))
		os.Exit(1)
	}
	authMap := authMap{}
	json.Unmarshal(body, &authMap)

	os.Setenv("STUDIO_TOKEN", authMap.StudioSession)

	return nil
}

//GetServerVersion permits to get server's version
func (c *clientMAP) GetServerVersion(debugV bool) error {
	colorRed := colorMessage.GetColorRed()
	request, err := http.NewRequest("GET", c.url, nil)
	request.Header.Set("Content-type", "application/json")
	if err != nil {
		return err
	}
	resp, err := c.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//If flag debug, print informations about the request
	if debugV {
		debug.Show("server version map", "", c.url, resp.StatusCode, body)
	}
	if err != nil {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(err.Error())
	}
	serverVersion := serverVersion{}
	json.Unmarshal(body, &serverVersion)
	os.Setenv("SERVER_VERSION", serverVersion.ServerVersion)

	return nil
}

//POST permits to send a post request to the API MAP
func (c *clientMAP) Post(requestBody []byte) (int, []byte, error) {
	request, err := http.NewRequest("POST", c.url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("studio-session", os.Getenv("STUDIO_TOKEN"))
	request.Header.Set("X-Client-Version", os.Getenv("SERVER_VERSION"))

	if err != nil {
		return 0, nil, err
	}
	resp, err := c.Do(request)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, body, nil
}

//Get permits to send a GET request to the API MAP
func (c *clientMAP) Get() (int, []byte, error) {
	request, err := http.NewRequest("GET", c.url, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("studio-session", os.Getenv("STUDIO_TOKEN"))
	request.Header.Set("X-Client-Version", os.Getenv("SERVER_VERSION"))
	if err != nil {
		return 0, nil, err
	}
	resp, err := c.Do(request)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, body, nil
}

//Put permits to send a PUT request to the API MAP
func (c *clientMAP) Put() (int, []byte, error) {
	request, err := http.NewRequest(http.MethodPut, c.url, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("studio-session", os.Getenv("STUDIO_TOKEN"))
	request.Header.Set("X-Client-Version", os.Getenv("SERVER_VERSION"))
	if err != nil {
		return 0, nil, err
	}
	resp, err := c.Do(request)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, body, nil
}

//Delete permits to send a DELETE request to the API MAP
func (c *clientMAP) Delete() (int, []byte, error) {
	request, err := http.NewRequest(http.MethodDelete, c.url, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("studio-session", os.Getenv("STUDIO_TOKEN"))
	request.Header.Set("X-Client-Version", os.Getenv("SERVER_VERSION"))
	if err != nil {
		return 0, nil, err
	}
	resp, err := c.Do(request)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, body, nil
}
