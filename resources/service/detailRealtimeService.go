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
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

//DetailRealtimeService represents the caracteristics of a service
type DetailRealtimeService struct {
	ID                     int                                   `json:"id" yaml:"id"`                   //Service ID
	Description            string                                `json:"description" yaml:"description"` //Service description
	State                  int                                   `json:"state" yaml:"state"`             //State of the service
	Status                 DetailRealtimeServiceStatus           `json:"status" yaml:"status"`
	StateType              int                                   `json:"state_type" yaml:"state_type"`                         //State type of the service
	Output                 string                                `json:"output" yaml:"output"`                                 //Service output
	MaxCheckAttempts       int                                   `json:"max_check_attempts" yaml:"max_check_attempts"`         //Maximum check attempts of the service
	NextCheck              string                                `json:"next_check" yaml:"next_check"`                         //Next check of the service
	LastUpdate             string                                `json:"last_update" yaml:"last_update"`                       //Last update of the service
	LastCheck              string                                `json:"last_check" yaml:"last_check"`                         //Last check of the service
	LastStateChange        string                                `json:"last_state_change" yaml:"last_state_change"`           //Last state change of the service
	LastHardStateChange    string                                `json:"last_hard_state_change" yaml:"last_hard_state_change"` //Last hard state change of the service
	Acknowledged           bool                                  `json:"is_acknowledged" yaml:"is_acknowledged"`               //If the service is acknowledge or not
	Activate               bool                                  `json:"is_active_checks" yaml:"is_active_checks"`             //If the service is activate or not
	Checked                bool                                  `json:"is_checked" yaml:"is_checked"`
	ScheduledDowntimeDepth int                                   `json:"scheduled_downtime_depth" yaml:"scheduled_downtime_depth"` //Schedule downtime depth of the service
	Acknowledgement        *DetailRealtimeServiceAcknowledgement `json:"acknowledgement" yaml:"acknowledgement"`
	Downtimes              []DetailRealtimeServiceDowntime       `json:"downtimes" yaml:"downtimes"`
}

type DetailRealtimeServiceStatus struct {
	Code         int    `json:"code" yaml:"code"`
	Name         string `json:"name" yaml:"name"`
	SeverityCode int    `json:"severity_code" yaml:"severity_code"`
}

type DetailRealtimeServiceAcknowledgement struct {
	AuthorID          int    `json:"author_id" yaml:"author_id"`
	AuthorName        string `json:"author_name" yaml:"author_name"`
	Comment           string `json:"comment" yaml:"comment"`
	EntryTime         string `json:"entry_time" yaml:"entry_time"`
	NotifyContact     bool   `json:"is_notify_contacts" yaml:"is_notify_contacts"`
	PersistentComment bool   `json:"is_persistent_comment" yaml:"is_persistent_comment"`
	Sticky            bool   `json:"is_sticky" yaml:"is_sticky"`
	HostID            int    `json:"host_id" yaml:"host_id"`
	PollerID          int    `json:"poller_id" yaml:"poller_id"`
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
	var values string = "Service detail for server " + s.Server.Name + ": \n"
	service := s.Server.Service
	if service != nil {
		values += "ID: " + strconv.Itoa((*service).ID) + "\t"
		values += "Description: " + (*service).Description + "\t"
		values += "State: " + strconv.Itoa((*service).State) + "\t"
		values += "Status code: " + strconv.Itoa((*service).Status.Code) + "\t"
		values += "Status name: " + (*service).Status.Name + "\t"
		values += "State type: " + strconv.Itoa((*service).StateType) + "\t"
		values += "Output: " + (*service).Output + "\t"
		values += "Max check attempts: " + strconv.Itoa((*service).MaxCheckAttempts) + "\t"
		values += "Next check: " + (*service).NextCheck + "\t"
		values += "Last update: " + (*service).LastUpdate + "\t"
		values += "Last check: " + (*service).LastCheck + "\t"
		values += "Last state change: " + (*service).LastStateChange + "\t"
		values += "Last hard state change: " + (*service).LastHardStateChange + "\t"
		values += "Acknowledged: " + strconv.FormatBool((*service).Acknowledged) + "\t"
		values += "Activate: " + strconv.FormatBool((*service).Activate) + "\t"
		values += "Checked: " + strconv.FormatBool((*service).Checked) + "\t"
		values += "Scheduled downtime depth: " + strconv.Itoa((*service).ScheduledDowntimeDepth) + "\n"
	} else {
		values += "service: null\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the service to csv
func (s DetailRealtimeServer) StringCSV() string {
	var values string = "Server,ID,Description,State,StatusCode,StatusName,StateType,Output,MaxCheckAttempts,NextCheck,LastUpdate,LastCheck,LastStateChange,LastHardStateChange,Acknowledged,Activate,Checked,ScheduledDowntimeDepth\n"
	values += s.Server.Name + ","
	service := s.Server.Service
	if service != nil {
		values += strconv.Itoa((*service).ID) + ","
		values += (*service).Description + ","
		values += strconv.Itoa((*service).State) + ","
		values += strconv.Itoa((*service).Status.Code) + ","
		values += (*service).Status.Name + ","
		values += strconv.Itoa((*service).StateType) + ","
		values += (*service).Output + ","
		values += strconv.Itoa((*service).MaxCheckAttempts) + ","
		values += (*service).NextCheck + ","
		values += (*service).LastUpdate + ","
		values += (*service).LastCheck + ","
		values += (*service).LastStateChange + ","
		values += (*service).LastHardStateChange + ","
		values += strconv.FormatBool((*service).Acknowledged) + ","
		values += strconv.FormatBool((*service).Activate) + ","
		values += strconv.FormatBool((*service).Checked) + ","
		values += strconv.Itoa((*service).ScheduledDowntimeDepth) + "\n"
	} else {
		values += ",,,,,,,,,,,,,,,,\n"
	}
	return fmt.Sprintf(values)
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
