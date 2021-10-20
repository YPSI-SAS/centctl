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

	"gopkg.in/yaml.v2"
)

//DetailService represents the caracteristics of a service
type DetailService struct {
	ID                   string `json:"id" yaml:"id"`                   //service ID
	Description          string `json:"description" yaml:"description"` //service name
	HostID               string `json:"host id" yaml:"host id"`
	HostName             string `json:"host name" yaml:"host name"`
	CheckCommand         string `json:"check command" yaml:"check command"`
	CheckCommandArg      string `json:"check command arg" yaml:"check command arg"`
	NormalCheckInterval  string `json:"normal check interval" yaml:"normal check interval"`
	RetryCheckInterval   string `json:"retry check interval" yaml:"retry check interval"`
	MaxCheckAttempts     string `json:"max check attempts" yaml:"max check attempts"`
	ActiveChecksEnabled  string `json:"active checks enabled" yaml:"active checks enabled"`
	PassiveChecksEnabled string `json:"passive checks enabled" yaml:"passive checks enabled"`
	Activate             string `json:"activate" yaml:"activate"`
}

//DetailResult represents a service Group array
type DetailResult struct {
	DetailServices []DetailService `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name    string         `json:"name" yaml:"name"`
	Service *DetailService `json:"service" yaml:"service"`
}

//StringText permits to display the caracteristics of the hosts to text
func (s DetailServer) StringText() string {
	var values string = "service detail for server " + s.Server.Name + ": \n"
	service := s.Server.Service
	if service != nil {
		values += "ID: " + (*service).ID + "\t"
		values += "Description: " + (*service).Description + "\t"
		values += "HostID: " + (*service).HostID + "\t"
		values += "HostName: " + (*service).HostName + "\t"
		values += "CheckCommand: " + (*service).CheckCommand + "\t"
		values += "CheckCommandArg: " + (*service).CheckCommandArg + "\t"
		values += "NormalCheckInterval: " + (*service).NormalCheckInterval + "\t"
		values += "RetryCheckInterval: " + (*service).RetryCheckInterval + "\t"
		values += "MaxCheckAttempts: " + (*service).MaxCheckAttempts + "\t"
		values += "ActiveChecksEnabled: " + (*service).ActiveChecksEnabled + "\t"
		values += "PassiveChecksEnabled: " + (*service).PassiveChecksEnabled + "\t"
		values += "Activate: " + (*service).Activate + "\n"
	} else {
		values += "service: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the hosts to csv
func (s DetailServer) StringCSV() string {
	var values string = "Server,ID,Description,HostID,HostName,CheckCommand,CheckCommandArg,NormalCheckInterval,RetryCheckInterval,MaxCheckAttempts,ActiveChecksEnabled,PassiveChecksEnabled,Activate\n"
	values += s.Server.Name + ","
	service := s.Server.Service
	if service != nil {
		values += (*service).ID + ","
		values += (*service).Description + ","
		values += (*service).HostID + ","
		values += (*service).HostName + ","
		values += (*service).CheckCommand + ","
		values += (*service).CheckCommandArg + ","
		values += (*service).NormalCheckInterval + ","
		values += (*service).RetryCheckInterval + ","
		values += (*service).MaxCheckAttempts + ","
		values += (*service).ActiveChecksEnabled + ","
		values += (*service).PassiveChecksEnabled + ","
		values += (*service).Activate + "\n"

	} else {
		values += ",,,,,,,,,,,\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the hosts to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the hosts to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
