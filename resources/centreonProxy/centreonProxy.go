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

package centreonProxy

import (
	"centctl/resources"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jszwec/csvutil"
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
	var values string

	centreonProxy := s.Server.CentreonProxy
	if centreonProxy != nil {
		elements := [][]string{{"0", "centreonProxy:"}, {"1", "URL: " + (*centreonProxy).URL}, {"1", "Port: " + strconv.Itoa((*centreonProxy).Port)}, {"1", "User: " + (*centreonProxy).User}, {"1", "Password: " + (*centreonProxy).Password}, {"1", "Protocol: " + (*centreonProxy).Protocol}}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "centreonProxy: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the commands to csv
func (s Server) StringCSV() string {
	var p []CentreonProxy
	if s.Server.CentreonProxy != nil {
		p = append(p, *s.Server.CentreonProxy)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
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
