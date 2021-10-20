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

package host

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

//RealtimeHostV2 represents the caracteristics of a host
type RealtimeHostV2 struct {
	ID           int              `json:"id" yaml:"hosts"`                    //Host ID
	Name         string           `json:"name" yaml:"name"`                   //Host name
	Alias        string           `json:"alias" yaml:"alias"`                 //Host alias
	Address      string           `json:"fqdn" yaml:"fqdn"`                   //Host address
	Status       RealtimeStatusV2 `json:"status" yaml:"status"`               //State of the host
	Acknowledged bool             `json:"acknowledged" yaml:"acknowledged"`   //If the host is acknowledge or not
	ActiveCheck  bool             `json:"active_checks" yaml:"active_checks"` //If the host is active or not
	PollerID     int              `json:"poller_id" yaml:"poller_id"`         //Poller ID
}

type RealtimeStatusV2 struct {
	Code         int    `json:"code" yaml:"code"`
	Name         string `json:"name" yaml:"name"`
	SeverityCode int    `json:"severity_code" yaml:"severity_code"`
}

//RealtimeServer represents a server with informations
type RealtimeServerV2 struct {
	Server RealtimeInformationsV2 `json:"server" yaml:"server"`
}

//RealtimeInformations represents the informations of the server
type RealtimeInformationsV2 struct {
	Name  string           `json:"name" yaml:"name"`
	Hosts []RealtimeHostV2 `json:"hosts" yaml:"hosts"`
}

type RealtimeResultBodyV2 struct {
	ListHosts []RealtimeHostV2 `json:"result" yaml:"result"`
}

//StringText permits to display the caracteristics of the hosts to text
func (s RealtimeServerV2) StringText() string {
	var values string = "Host list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Hosts); i++ {
		values += "ID: " + strconv.Itoa(s.Server.Hosts[i].ID) + "\t"
		values += "Name: " + s.Server.Hosts[i].Name + "\t"
		values += "Alias: " + s.Server.Hosts[i].Alias + "\t"
		values += "IP address: " + s.Server.Hosts[i].Address + "\t"
		values += "Status code: " + strconv.Itoa(s.Server.Hosts[i].Status.Code) + "\t"
		values += "Status name: " + s.Server.Hosts[i].Status.Name + "\t"
		values += "Acknowledged: " + strconv.FormatBool(s.Server.Hosts[i].Acknowledged) + "\t"
		values += "ActiveCheck: " + strconv.FormatBool(s.Server.Hosts[i].ActiveCheck) + "\t"
		values += "Poller ID: " + strconv.Itoa(s.Server.Hosts[i].PollerID) + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the hosts to csv
func (s RealtimeServerV2) StringCSV() string {
	var values string = "Server,ID,Name,Alias,IPAddress,StatusCode,StatusName,Acknowledged,ActiveCheck,PollerID	\n"
	for i := 0; i < len(s.Server.Hosts); i++ {
		values += "\"" + s.Server.Name + "\"" + "," + "\"" + strconv.Itoa(s.Server.Hosts[i].ID) + "\"" + "," + "\"" + s.Server.Hosts[i].Name + "\"" + "," + "\"" + s.Server.Hosts[i].Alias + "\"" + "," + "\"" + s.Server.Hosts[i].Address + "\"" + "," + "\"" + strconv.Itoa(s.Server.Hosts[i].Status.Code) + "\"" + "," + "\"" + s.Server.Hosts[i].Status.Name + "\"" + "," + "\"" + strconv.FormatBool(s.Server.Hosts[i].Acknowledged) + "\"" + "," + "\"" + strconv.FormatBool(s.Server.Hosts[i].ActiveCheck) + "\"" + "," + "\"" + strconv.Itoa(s.Server.Hosts[i].PollerID) + "\"" + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the hosts to json
func (s RealtimeServerV2) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the hosts to yaml
func (s RealtimeServerV2) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
