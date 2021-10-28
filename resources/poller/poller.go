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
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jszwec/csvutil"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//Poller represents the caracteristics of a poller
type Poller struct {
	ID          int    `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Address     string `json:"address" yaml:"address"`
	IsRunning   bool   `json:"is_running" yaml:"is_running"`
	LastAlive   int64  `json:"last_alive" yaml:"last_alive"`
	Version     string `json:"version" yaml:"version"`
	Description string `json:"description" yaml:"description"`
}

type ResultPoller struct {
	Pollers []Poller `json:"result" yaml:"result"`
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
	sort.SliceStable(s.Server.Pollers, func(i, j int) bool {
		return strings.ToLower(s.Server.Pollers[i].Name) < strings.ToLower(s.Server.Pollers[j].Name)
	})
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Address", "IsRunning", "LastAlive", "Version", "Description"})
	for i := 0; i < len(s.Server.Pollers); i++ {
		table = append(table, []string{strconv.Itoa(s.Server.Pollers[i].ID), s.Server.Pollers[i].Name, s.Server.Pollers[i].Address, strconv.FormatBool(s.Server.Pollers[i].IsRunning), (time.Unix(s.Server.Pollers[i].LastAlive, 0).Format(time.UnixDate)), s.Server.Pollers[i].Version, s.Server.Pollers[i].Description})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the pollers to csv
func (s Server) StringCSV() string {
	b, _ := csvutil.Marshal(s.Server.Pollers)
	return string(b)
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
