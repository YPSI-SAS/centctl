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

package timePeriod

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailTimePeriod represents the caracteristics of a TimePeriod
type DetailTimePeriod struct {
	ID         string                      `json:"id" yaml:"id"`       //TimePeriod ID
	Name       string                      `json:"name" yaml:"name"`   //TimePeriod name
	Alias      string                      `json:"alias" yaml:"alias"` //TimePeriod expression
	Monday     string                      `json:"monday" yaml:"monday"`
	Tuesday    string                      `json:"tuesday" yaml:"tuesday"`
	Wednesday  string                      `json:"wednesday" yaml:"wednesday"`
	Thursday   string                      `json:"thursday" yaml:"thursday"`
	Friday     string                      `json:"friday" yaml:"friday"`
	Saturday   string                      `json:"saturday" yaml:"saturday"`
	Sunday     string                      `json:"sunday" yaml:"sunday"`
	Exceptions []DetailTimePeriodException `json:"exceptions" yaml:"exceptions"`
}

//DetailResult represents a poller array
type DetailResult struct {
	TimePeriods []DetailTimePeriod `json:"result" yaml:"result"`
}

//DetailTimePeriodException represents the caracteristics of an exception
type DetailTimePeriodException struct {
	Days      string `json:"days" yaml:"days"`
	Timerange string `json:"timerange" yaml:"timerange"`
}

//DetailResultException represents a poller array
type DetailResultException struct {
	TimePeriodExceptions []DetailTimePeriodException `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name        string            `json:"name" yaml:"name"`
	TimePeriods *DetailTimePeriod `json:"timePeriods" yaml:"timePeriods"`
}

//StringText permits to display the caracteristics of the TimePeriods to text
func (s DetailServer) StringText() string {
	var values string = "TimePeriod list for server " + s.Server.Name + ": \n"

	timePeriod := s.Server.TimePeriods
	if timePeriod != nil {
		values += "ID: " + (*timePeriod).ID + "\t"
		values += "Name: " + (*timePeriod).Name + "\t"
		values += "Alias: " + (*timePeriod).Alias + "\t"
		values += "Monday: " + (*timePeriod).Monday + "\t"
		values += "Tuesday: " + (*timePeriod).Tuesday + "\t"
		values += "Wednesday: " + (*timePeriod).Wednesday + "\t"
		values += "Thursday: " + (*timePeriod).Thursday + "\t"
		values += "Friday: " + (*timePeriod).Friday + "\t"
		values += "Saturday: " + (*timePeriod).Saturday + "\t"
		values += "Sunday: " + (*timePeriod).Sunday + "\n"
	} else {
		values += "timePeriod: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the TimePeriods to csv
func (s DetailServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday\n"
	values += s.Server.Name + ","
	timePeriod := s.Server.TimePeriods
	if timePeriod != nil {
		values += (*timePeriod).ID + ","
		values += (*timePeriod).Name + ","
		values += (*timePeriod).Alias + ","
		values += (*timePeriod).Monday + ","
		values += (*timePeriod).Tuesday + ","
		values += (*timePeriod).Wednesday + ","
		values += (*timePeriod).Thursday + ","
		values += (*timePeriod).Friday + ","
		values += (*timePeriod).Saturday + ","
		values += (*timePeriod).Sunday + "\n"
	} else {
		values += ",,,,,,,,,\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the TimePeriods to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the TimePeriods to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
