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

	"gopkg.in/yaml.v2"
)

//Host represents the caracteristics of a host
type Host struct {
	ID       string `json:"id" yaml:"id"`             //Host ID
	Name     string `json:"name" yaml:"name"`         //Host name
	Alias    string `json:"alias" yaml:"alias"`       //Host alias
	Address  string `json:"address" yaml:"address"`   //Host address
	Activate string `json:"activate" yaml:"activate"` //If the host is activate or not
}

//Result represents a poller array
type Result struct {
	Hosts []Host `json:"result" yaml:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name  string `json:"name" yaml:"name"`
	Hosts []Host `json:"hosts" yaml:"hosts"`
}

//StringText permits to display the caracteristics of the hosts to text
func (s Server) StringText() string {
	var values string = "Host list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Hosts); i++ {
		values += "ID: " + s.Server.Hosts[i].ID + "\t"
		values += "Name: " + s.Server.Hosts[i].Name + "\t"
		values += "Alias: " + s.Server.Hosts[i].Alias + "\t"
		values += "IP address: " + s.Server.Hosts[i].Address + "\t"
		values += "Activate: " + s.Server.Hosts[i].Activate + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the hosts to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Name,Alias,IPAddress,Activate\n"
	for i := 0; i < len(s.Server.Hosts); i++ {
		values += s.Server.Name + "," + s.Server.Hosts[i].ID + "," + s.Server.Hosts[i].Name + "," + s.Server.Hosts[i].Alias + "," + s.Server.Hosts[i].Address + "," + s.Server.Hosts[i].Activate + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the hosts to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the hosts to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
