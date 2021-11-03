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
	"fmt"
	"strconv"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailTimelineService represents the caracteristics of a service
type DetailTimelineService struct {
	ID                            int                           `json:"id" yaml:"id"`                 //service ID
	Type                          string                        `json:"type" yaml:"type"`             //service name
	Date                          string                        `json:"date" yaml:"date" `            //service alias
	StartDate                     string                        `json:"start_date" yaml:"start_date"` //service address
	EndDate                       string                        `json:"end_date" yaml:"end_date"`     //service output
	Content                       string                        `json:"content" yaml:"content"`
	Tries                         int                           `json:"tries" yaml:"tries"` //Maximum check attempts of the service
	*DetailTimelineServiceStatus  `json:"status" yaml:"status"` //State of the service
	*DetailTimelineServiceContact `json:"contact" yaml:"contact"`
}

type DetailTimelineServiceContact struct {
	ID   int    `json:"id"  yaml:"id" csv:"ContactID"`
	Name string `json:"name" yaml:"name" csv:"ContactName"`
}

type DetailTimelineServiceStatus struct {
	Code         int    `json:"code" yaml:"code" csv:"StatusCode"`
	Name         string `json:"name" yaml:"name" csv:"StatusName"`
	SeverityCode int    `json:"severity_code" yaml:"severity_code" csv:"StatusSeverityCode"`
}

//TimelineServiceResult represents a service Group array
type TimelineServiceResult struct {
	DetailTimelineServices []DetailTimelineService `json:"result" yaml:"result"`
}

//DetailTimelineServer represents a server with informations
type DetailTimelineServer struct {
	Server DetailTimelineInformations `json:"server" yaml:"server"`
}

//DetailTimelineInformations represents the informations of the server
type DetailTimelineInformations struct {
	Name            string                  `json:"name" yaml:"name"`
	TimelineService []DetailTimelineService `json:"timeline_service" yaml:"timeline_service"`
}

//StringText permits to display the caracteristics of the hosts to text
func (s DetailTimelineServer) StringText() string {
	var values string
	elements := [][]string{{"0", "Timeline service:"}}
	for i := 0; i < len(s.Server.TimelineService); i++ {
		elements = append(elements, []string{"1", "ID: " + strconv.Itoa(s.Server.TimelineService[i].ID)})
		elements = append(elements, []string{"2", "Type: " + s.Server.TimelineService[i].Type})
		elements = append(elements, []string{"2", "Date: " + s.Server.TimelineService[i].Date})
		elements = append(elements, []string{"2", "Start date: " + s.Server.TimelineService[i].StartDate})
		elements = append(elements, []string{"2", "End date: " + s.Server.TimelineService[i].EndDate})
		elements = append(elements, []string{"2", "Content: " + s.Server.TimelineService[i].Content})
		elements = append(elements, []string{"2", "Tries: " + strconv.Itoa(s.Server.TimelineService[i].Tries)})
		if s.Server.TimelineService[i].DetailTimelineServiceContact != nil {
			elements = append(elements, []string{"2", "Contact:"})
			elements = append(elements, []string{"3", s.Server.TimelineService[i].DetailTimelineServiceContact.Name + " (ID: " + strconv.Itoa(s.Server.TimelineService[i].DetailTimelineServiceContact.ID) + ")"})
		} else {
			elements = append(elements, []string{"2", "Contact:[]"})
		}
		if s.Server.TimelineService[i].DetailTimelineServiceStatus != nil {
			elements = append(elements, []string{"2", "Status:"})
			elements = append(elements, []string{"3", s.Server.TimelineService[i].DetailTimelineServiceStatus.Name + " (ID: " + strconv.Itoa(s.Server.TimelineService[i].DetailTimelineServiceStatus.Code) + ")"})
		} else {
			elements = append(elements, []string{"2", "Status:[]"})
		}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the TimelineService to csv
func (s DetailTimelineServer) StringCSV() string {
	b, _ := csvutil.Marshal(s.Server.TimelineService)
	return string(b)
}

//StringJSON permits to display the caracteristics of the TimelineService to json
func (s DetailTimelineServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the hosts to yaml
func (s DetailTimelineServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
