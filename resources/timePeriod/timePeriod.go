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

package timePeriod

import (
	"centctl/resources"
	"encoding/json"

	"github.com/jszwec/csvutil"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//TimePeriod represents the caracteristics of a TimePeriod
type TimePeriod struct {
	ID        string `json:"id" yaml:"id"`       //TimePeriod ID
	Name      string `json:"name" yaml:"name"`   //TimePeriod name
	Alias     string `json:"alias" yaml:"alias"` //TimePeriod expression
	Monday    string `json:"monday" yaml:"monday"`
	Tuesday   string `json:"tuesday" yaml:"tuesday"`
	Wednesday string `json:"wednesday" yaml:"wednesday"`
	Thursday  string `json:"thursday" yaml:"thursday"`
	Friday    string `json:"friday" yaml:"friday"`
	Saturday  string `json:"saturday" yaml:"saturday"`
	Sunday    string `json:"sunday" yaml:"sunday"`
}

//Result represents a poller array
type Result struct {
	TimePeriods []TimePeriod `json:"result" yaml:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name        string       `json:"name" yaml:"name"`
	TimePeriods []TimePeriod `json:"timePeriods" yaml:"timePeriods"`
}

//StringText permits to display the caracteristics of the TimePeriods to text
func (s Server) StringText() string {
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Alias", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"})
	for i := 0; i < len(s.Server.TimePeriods); i++ {
		table = append(table, []string{s.Server.TimePeriods[i].ID, s.Server.TimePeriods[i].Name, s.Server.TimePeriods[i].Alias, s.Server.TimePeriods[i].Monday, s.Server.TimePeriods[i].Tuesday, s.Server.TimePeriods[i].Wednesday, s.Server.TimePeriods[i].Thursday, s.Server.TimePeriods[i].Friday, s.Server.TimePeriods[i].Saturday, s.Server.TimePeriods[i].Sunday})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the TimePeriods to csv
func (s Server) StringCSV() string {
	b, _ := csvutil.Marshal(s.Server.TimePeriods)
	return string(b)
}

//StringJSON permits to display the caracteristics of the TimePeriods to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the TimePeriods to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
