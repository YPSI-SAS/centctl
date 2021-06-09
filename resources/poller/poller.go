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
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//Poller represents the caracteristics of a poller
type Poller struct {
	Type     string   `json:"type" yaml:"type"`
	Label    string   `json:"label" yaml:"label"`
	Metadata Metadata `json:"metadata" yaml:"metadata"`
}

type Metadata struct {
	CentreonID string `json:"centreon-id" yaml:"centreon-id"`
	HostName   string `json:"hostname" yaml:"hostname"`
	Address    string `json:"address" yaml:"address"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name    string   `json:"name" yaml:"name"`
	Pollers []Poller `json:"pollers" yaml:"pollers"`
}

//StringText permits to display the caracteristics of the pollers to text
func (s Server) StringText() string {
	var values string = "Poller list for server" + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Pollers); i++ {
		values += "Type: " + s.Server.Pollers[i].Type + "\t"
		values += "Label: " + s.Server.Pollers[i].Label + "\t"
		values += "CentreonID: " + s.Server.Pollers[i].Metadata.CentreonID + "\t"
		values += "Hosname: " + s.Server.Pollers[i].Metadata.HostName + "\t"
		values += "Address: " + s.Server.Pollers[i].Metadata.Address + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the pollers to csv
func (s Server) StringCSV() string {
	var values string = "Server,Type,Label,CentreonID,Hostname,Address\n"
	for i := 0; i < len(s.Server.Pollers); i++ {
		values += s.Server.Name + "," + s.Server.Pollers[i].Type + "," + s.Server.Pollers[i].Label + "," + s.Server.Pollers[i].Metadata.CentreonID + "," + s.Server.Pollers[i].Metadata.HostName + "," + s.Server.Pollers[i].Metadata.Address + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the pollers to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the pollers to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
