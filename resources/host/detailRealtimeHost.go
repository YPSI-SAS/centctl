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

package host

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"
)

//DetailRealtimeHost represents the caracteristics of a host
type DetailRealtimeHost struct {
	ID                  int                                `json:"id" yaml:"id"`                 //Host ID
	Name                string                             `json:"name" yaml:"name"`             //Host name
	Alias               string                             `json:"alias" yaml:"alias" `          //Host alias
	Address             string                             `json:"address_ip" yaml:"address_ip"` //Host address
	State               int                                `json:"state" yaml:"state"`           //Status of the host
	StateType           int                                `json:"state_type" yaml:"state_type"` //State type of the host
	Output              string                             `json:"output" yaml:"output"`         //Host output
	CheckCommand        string                             `json:"check_command" yaml:"check_command"`
	MaxCheckAttempts    int                                `json:"max_check_attempts" yaml:"max_check_attempts"`         //Maximum check attempts of the host
	CheckAttempt        int                                `json:"check_attempt" yaml:"check_attempt"`                   //Check attempt of the host
	LastCheck           string                             `json:"last_check" yaml:"last_check"`                         //Last check of the host
	LastStateChange     string                             `json:"last_state_change" yaml:"last_state_change"`           //Last state change of the host
	LastHardStateChange string                             `json:"last_hard_state_change" yaml:"last_hard_state_change"` //Last hard state change of the host
	Acknowledged        bool                               `json:"acknowledged" yaml:"acknowledged"`                     //If the host is acknowledge or not
	Activate            bool                               `json:"active_checks" yaml:"active_checks"`                   //If the host is activate or not
	PollerName          string                             `json:"poller_name" yaml:"poller_name"`                       //Poller name of the host
	PollerID            int                                `json:"poller_id" yaml:"poller_id"`                           //Poller ID of the host
	PassiveChecks       bool                               `json:"passive_checks" yaml:"passive_checks"`                 //Accept passive results
	Notify              bool                               `json:"notify" yaml:"notify"`                                 //notification is enabled
	Acknowledgement     *DetailRealtimeHostAcknowledgement `json:"acknowledgement" yaml:"acknowledgement"`
	Downtimes           []DetailRealtimeHostDowntime       `json:"downtimes" yaml:"downtimes"`
}

type DetailRealtimeHostAcknowledgement struct {
	AuthorID          int    `json:"author_id" yaml:"author_id"`
	AuthorName        string `json:"author_name" yaml:"author_name"`
	Comment           string `json:"comment" yaml:"comment"`
	EntryTime         string `json:"entry_time" yaml:"entry_time"`
	NotifyContact     bool   `json:"is_notify_contacts" yaml:"is_notify_contacts"`
	PersistentComment bool   `json:"is_persistent_comment" yaml:"is_persistent_comment"`
	Sticky            bool   `json:"is_sticky" yaml:"is_sticky"`
}

type DetailRealtimeHostDowntime struct {
	AuthorID   int    `json:"author_id" yaml:"author_id"`
	AuthorName string `json:"author_name" yaml:"author_name"`
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
	Name string              `json:"name" yaml:"name"`
	Host *DetailRealtimeHost `json:"host" yaml:"host"`
}

//StringText permits to display the caracteristics of the hosts to text
func (s DetailRealtimeServer) StringText() string {
	var values string = "Host detail for server " + s.Server.Name + ": \n"
	host := s.Server.Host
	if host != nil {
		values += "ID: " + strconv.Itoa((*host).ID) + "\t"
		values += "Name: " + (*host).Name + "\t"
		values += "Alias: " + (*host).Alias + "\t"
		values += "IP address: " + (*host).Address + "\t"
		values += "State: " + strconv.Itoa((*host).State) + "\t"
		values += "State type: " + strconv.Itoa((*host).StateType) + "\t"
		values += "Output: " + (*host).Output + "\t"
		values += "Check command: " + (*host).CheckCommand + "\t"
		values += "Max check attempts: " + strconv.Itoa((*host).MaxCheckAttempts) + "\t"
		values += "Check attempt: " + strconv.Itoa((*host).CheckAttempt) + "\t"
		values += "Last check: " + (*host).LastCheck + "\t"
		values += "Last state change: " + (*host).LastStateChange + "\t"
		values += "Last hard state change: " + (*host).LastHardStateChange + "\t"
		values += "Acknowledged: " + strconv.FormatBool((*host).Acknowledged) + "\t"
		values += "Activate: " + strconv.FormatBool((*host).Activate) + "\t"
		values += "Poller name: " + (*host).PollerName + "\t"
		values += "Poller id: " + strconv.Itoa((*host).PollerID) + "\t"
		values += "Passive checks: " + strconv.FormatBool((*host).PassiveChecks) + "\t"
		values += "Notify: " + strconv.FormatBool((*host).Notify) + "\n"
	} else {
		values += "Host: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the hosts to csv
func (s DetailRealtimeServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,IPAddress,State,StateType,Output,CheckCommand,MaxCheckAttempts,CheckAttempt,LastCheck,LastStateChange,LastHardStateChange,Acknowledged,Activate,PollerName,PollerID,PassiveChecks,Notify\n"
	values += s.Server.Name + ","
	host := s.Server.Host
	if host != nil {
		values += strconv.Itoa((*host).ID) + ","
		values += (*host).Name + ","
		values += (*host).Alias + ","
		values += (*host).Address + ","
		values += strconv.Itoa((*host).State) + ","
		values += strconv.Itoa((*host).StateType) + ","
		values += (*host).Output + ","
		values += (*host).CheckCommand + ","
		values += strconv.Itoa((*host).MaxCheckAttempts) + ","
		values += strconv.Itoa((*host).CheckAttempt) + ","
		values += (*host).LastCheck + ","
		values += (*host).LastStateChange + ","
		values += (*host).LastHardStateChange + ","
		values += strconv.FormatBool((*host).Acknowledged) + ","
		values += strconv.FormatBool((*host).Activate) + ","
		values += (*host).PollerName + ","
		values += strconv.Itoa((*host).PollerID) + ","
		values += strconv.FormatBool((*host).PassiveChecks) + ","
		values += strconv.FormatBool((*host).Notify) + "\n"

	} else {
		values += ",,,,,,,,,,,,,,,,,,\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the hosts to json
func (s DetailRealtimeServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the hosts to yaml
func (s DetailRealtimeServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
