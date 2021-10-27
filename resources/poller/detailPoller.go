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

package poller

import (
	"centctl/resources"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailPoller represents the caracteristics of a poller
type DetailPoller struct {
	Type     string         `json:"type" yaml:"type"`
	Label    string         `json:"label" yaml:"label"`
	Metadata DetailMetadata `json:"metadata" yaml:"metadata"`
}

type DetailMetadata struct {
	CentreonID string `json:"centreon-id" yaml:"centreon-id"`
	HostName   string `json:"hostname" yaml:"hostname"`
	Address    string `json:"address" yaml:"address"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name   string        `json:"name" yaml:"name"`
	Poller *DetailPoller `json:"poller" yaml:"poller"`
}

//StringText permits to display the caracteristics of the pollers to text
func (s DetailServer) StringText() string {
	var values string
	poller := s.Server.Poller
	if poller != nil {
		elements := [][]string{{"0", "Poller:"}}
		elements = append(elements, []string{"1", "Type: " + (*poller).Type})
		elements = append(elements, []string{"1", "Label: " + (*poller).Label})
		elements = append(elements, []string{"1", "CentreonID: " + (*poller).Metadata.CentreonID})
		elements = append(elements, []string{"1", "HostName: " + (*poller).Metadata.HostName})
		elements = append(elements, []string{"1", "Address: " + (*poller).Metadata.Address})
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "poller: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the pollers to csv
func (s DetailServer) StringCSV() string {
	var values string = "Server,Type,Label,CentreonID,Hostname,Address\n"
	values += s.Server.Name + ","
	poller := s.Server.Poller
	if poller != nil {
		values += "\"" + (*poller).Type + "\"" + "," + "\"" + (*poller).Label + "\"" + "," + "\"" + (*poller).Metadata.CentreonID + "\"" + "," + "\"" + (*poller).Metadata.HostName + "\"" + "," + "\"" + (*poller).Metadata.Address + "\"" + "\n"
	} else {
		values += ",,,,\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the pollers to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the pollers to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
