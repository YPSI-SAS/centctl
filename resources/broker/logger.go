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

package broker

import (
	"centctl/resources"
	"encoding/json"

	"github.com/jszwec/csvutil"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//BrokerLogger represents the caracteristics of a BrokerLogger
type BrokerLogger struct {
	ID   string `json:"id" yaml:"id"`     //BrokerLogger ID
	Name string `json:"name" yaml:"name"` //BrokerLogger name
}

//ResultLogger represents a poller array
type ResultLogger struct {
	BrokerLoggers []BrokerLogger `json:"result" yaml:"result"`
}

//ServerLogger represents a server with informations
type ServerLogger struct {
	Server InformationsLogger `json:"server" yaml:"server"`
}

//InformationsLogger represents the informations of the server
type InformationsLogger struct {
	Name          string         `json:"name" yaml:"name"`
	BrokerLoggers []BrokerLogger `json:"broker_loggers" yaml:"broker_loggers"`
}

//StringText permits to display the caracteristics of the BrokerLoggers to text
func (s ServerLogger) StringText() string {
	var table pterm.TableData
	table = append(table, []string{"ID", "Name"})
	for i := 0; i < len(s.Server.BrokerLoggers); i++ {
		table = append(table, []string{s.Server.BrokerLoggers[i].ID, s.Server.BrokerLoggers[i].Name})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the BrokerLoggers to csv
func (s ServerLogger) StringCSV() string {
	b, _ := csvutil.Marshal(s.Server.BrokerLoggers)
	return string(b)
}

//StringJSON permits to display the caracteristics of the BrokerLoggers to json
func (s ServerLogger) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the BrokerLoggers to yaml
func (s ServerLogger) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
