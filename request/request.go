package request

import (
	"bytes"
	"centctl/debug"
	"centctl/host"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//CreateBodyRequest permits to create a body request for the request next
func CreateBodyRequest(action string, object string, values string) ([]byte, error) {
	var requestBody []byte
	var err error
	if values == "" && object != "" {
		requestBody, err = json.Marshal(map[string]string{
			"action": action,
			"object": object,
		})
	} else if object == "" {
		requestBody, err = json.Marshal(map[string]string{
			"action": action,
			"values": values,
		})
	} else if object != "" && values != "" {

		requestBody, err = json.Marshal(map[string]string{
			"action": action,
			"object": object,
			"values": values,
		})
	}
	return requestBody, err
}

type client struct {
	http.Client
	url string
}

//Client is the interface which implements the functions request
type Client interface {
	CentreonCLAPI(requestBody []byte) (int, []byte, error)
	Get() (int, []byte, error)
	NamePollerHost(hostName string, debugV bool) (string, error)
	ExportConf(pollerName string, debugV bool) error
}

//NewClient permits to create a new client associate with an url
func NewClient(url string) Client {
	return &client{
		http.Client{
			Timeout: time.Duration(30) * time.Second,
		},
		url,
	}
}

//CentreonCLAPI permits to send a request to CentreonAPI
func (c *client) CentreonCLAPI(requestBody []byte) (int, []byte, error) {
	request, err := http.NewRequest("POST", c.url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("centreon-auth-token", os.Getenv("TOKEN"))

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

//Get permits to send a request get to CentreonAPI
func (c *client) Get() (int, []byte, error) {
	request, err := http.NewRequest("GET", c.url, nil)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("centreon-auth-token", os.Getenv("TOKEN"))
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

//NamePollerHost permits to retrieve the name of the poller of this host
func (c *client) NamePollerHost(hostName string, debugV bool) (string, error) {
	statusCode, body, err := c.Get()
	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("search name poller host", "", c.url, statusCode, body)
	}
	if err != nil {
		return "", nil
	}

	var hosts []host.RealtimeHost
	json.Unmarshal(body, &hosts)
	pollerName := ""
	for _, val := range hosts {
		if val.Name == hostName {
			pollerName = val.PollerName
		}
	}
	if pollerName == "" {

		return "", fmt.Errorf("the name of host is incorrect")
	}
	return pollerName, nil
}

//ExportConf permits to export the configuration of one poller
func (c *client) ExportConf(pollerName string, debugV bool) error {
	requestBody, err := CreateBodyRequest("APPLYCFG", "", pollerName)
	if err != nil {
		return err
	}

	statusCode, body, err := c.CentreonCLAPI(requestBody)
	//If flag debug, print informations about the request API
	if debugV {
		debug.Show("export configuration", string(requestBody), c.url, statusCode, body)
	}
	if err != nil {
		return err
	}

	if !strings.Contains(string(body), "result") {
		fmt.Println("erreur: ", string(body))
		os.Exit(1)
	}
	fmt.Printf("The configuration of the poller %v is exported\n", pollerName)
	return nil
}
