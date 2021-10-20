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

package centreonProxy

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

//CentreonProxy represents the caracteristics of a CentreonProxy
type CentreonProxy struct {
	URL      string `json:"url" yaml:"url"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Protocol string `json:"protocol" yaml:"protocol"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name          string         `json:"name" yaml:"name"`
	CentreonProxy *CentreonProxy `json:"centreon_proxy" yaml:"centreon_proxy"`
}

//StringText permits to display the caracteristics of the commands to text
func (s Server) StringText() string {
	var values string = "CentreonProxy list for server " + s.Server.Name + ": \n"

	centreonProxy := s.Server.CentreonProxy
	if centreonProxy != nil {
		values += "URL: " + (*centreonProxy).URL + "\t"
		values += "Port: " + strconv.Itoa((*centreonProxy).Port) + "\t"
		values += "User: " + (*centreonProxy).User + "\t"
		values += "Password: " + (*centreonProxy).Password + "\t"
		values += "Protocol: " + (*centreonProxy).Protocol + "\n"
	} else {
		values += "centreonProxy: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the commands to csv
func (s Server) StringCSV() string {
	var values string = "Server,URL,Port,User,Password,Protocol\n"
	values += s.Server.Name + ","
	centreonProxy := s.Server.CentreonProxy
	if centreonProxy != nil {
		values += (*centreonProxy).URL + ","
		values += strconv.Itoa((*centreonProxy).Port) + ","
		values += (*centreonProxy).User + ","
		values += (*centreonProxy).Password + ","
		values += (*centreonProxy).Protocol + "\n"
	} else {
		values += ",,,,\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the commands to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the commands to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
