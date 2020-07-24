/*
MIT License

Copyright (c) 2020 YPSI SAS
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

	"gopkg.in/yaml.v2"
)

//DetailService represents the caracteristics of a service
type DetailService struct {
	ServiceID              string `json:"service_id"`               //Service ID
	Description            string `json:"description"`              //Service description
	HostID                 string `json:"host_id"`                  //Host ID of the service
	HostName               string `json:"host_name"`                //Host name of the service
	State                  string `json:"state"`                    //State of the service
	StateType              string `json:"state_type"`               //State type of the service
	Output                 string `json:"output"`                   //Service output
	Perfdata               string `json:"perfdata"`                 //Service perfdata
	MaxCheckAttempts       string `json:"max_check_attempts"`       //Maximum check attempts of the service
	CurrentAttempt         string `json:"current_attempt"`          //Current attempts of the service
	NextCheck              string `json:"next_check"`               //Next check of the service
	LastUpdate             string `json:"last_update"`              //Last update of the service
	LastCheck              string `json:"last_check"`               //Last check of the service
	LastStateChange        string `json:"last_state_change"`        //Last state change of the service
	LastHardStateChange    string `json:"last_hard_state_change"`   //Last hard state change of the service
	Acknowledged           string `json:"acknowledged"`             //If the service is acknowledge or not
	Activate               string `json:"active_checks"`            //If the service is activate or not
	PollerName             string `json:"instance_name"`            //Poller name who check this host
	Criticality            string `json:"criticality"`              //Criticality of the service
	PassiveChecks          string `json:"passive_checks"`           //Accept passive results
	Notify                 string `json:"notify"`                   //Notification is enabled
	ScheduledDowntimeDepth string `json:"scheduled_downtime_depth"` //Schedule downtime depth of the service
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name    string        `json:"name"`
	Service DetailService `json:"service"`
}

//StringText permits to display the caracteristics of the service to text
func (s DetailServer) StringText() string {
	var values string = "Service detail for server " + s.Server.Name + ": \n"

	values += "ID: " + s.Server.Service.ServiceID + "\t"
	values += "Description: " + s.Server.Service.Description + "\t"
	values += "Host ID: " + s.Server.Service.HostID + "\t"
	values += "Host name: " + s.Server.Service.HostName + "\t"
	values += "State: " + GetState(s.Server.Service.State) + "\t"
	values += "State type: " + getStateType(s.Server.Service.StateType) + "\t"
	values += "Output: " + s.Server.Service.Output + "\t"
	values += "Perfdata: " + s.Server.Service.Perfdata + "\t"
	values += "Max check attempts: " + s.Server.Service.MaxCheckAttempts + "\t"
	values += "Current attempt: " + s.Server.Service.CurrentAttempt + "\t"
	values += "Next check: " + s.Server.Service.NextCheck + "\t"
	values += "Last update: " + s.Server.Service.LastUpdate + "\t"
	values += "Last check: " + s.Server.Service.LastCheck + "\t"
	values += "Last state change: " + s.Server.Service.LastStateChange + "\t"
	values += "Last hard state change: " + s.Server.Service.LastHardStateChange + "\t"
	values += "Acknowledged: " + GetAcknowledgment(s.Server.Service.Acknowledged) + "\t"
	values += "Activate: " + s.Server.Service.Activate + "\t"
	values += "Poller name: " + s.Server.Service.PollerName + "\t"
	values += "Criticality: " + s.Server.Service.Criticality + "\t"
	values += "Passive checks: " + s.Server.Service.PassiveChecks + "\t"
	values += "Notify: " + s.Server.Service.Notify + "\t"
	values += "Scheduled downtime depth: " + s.Server.Service.ScheduledDowntimeDepth + "\n"

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the service to csv
func (s DetailServer) StringCSV() string {
	var values string = "Server,ID,Description,HostID,HostName,State,StateType,Output,Perfdata,MaxCheckAttempts,CheckAttempt,CurrentAttempt,NextCheck,LastUpdate,LastCheck,LastStateChange,LastHardStateChange,Acknowledged,Activate,PollerName,Criticality,PassiveChecks,Notify,ScheduledDowntimeDepth\n"
	values += s.Server.Name + ","
	values += s.Server.Service.ServiceID + ","
	values += s.Server.Service.Description + ","
	values += s.Server.Service.HostID + ","
	values += s.Server.Service.HostName + ","
	values += GetState(s.Server.Service.State) + ","
	values += getStateType(s.Server.Service.StateType) + ","
	values += s.Server.Service.Output + ","
	values += s.Server.Service.Perfdata + ","
	values += s.Server.Service.MaxCheckAttempts + ","
	values += s.Server.Service.CurrentAttempt + ","
	values += s.Server.Service.NextCheck + ","
	values += s.Server.Service.LastUpdate + ","
	values += s.Server.Service.LastCheck + ","
	values += s.Server.Service.LastStateChange + ","
	values += s.Server.Service.LastHardStateChange + ","
	values += GetAcknowledgment(s.Server.Service.Acknowledged) + ","
	values += s.Server.Service.Activate + ","
	values += s.Server.Service.PollerName + ","
	values += s.Server.Service.Criticality + ","
	values += s.Server.Service.PassiveChecks + ","
	values += s.Server.Service.Notify + ","
	values += s.Server.Service.ScheduledDowntimeDepth + "\n"
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the service to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the service to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}

//getStateType permits to obtain the value of the state type
func getStateType(stateTypeV string) string {
	state := ""
	switch stateTypeV {
	case "0":
		state = "SOFT"
	case "1":
		state = "HARD"
	}
	return state
}
