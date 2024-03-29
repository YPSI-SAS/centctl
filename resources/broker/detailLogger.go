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
	"fmt"
	"sort"
	"strings"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

type DetailBrokerLogger struct {
	ID         string           `json:"ID" yaml:"ID"`
	BrokerName string           `json:"broker_name" yaml:"broker_name"`
	Parameters DetailParameters `json:"parameters" yaml:"parameters"`
}

//DetailResultLogger represents a poller array
type DetailResultLogger struct {
	BrokerLoggers []DetailParameter `json:"result" yaml:"result"`
}

//DetailServerLogger represents a server with informations
type DetailServerLogger struct {
	Server DetailInformationsLogger `json:"server" yaml:"server"`
}

//DetailInformationsLogger represents the informations of the server
type DetailInformationsLogger struct {
	Name         string             `json:"name" yaml:"name"`
	BrokerLogger DetailBrokerLogger `json:"broker_logger" yaml:"broker_logger"`
}

//StringText permits to display the caracteristics of the BrokerLoggers to text
func (s DetailServerLogger) StringText() string {
	var values string
	for i := 0; i < len(s.Server.BrokerLogger.Parameters); i++ {
		values += "ID: " + s.Server.BrokerLogger.Parameters[i].ParamKey + "\t"
		values += "Name: " + s.Server.BrokerLogger.Parameters[i].ParamValue + "\n"
	}
	elements := [][]string{{"0", "Broker logger:"}, {"1", "ID: " + s.Server.BrokerLogger.ID}, {"1", "BrokerName: " + s.Server.BrokerLogger.BrokerName}}
	elements = append(elements, []string{"1", "Parameters:"})
	sort.SliceStable(s.Server.BrokerLogger.Parameters, func(i, j int) bool {
		return strings.ToLower(s.Server.BrokerLogger.Parameters[i].ParamKey) < strings.ToLower(s.Server.BrokerLogger.Parameters[j].ParamKey)
	})
	for _, params := range s.Server.BrokerLogger.Parameters {
		elements = append(elements, []string{"2", params.ParamKey + " \t(value=" + params.ParamValue + ")"})
	}
	items := resources.GenerateListItems(elements, "")
	values = resources.BulletList(items)

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the BrokerLoggers to csv
func (s DetailServerLogger) StringCSV() string {
	p := []DetailBrokerLogger{s.Server.BrokerLogger}
	b, _ := csvutil.Marshal(p)
	return string(b)
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
