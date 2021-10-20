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

package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

//RealtimeService represents the caracteristics of a service
type RealtimeService struct {
	ServiceID    int            `json:"id" yaml:"id"`                       //Service ID
	Name         string         `json:"name" yaml:"name"`                   //Service description
	Parent       RealtimeParent `json:"parent" yaml:"parent"`               //Parent of service
	Status       RealtimeStatus `json:"status" yaml:"status"`               //State of the service
	Information  string         `json:"information" yaml:"information"`     //Srevice Output
	Acknowledged bool           `json:"acknowledged" yaml:"acknowledged"`   //If the service is acknowledge or not
	ActiveCheck  bool           `json:"active_checks" yaml:"active_checks"` //If the service is activate or not
}

type RealtimeParent struct {
	ID       int    `json:"id" yaml:"id"`
	Name     string `json:"name" yaml:"name"`
	Address  string `json:"fqdn" yaml:"fqdn"`
	PollerID int    `json:"poller_id" yaml:"poller_id"` //Poller ID
}

type RealtimeStatus struct {
	Code         int    `json:"code" yaml:"code"`
	Name         string `json:"name" yaml:"name"`
	SeverityCode int    `json:"severity_code" yaml:"severity_code"`
}

//RealtimeServer represents a server with informations
type RealtimeServer struct {
	Server RealtimeInformations `json:"server" yaml:"server"`
}

//RealtimeInformations represents the informations of the server
type RealtimeInformations struct {
	Name     string            `json:"name" yaml:"name"`
	Services []RealtimeService `json:"services" yaml:"services"`
}

type RealtimeResultBody struct {
	ListServices []RealtimeService `json:"result" yaml:"result"`
}

//StringText permits to display the caracteristics of the services to text
func (s RealtimeServer) StringText() string {
	var values string = "Service list for server=" + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Services); i++ {
		values += "ID: " + strconv.Itoa(s.Server.Services[i].ServiceID) + "\t"
		values += "Name: " + s.Server.Services[i].Name + "\t"
		values += "Parent ID: " + strconv.Itoa(s.Server.Services[i].Parent.ID) + "\t"
		values += "Parent name: " + s.Server.Services[i].Parent.Name + "\t"
		values += "Parent address: " + s.Server.Services[i].Parent.Address + "\t"
		values += "Parent pollerID: " + strconv.Itoa(s.Server.Services[i].Parent.PollerID) + "\t"
		values += "Status code: " + strconv.Itoa(s.Server.Services[i].Status.Code) + "\t"
		values += "Status name: " + s.Server.Services[i].Status.Name + "\t"
		values += "Information: " + s.Server.Services[i].Information + "\t"
		values += "Acknowledged: " + strconv.FormatBool(s.Server.Services[i].Acknowledged) + "\t"
		values += "ActiveCheck: " + strconv.FormatBool(s.Server.Services[i].ActiveCheck) + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the services to csv
func (s RealtimeServer) StringCSV() string {
	var values string = "Server,ID,Name,ParentID,ParentName,ParentPollerID,ParentAddress,StatusCode,StatusName,Information,Acknowledged,Activate\n"
	for i := 0; i < len(s.Server.Services); i++ {
		values += s.Server.Name + "," + strconv.Itoa(s.Server.Services[i].ServiceID) + "," + s.Server.Services[i].Name + "," + strconv.Itoa(s.Server.Services[i].Parent.PollerID) + "," + strconv.Itoa(s.Server.Services[i].Parent.ID) + "," + s.Server.Services[i].Parent.Name + "," + s.Server.Services[i].Parent.Address + "," + strconv.Itoa(s.Server.Services[i].Status.Code) + "," + s.Server.Services[i].Status.Name + "," + s.Server.Services[i].Information + "," + strconv.FormatBool(s.Server.Services[i].Acknowledged) + "," + strconv.FormatBool(s.Server.Services[i].ActiveCheck) + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the services to json
func (s RealtimeServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the services to yaml
func (s RealtimeServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
