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

package service

import (
	"centctl/resources"
	"encoding/json"

	"github.com/jszwec/csvutil"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//Service represents the caracteristics of a service
type Service struct {
	ServiceID   string `json:"id" yaml:"id"`                   //Service ID
	Description string `json:"description" yaml:"description"` //Service description
	HostID      string `json:"host id" yaml:"host id"`         //Host ID of the service
	HostName    string `json:"host name" yaml:"host name"`     //Host name of the service
	Activate    string `json:"activate" yaml:"activate"`       //If the service is activate or not
}

//Result represents a poller array
type Result struct {
	Services []Service `json:"result" yaml:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name     string    `json:"name" yaml:"name"`
	Services []Service `json:"services" yaml:"services"`
}

//StringText permits to display the caracteristics of the services to text
func (s Server) StringText() string {
	var table pterm.TableData
	table = append(table, []string{"ID", "Description", "Host ID", "Host name", "Activate"})
	for i := 0; i < len(s.Server.Services); i++ {
		table = append(table, []string{s.Server.Services[i].ServiceID, s.Server.Services[i].Description, s.Server.Services[i].HostID, s.Server.Services[i].HostName, s.Server.Services[i].Activate})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the services to csv
func (s Server) StringCSV() string {
	b, _ := csvutil.Marshal(s.Server.Services)
	return string(b)
}

//StringJSON permits to display the caracteristics of the services to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the services to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
