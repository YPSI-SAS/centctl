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

//DetailRealtimeService represents the caracteristics of a service
type DetailRealtimeService struct {
	ID                                    int    `json:"id" yaml:"id"`                   //Service ID
	Description                           string `json:"description" yaml:"description"` //Service description
	State                                 int    `json:"state" yaml:"state"`             //State of the service
	DetailRealtimeServiceStatus           `json:"status" yaml:"status"`
	StateType                             int    `json:"state_type" yaml:"state_type"`                         //State type of the service
	Output                                string `json:"output" yaml:"output"`                                 //Service output
	MaxCheckAttempts                      int    `json:"max_check_attempts" yaml:"max_check_attempts"`         //Maximum check attempts of the service
	NextCheck                             string `json:"next_check" yaml:"next_check"`                         //Next check of the service
	LastUpdate                            string `json:"last_update" yaml:"last_update"`                       //Last update of the service
	LastCheck                             string `json:"last_check" yaml:"last_check"`                         //Last check of the service
	LastStateChange                       string `json:"last_state_change" yaml:"last_state_change"`           //Last state change of the service
	LastHardStateChange                   string `json:"last_hard_state_change" yaml:"last_hard_state_change"` //Last hard state change of the service
	Acknowledged                          bool   `json:"is_acknowledged" yaml:"is_acknowledged"`               //If the service is acknowledge or not
	Activate                              bool   `json:"is_active_checks" yaml:"is_active_checks"`             //If the service is activate or not
	Checked                               bool   `json:"is_checked" yaml:"is_checked"`
	ScheduledDowntimeDepth                int    `json:"scheduled_downtime_depth" yaml:"scheduled_downtime_depth"` //Schedule downtime depth of the service
	*DetailRealtimeServiceAcknowledgement `json:"acknowledgement" yaml:"acknowledgement"`
	Downtimes                             DetailRealtimeServiceDowntimes `json:"downtimes" yaml:"downtimes"`
}

type DetailRealtimeServiceDowntimes []DetailRealtimeServiceDowntime

func (t DetailRealtimeServiceDowntimes) MarshalCSV() ([]byte, error) {
	var value string
	for i, downtime := range t {
		value += strconv.Itoa(downtime.AuthorID) + "|" + downtime.AuthorName + "|" + downtime.Comment + "|" + strconv.Itoa(downtime.Duration) + "|" + downtime.EntryTime + "|" + downtime.StartTime + "|" + downtime.EndTime + "|" + strconv.FormatBool(downtime.Started) + "|" + strconv.FormatBool(downtime.Fixed)
		if i < len(t)-1 {
			value += ","
		}
	}
	return []byte(value), nil
}

type DetailRealtimeServiceStatus struct {
	Code         int    `json:"code" yaml:"code" csv:"StatusCode"`
	Name         string `json:"name" yaml:"name" csv:"StatusName"`
	SeverityCode int    `json:"severity_code" yaml:"severity_code" csv:"StatusSeverityCode"`
}

type DetailRealtimeServiceAcknowledgement struct {
	AuthorID          int    `json:"author_id" yaml:"author_id" csv:"AckAuthorID"`
	AuthorName        string `json:"author_name" yaml:"author_name" csv:"AckAuthorName"`
	Comment           string `json:"comment" yaml:"comment" csv:"AckComment"`
	EntryTime         string `json:"entry_time" yaml:"entry_time" csv:"AckEntryTime"`
	NotifyContact     bool   `json:"is_notify_contacts" yaml:"is_notify_contacts" csv:"AckNotifyContact"`
	PersistentComment bool   `json:"is_persistent_comment" yaml:"is_persistent_comment" csv:"AckPersistentComment"`
	Sticky            bool   `json:"is_sticky" yaml:"is_sticky" csv:"AckSticky"`
	HostID            int    `json:"host_id" yaml:"host_id" csv:"AckHostID"`
	PollerID          int    `json:"poller_id" yaml:"poller_id" csv:"AckPollerID"`
}

type DetailRealtimeServiceDowntime struct {
	AuthorID   int    `json:"author_id" yaml:"author_id"`
	AuthorName string `json:"author_name" yaml:"author_name"`
	HostID     int    `json:"host_id" yaml:"host_id"`
	Comment    string `json:"comment" yaml:"comment"`
	Duration   int    `json:"duration" yaml:"duration"`
	EntryTime  string `json:"entry_time" yaml:"entry_time"`
	StartTime  string `json:"start_time" yaml:"start_time"`
	EndTime    string `json:"end_time" yaml:"end_time"`
	Started    bool   `json:"is_started" yaml:"is_started"`
	Fixed      bool   `json:"is_fixed" yaml:"is_fixed"`
}

//DetailRealtimeServer represents a server with informations
type DetailRealtimeServer struct {
	Server DetailRealtimeInformations `json:"server" yaml:"server"`
}

//DetailRealtimeInformations represents the informations of the server
type DetailRealtimeInformations struct {
	Name    string                 `json:"name" yaml:"name"`
	Service *DetailRealtimeService `json:"service" yaml:"service"`
}

//StringText permits to display the caracteristics of the service to text
func (s DetailRealtimeServer) StringText() string {
	var values string
	service := s.Server.Service
	if service != nil {
		elements := [][]string{{"0", "Service:"}}
		elements = append(elements, []string{"1", "ID: " + strconv.Itoa((*service).ID)})
		elements = append(elements, []string{"1", "Description: " + (*service).Description})
		elements = append(elements, []string{"1", "State: " + strconv.Itoa((*service).State)})
		elements = append(elements, []string{"1", "State type: " + strconv.Itoa((*service).StateType)})
		elements = append(elements, []string{"1", "Status: " + (*service).DetailRealtimeServiceStatus.Name + "(Code: " + strconv.Itoa((*service).DetailRealtimeServiceStatus.Code) + ")"})
		elements = append(elements, []string{"1", "Output: " + (*service).Output})
		elements = append(elements, []string{"1", "Max check attempts: " + strconv.Itoa((*service).MaxCheckAttempts)})
		elements = append(elements, []string{"1", "Next check: " + (*service).NextCheck})
		elements = append(elements, []string{"1", "Last updata: " + (*service).LastUpdate})
		elements = append(elements, []string{"1", "Last check: " + (*service).LastCheck})
		elements = append(elements, []string{"1", "Last state change: " + (*service).LastStateChange})
		elements = append(elements, []string{"1", "Last hard state change: " + (*service).LastHardStateChange})
		elements = append(elements, []string{"1", "Acknowledged: " + strconv.FormatBool((*service).Acknowledged)})
		elements = append(elements, []string{"1", "Activate: " + strconv.FormatBool((*service).Activate)})
		elements = append(elements, []string{"1", "Checked: " + strconv.FormatBool((*service).Checked)})
		elements = append(elements, []string{"1", "Schedule downtime depth: " + strconv.Itoa((*service).ScheduledDowntimeDepth)})

		if (*service).DetailRealtimeServiceAcknowledgement != nil {
			elements = append(elements, []string{"1", "Acknowledgement:"})
			elements = append(elements, []string{"2", "Author: " + (*service).DetailRealtimeServiceAcknowledgement.AuthorName + " (ID: " + strconv.Itoa((*service).DetailRealtimeServiceAcknowledgement.AuthorID) + ")"})
			elements = append(elements, []string{"2", "Comment: " + (*service).DetailRealtimeServiceAcknowledgement.Comment})
			elements = append(elements, []string{"2", "Entry time: " + (*service).DetailRealtimeServiceAcknowledgement.EntryTime})
			elements = append(elements, []string{"2", "Notify contact: " + strconv.FormatBool((*service).DetailRealtimeServiceAcknowledgement.NotifyContact)})
			elements = append(elements, []string{"2", "Persistent Comment: " + strconv.FormatBool((*service).DetailRealtimeServiceAcknowledgement.PersistentComment)})
			elements = append(elements, []string{"2", "Sticky: " + strconv.FormatBool((*service).DetailRealtimeServiceAcknowledgement.Sticky)})
			elements = append(elements, []string{"2", "Host ID: " + strconv.Itoa((*service).DetailRealtimeServiceAcknowledgement.HostID)})
			elements = append(elements, []string{"2", "Poller ID: " + strconv.Itoa((*service).DetailRealtimeServiceAcknowledgement.PollerID)})
		} else {
			elements = append(elements, []string{"2", "Acknowledgement:[]"})
		}

		if len((*service).Downtimes) == 0 {
			elements = append(elements, []string{"1", "Downtimes: []"})
		} else {
			elements = append(elements, []string{"1", "Downtimes:"})
			for _, downtime := range (*service).Downtimes {
				elements = append(elements, []string{"2", "Author: " + downtime.AuthorName + " (ID: " + strconv.Itoa(downtime.AuthorID) + ")"})
				elements = append(elements, []string{"2", "Host ID: " + strconv.Itoa(downtime.HostID)})
				elements = append(elements, []string{"3", "Comment: " + downtime.Comment})
				elements = append(elements, []string{"3", "Duration: " + strconv.Itoa(downtime.Duration)})
				elements = append(elements, []string{"3", "Entry time: " + downtime.EntryTime})
				elements = append(elements, []string{"3", "Start time: " + downtime.StartTime})
				elements = append(elements, []string{"3", "End time: " + downtime.EndTime})
				elements = append(elements, []string{"3", "Started: " + strconv.FormatBool(downtime.Started)})
				elements = append(elements, []string{"3", "Fixed: " + strconv.FormatBool(downtime.Fixed)})

			}
		}

		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "service: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the service to csv
func (s DetailRealtimeServer) StringCSV() string {
	var p []DetailRealtimeService
	if s.Server.Service != nil {
		p = append(p, *s.Server.Service)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the service to json
func (s DetailRealtimeServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the service to yaml
func (s DetailRealtimeServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
