package host

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//DetailHost represents the caracteristics of a host
type DetailHost struct {
	ID                  string `json:"id"`                     //Host ID
	Name                string `json:"name"`                   //Host name
	Alias               string `json:"alias"`                  //Host alias
	Address             string `json:"address"`                //Host address
	State               string `json:"state"`                  //State of the host
	StateType           string `json:"state_type"`             //State type of the host
	Output              string `json:"output"`                 //Host output
	MaxCheckAttempts    string `json:"max_check_attempts"`     //Maximum check attempts of the host
	CheckAttempt        string `json:"check_attempt"`          //Check attempt of the host
	LastCheck           string `json:"last_check"`             //Last check of the host
	LastStateChange     string `json:"last_state_change"`      //Last state change of the host
	LastHardStateChange string `json:"last_hard_state_change"` //Last hard state change of the host
	Acknowledged        string `json:"acknowledged"`           //If the host is acknowledge or not
	Activate            string `json:"active_checks"`          //If the host is activate or not
	PollerName          string `json:"instance_name"`          //Poller name of the host
	Criticality         string `json:"criticality"`            //Criticality of the host
	PassiveChecks       string `json:"passive_checks"`         //Accept passive results
	Notify              string `json:"notify"`                 //notification is enabled
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name string     `json:"name"`
	Host DetailHost `json:"host"`
}

//StringText permits to display the caracteristics of the hosts to text
func (s DetailServer) StringText() string {
	var values string = "Host detail for server " + s.Server.Name + ": \n"

	values += "ID: " + s.Server.Host.ID + "\t"
	values += "Name: " + s.Server.Host.Name + "\t"
	values += "Alias: " + s.Server.Host.Alias + "\t"
	values += "IP address: " + s.Server.Host.Address + "\t"
	values += "State: " + GetState(s.Server.Host.State) + "\t"
	values += "State type: " + getStateType(s.Server.Host.StateType) + "\t"
	values += "Output: " + s.Server.Host.Output + "\t"
	values += "Max check attempts: " + s.Server.Host.MaxCheckAttempts + "\t"
	values += "Check attempt: " + s.Server.Host.CheckAttempt + "\t"
	values += "Last check: " + s.Server.Host.LastCheck + "\t"
	values += "Last state change: " + s.Server.Host.LastStateChange + "\t"
	values += "Last hard state change: " + s.Server.Host.LastHardStateChange + "\t"
	values += "Acknowledged: " + GetAcknowledgment(s.Server.Host.Acknowledged) + "\t"
	values += "Activate: " + s.Server.Host.Activate + "\t"
	values += "Poller name: " + s.Server.Host.PollerName + "\t"
	values += "Criticality: " + s.Server.Host.Criticality + "\t"
	values += "Passive checks: " + s.Server.Host.PassiveChecks + "\t"
	values += "Notify: " + s.Server.Host.Notify + "\n"

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the hosts to csv
func (s DetailServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,IPAddress,State,StateType,Output,MaxCheckAttempts,CheckAttempt,LastCheck,LastStateChange,LastHardStateChange,Acknowledged,Activate,PollerName,Criticality,PassiveChecks,Notify\n"
	values += s.Server.Name + ","
	values += s.Server.Host.ID + ","
	values += s.Server.Host.Name + ","
	values += s.Server.Host.Alias + ","
	values += s.Server.Host.Address + ","
	values += GetState(s.Server.Host.State) + ","
	values += getStateType(s.Server.Host.StateType) + ","
	values += s.Server.Host.Output + ","
	values += s.Server.Host.MaxCheckAttempts + ","
	values += s.Server.Host.CheckAttempt + ","
	values += s.Server.Host.LastCheck + ","
	values += s.Server.Host.LastStateChange + ","
	values += s.Server.Host.LastHardStateChange + ","
	values += GetAcknowledgment(s.Server.Host.Acknowledged) + ","
	values += s.Server.Host.Activate + ","
	values += s.Server.Host.PollerName + ","
	values += s.Server.Host.Criticality + ","
	values += s.Server.Host.PassiveChecks + ","
	values += s.Server.Host.Notify + "\n"
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the hosts to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the hosts to yaml
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
