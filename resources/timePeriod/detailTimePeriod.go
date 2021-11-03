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
	"fmt"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailTimePeriod represents the caracteristics of a TimePeriod
type DetailTimePeriod struct {
	ID         string                     `json:"id" yaml:"id"`       //TimePeriod ID
	Name       string                     `json:"name" yaml:"name"`   //TimePeriod name
	Alias      string                     `json:"alias" yaml:"alias"` //TimePeriod expression
	Monday     string                     `json:"monday" yaml:"monday"`
	Tuesday    string                     `json:"tuesday" yaml:"tuesday"`
	Wednesday  string                     `json:"wednesday" yaml:"wednesday"`
	Thursday   string                     `json:"thursday" yaml:"thursday"`
	Friday     string                     `json:"friday" yaml:"friday"`
	Saturday   string                     `json:"saturday" yaml:"saturday"`
	Sunday     string                     `json:"sunday" yaml:"sunday"`
	Exceptions DetailTimePeriodExceptions `json:"exceptions" yaml:"exceptions"`
}

type DetailTimePeriodExceptions []DetailTimePeriodException

func (t DetailTimePeriodExceptions) MarshalCSV() ([]byte, error) {
	var value string
	for i, parent := range t {
		value += parent.Days + "|" + parent.Timerange
		if i < len(t)-1 {
			value += ","
		}
	}
	return []byte(value), nil
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
	var values string

	timePeriod := s.Server.TimePeriods
	if timePeriod != nil {
		elements := [][]string{{"0", "TimePeriod:"}}
		elements = append(elements, []string{"1", "ID: " + (*timePeriod).ID})
		elements = append(elements, []string{"1", "Name: " + (*timePeriod).Name})
		elements = append(elements, []string{"1", "Alias: " + (*timePeriod).Alias})
		elements = append(elements, []string{"1", "Monday: " + (*timePeriod).Monday})
		elements = append(elements, []string{"1", "Tuesday: " + (*timePeriod).Tuesday})
		elements = append(elements, []string{"1", "Wednesday: " + (*timePeriod).Wednesday})
		elements = append(elements, []string{"1", "Thursday: " + (*timePeriod).Thursday})
		elements = append(elements, []string{"1", "Friday: " + (*timePeriod).Friday})
		elements = append(elements, []string{"1", "Saturday: " + (*timePeriod).Saturday})
		elements = append(elements, []string{"1", "Sunday: " + (*timePeriod).Sunday})

		if len((*timePeriod).Exceptions) == 0 {
			elements = append(elements, []string{"1", "Exceptions: []"})
		} else {
			elements = append(elements, []string{"1", "Exceptions:"})
			for _, server := range (*timePeriod).Exceptions {
				elements = append(elements, []string{"2", "Days: " + server.Days + "\tTimerange: " + server.Timerange})
			}
		}

		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "timePeriod: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the TimePeriods to csv
func (s DetailServer) StringCSV() string {
	var p []DetailTimePeriod
	if s.Server.TimePeriods != nil {
		p = append(p, *s.Server.TimePeriods)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
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
