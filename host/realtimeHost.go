package host

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//RealtimeHost represents the caracteristics of a host
type RealtimeHost struct {
	ID           string `json:"id"`            //Host ID
	Name         string `json:"name"`          //Host name
	Alias        string `json:"alias"`         //Host alias
	Address      string `json:"address"`       //Host address
	State        string `json:"state"`         //State of the host
	Acknowledged string `json:"acknowledged"`  //If the host is acknowledge or not
	Activate     string `json:"active_checks"` //If the host is activate or not
	PollerName   string `json:"instance_name"` //Poller name of the host
}

//RealtimeServer represents a server with informations
type RealtimeServer struct {
	Server RealtimeInformations `json:"server"`
}

//RealtimeInformations represents the informations of the server
type RealtimeInformations struct {
	Name  string         `json:"name"`
	Hosts []RealtimeHost `json:"hosts"`
}

//StringText permits to display the caracteristics of the hosts to text
func (s RealtimeServer) StringText() string {
	var values string = "Host list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Hosts); i++ {
		values += "ID: " + s.Server.Hosts[i].ID + "\t"
		values += "Name: " + s.Server.Hosts[i].Name + "\t"
		values += "Alias: " + s.Server.Hosts[i].Alias + "\t"
		values += "IP address: " + s.Server.Hosts[i].Address + "\t"
		values += "State: " + GetState(s.Server.Hosts[i].State) + "\t"
		values += "Acknowledged: " + GetAcknowledgment(s.Server.Hosts[i].Acknowledged) + "\t"
		values += "Activate: " + s.Server.Hosts[i].Activate + "\t"
		values += "Poller name: " + s.Server.Hosts[i].PollerName + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the hosts to csv
func (s RealtimeServer) StringCSV() string {
	var values string = "Server,ID,Name,Alias,IPAddress,State,Acknowledged,Activate,PollerName\n"
	for i := 0; i < len(s.Server.Hosts); i++ {
		values += s.Server.Name + "," + s.Server.Hosts[i].ID + "," + s.Server.Hosts[i].Name + "," + s.Server.Hosts[i].Alias + "," + s.Server.Hosts[i].Address + "," + GetState(s.Server.Hosts[i].State) + "," + GetAcknowledgment(s.Server.Hosts[i].Acknowledged) + "," + s.Server.Hosts[i].Activate + "," + s.Server.Hosts[i].PollerName + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the hosts to json
func (s RealtimeServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the hosts to yaml
func (s RealtimeServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}

//GetState permits to obtain the value of the state
func GetState(stateValue string) string {
	state := ""
	switch stateValue {
	case "0":
		state = "UP"
	case "1":
		state = "DOWN"
	case "2":
		state = "UNREACHABLE"
	case "3":
		state = "PENDING"
	}
	return state
}

//GetAcknowledgment permits to obtain the value of the acknowledgement
func GetAcknowledgment(acknowledgeValue string) string {
	acknowledge := ""
	switch acknowledgeValue {
	case "0":
		acknowledge = "no"
	case "1":
		acknowledge = "yes"
	}
	return acknowledge
}
