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
