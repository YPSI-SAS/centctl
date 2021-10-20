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

package broker

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailBrokerLogger represents the caracteristics of a BrokerLogger
type DetailBrokerLogger struct {
	ParamKey   string `json:"parameter key" yaml:"parameter key"`     //BrokerLogger param key
	ParamValue string `json:"parameter value" yaml:"parameter value"` //BrokerLogger param value
}

//DetailResultLogger represents a poller array
type DetailResultLogger struct {
	BrokerLoggers []DetailBrokerLogger `json:"result" yaml:"result"`
}

//DetailServerLogger represents a server with informations
type DetailServerLogger struct {
	Server DetailInformationsLogger `json:"server" yaml:"server"`
}

//DetailInformationsLogger represents the informations of the server
type DetailInformationsLogger struct {
	Name         string               `json:"name" yaml:"name"`
	BrokerLogger []DetailBrokerLogger `json:"broker_logger" yaml:"broker_logger"`
}

//StringText permits to display the caracteristics of the BrokerLoggers to text
func (s DetailServerLogger) StringText() string {
	var values string = "BrokerLogger list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.BrokerLogger); i++ {
		values += "ID: " + s.Server.BrokerLogger[i].ParamKey + "\t"
		values += "Name: " + s.Server.BrokerLogger[i].ParamValue + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the BrokerLoggers to csv
func (s DetailServerLogger) StringCSV() string {
	var values string = "Server,ID,Name\n"
	values += s.Server.Name + ","
	for i := 0; i < len(s.Server.BrokerLogger); i++ {
		values += s.Server.BrokerLogger[i].ParamKey + "," + s.Server.BrokerLogger[i].ParamValue + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the BrokerLoggers to json
func (s DetailServerLogger) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the BrokerLoggers to yaml
func (s DetailServerLogger) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
