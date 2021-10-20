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

package resourceCFG

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//ResourceCFG represents the caracteristics of a resourceCFG
type ResourceCFG struct {
	ID       string   `json:"id" yaml:"id"`           //resourceCFG ID
	Name     string   `json:"name" yaml:"name"`       //resourceCFG name
	Value    string   `json:"value" yaml:"value"`     //resourceCFG value
	Comment  string   `json:"comment" yaml:"comment"` //resourceCFG comment
	Activate string   `json:"activate"`
	Instance []string `json:"instance"`
}

//Result represents a poller array
type Result struct {
	ResourceCFG []ResourceCFG `json:"result" yaml:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name        string        `json:"name" yaml:"name"`
	ResourceCFG []ResourceCFG `json:"resourceCFG" yaml:"resourceCFG"`
}

//StringText permits to display the caracteristics of the resourceCFG to text
func (s Server) StringText() string {
	var values string = "ResourceCFG list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.ResourceCFG); i++ {
		values += "ID: " + s.Server.ResourceCFG[i].ID + "\t"
		values += "Name: " + s.Server.ResourceCFG[i].Name + "\t"
		values += "Value: " + s.Server.ResourceCFG[i].Value + "\t"
		values += "Comment: " + s.Server.ResourceCFG[i].Comment + "\t"
		values += "Activate: " + s.Server.ResourceCFG[i].Activate + "\t"
		values += "Instance: "
		for index, inst := range s.Server.ResourceCFG[i].Instance {
			values += inst
			if index != len(s.Server.ResourceCFG[i].Instance)-1 {
				values += ", "
			}
		}
		values += "\n"

	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the resourceCFG to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Name,Value,Comment,Activate,Instance\n"
	for i := 0; i < len(s.Server.ResourceCFG); i++ {
		values += s.Server.Name + "," + s.Server.ResourceCFG[i].ID + "," + s.Server.ResourceCFG[i].Name + "," + s.Server.ResourceCFG[i].Value + "," + s.Server.ResourceCFG[i].Comment + "," + s.Server.ResourceCFG[i].Activate + ","
		for index, inst := range s.Server.ResourceCFG[i].Instance {
			values += inst
			if index != len(s.Server.ResourceCFG[i].Instance)-1 {
				values += "|"
			}
		}
		values += "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the resourceCFG to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the resourceCFG to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
