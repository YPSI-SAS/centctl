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

package request

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBodyRequestWithoutValues(t *testing.T) {
	action := "add"
	object := "service"
	values := ""
	body, err := CreateBodyRequest(action, object, values)

	assert.NoError(t, err)
	assert.Equal(t, "{\"action\":\"add\",\"object\":\"service\"}", string(body))
}

func TestCreateBodyRequestWithoutObject(t *testing.T) {
	action := "add"
	object := ""
	values := "value1;value2"
	body, err := CreateBodyRequest(action, object, values)

	assert.NoError(t, err)
	assert.Equal(t, "{\"action\":\"add\",\"values\":\"value1;value2\"}", string(body))
}

func TestCreateBodyRequest(t *testing.T) {
	action := "add"
	object := "service"
	values := "value1;value2"
	body, err := CreateBodyRequest(action, object, values)

	assert.NoError(t, err)
	assert.Equal(t, "{\"action\":\"add\",\"object\":\"service\",\"values\":\"value1;value2\"}", string(body))
}

func TestCentreonCLAPI(t *testing.T) {
	os.Setenv("TOKEN", "tokentest=")
	requestBody, _ := CreateBodyRequest("Show", "contact", "")
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/centreonCLAPI" && r.Method == "POST" && r.Header.Get("Content-type") == "application/json" && r.Header.Get("centreon-auth-token") == "tokentest=" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`{"server": {"name": "Melissa","contacts": [{"id": "1","name": "Admin Centreon","alias": "admin","email": "centreon@ypsi.fr"},{"id": "17","name": "Guest","alias": "guest","email": "guest@localhost"}]}}`))
			}
		}),
	)
	defer ts.Close()

	client := NewClient(ts.URL + "/centreonCLAPI")

	_, body, err := client.CentreonCLAPI(requestBody)
	assert.NoError(t, err)
	assert.Equal(t, `{"server": {"name": "Melissa","contacts": [{"id": "1","name": "Admin Centreon","alias": "admin","email": "centreon@ypsi.fr"},{"id": "17","name": "Guest","alias": "guest","email": "guest@localhost"}]}}`, string(body))
}

func TestGet(t *testing.T) {
	os.Setenv("TOKEN", "tokentest=")
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/GET" && r.Method == "GET" && r.Header.Get("Content-type") == "application/json" && r.Header.Get("centreon-auth-token") == "tokentest=" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`{"server": {"name": "Melissa","hosts": [{"id": "17","name": "Hote_de_test","alias": "essai","address": "127.0.0.1","state": "0","acknowledged": "0","active_checks": "1","instance_name": "Central"}]}}`))
			}
		}),
	)
	defer ts.Close()

	client := NewClient(ts.URL + "/GET")

	_, body, err := client.Get()
	assert.NoError(t, err)
	assert.Equal(t, `{"server": {"name": "Melissa","hosts": [{"id": "17","name": "Hote_de_test","alias": "essai","address": "127.0.0.1","state": "0","acknowledged": "0","active_checks": "1","instance_name": "Central"}]}}`, string(body))
}

func TestNamePollerHost(t *testing.T) {
	hostName := "hostTEST"
	os.Setenv("TOKEN", "tokentest=")
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/NamePoller" && r.Method == "GET" && r.Header.Get("Content-type") == "application/json" && r.Header.Get("centreon-auth-token") == "tokentest=" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`[{"id":"17","name":"hostTEST","alias":"hostTEST","address":"127.0.0.1","state":"0","acknowledged":"0","instance_name":"pollerTEST","active_checks":"1"}]`))
			}
		}),
	)
	defer ts.Close()

	client := NewClient(ts.URL + "/NamePoller")

	pollerName, err := client.NamePollerHost(hostName, false)
	assert.NoError(t, err)
	assert.Equal(t, "pollerTEST", pollerName)
}

func TestNamePollerHostWithHostIncorrect(t *testing.T) {
	hostName := "hostTEST"
	os.Setenv("TOKEN", "tokentest=")
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/NamePoller" && r.Method == "GET" && r.Header.Get("Content-type") == "application/json" && r.Header.Get("centreon-auth-token") == "tokentest=" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`[{"id":"17","name":"hosTEST","alias":"hostTEST","address":"127.0.0.1","state":"0","acknowledged":"0","instance_name":"pollerTEST","active_checks":"1"}]`))
			}
		}),
	)
	defer ts.Close()

	client := NewClient(ts.URL + "/NamePoller")

	_, err := client.NamePollerHost(hostName, false)
	assert.EqualError(t, err, "the name of host is incorrect")
}

func TestExportConf(t *testing.T) {
	pollerName := "pollerTEST"
	os.Setenv("TOKEN", "tokentest=")
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/exportConf" && r.Method == "POST" && r.Header.Get("Content-type") == "application/json" && r.Header.Get("centreon-auth-token") == "tokentest=" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(`result`))
			}
		}),
	)
	defer ts.Close()

	client := NewClient(ts.URL + "/exportConf")

	err := client.ExportConf(pollerName, false)

	assert.NoError(t, err, nil)
}
