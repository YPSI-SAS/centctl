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

package poller

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//Poller represents the caracteristics of a poller
type Poller struct {
	ID        string `json:"id"`         //Poller ID
	Name      string `json:"name"`       //Poller Name
	IPAddress string `json:"ip address"` //IP address of the poller
}

//Result represents a poller array
type Result struct {
	Pollers []Poller `json:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name    string   `json:"name"`
	Pollers []Poller `json:"pollers"`
}

//StringText permits to display the caracteristics of the pollers to text
func (s Server) StringText() string {
	var values string = "Poller list for server" + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Pollers); i++ {
		values += "ID: " + s.Server.Pollers[i].ID + "\t"
		values += "Name: " + s.Server.Pollers[i].Name + "\t"
		values += "IP Address: " + s.Server.Pollers[i].IPAddress + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the pollers to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Name,IPAddress\n"
	for i := 0; i < len(s.Server.Pollers); i++ {
		values += s.Server.Name + "," + s.Server.Pollers[i].ID + "," + s.Server.Pollers[i].Name + "," + s.Server.Pollers[i].IPAddress + "\n"
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
