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
	"centctl/debug"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type clientV2 struct {
	http.Client
	url string
}

//Client is the interface which implements the functions request
type ClientV2 interface {
	Get() (int, []byte, error)
	Put(requestBody []byte) (int, []byte, error)
	Post(requestBody []byte) (int, []byte, error)
}

//NewClientV2 permits to create a new client associate with an url
func NewClientV2(url string) ClientV2 {
	return &clientV2{
		http.Client{
			Timeout: time.Duration(30) * time.Second,
		},
		url,
	}
}

//Get permits to send a request get to CentreonAPI
func (c *clientV2) Get() (int, []byte, error) {
	request, err := http.NewRequest("GET", c.url, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("X-AUTH-TOKEN", os.Getenv("TOKEN"))
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

type Poller struct {
	ID int `json:"poller_id"`
}

func (c *clientV2) Post(requestBody []byte) (int, []byte, error) {
	request, err := http.NewRequest("POST", c.url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("X-AUTH-TOKEN", os.Getenv("TOKEN"))

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

func (c *clientV2) Put(requestBody []byte) (int, []byte, error) {
	request, err := http.NewRequest("PUT", c.url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("X-AUTH-TOKEN", os.Getenv("TOKEN"))

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

//IDPollerHost permits to retrieve the ID of the poller of this host
func IDPollerHost(hostID int, debugV bool) (int, error) {
	urlCentreon := os.Getenv("URL") + "/api/beta/monitoring/hosts/" + strconv.Itoa(hostID)
	client := NewClientV2(urlCentreon)
	statusCode, body, err := client.Get()
	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("search ID poller host", "", urlCentreon, statusCode, body)
	}
	if err != nil {
		return -1, nil
	}

	var pollerHost Poller
	pollerHost.ID = -1
	json.Unmarshal(body, &pollerHost)

	if pollerHost.ID == -1 {

		return -1, fmt.Errorf("the ID of host is incorrect")
	}
	return pollerHost.ID, nil
}
