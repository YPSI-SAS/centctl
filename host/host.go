package host

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//Host represents the caracteristics of a host
type Host struct {
	ID       string `json:"id"`       //Host ID
	Name     string `json:"name"`     //Host name
	Alias    string `json:"alias"`    //Host alias
	Address  string `json:"address"`  //Host address
	Activate string `json:"activate"` //If the host is activate or not
}

//Result represents a poller array
type Result struct {
	Hosts []Host `json:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name  string `json:"name"`
	Hosts []Host `json:"hosts"`
}

//StringText permits to display the caracteristics of the hosts to text
func (s Server) StringText() string {
	var values string = "Host list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Hosts); i++ {
		values += "ID: " + s.Server.Hosts[i].ID + "\t"
		values += "Name: " + s.Server.Hosts[i].Name + "\t"
		values += "Alias: " + s.Server.Hosts[i].Alias + "\t"
		values += "IP address: " + s.Server.Hosts[i].Address + "\t"
		values += "Activate: " + s.Server.Hosts[i].Activate + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the hosts to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Name,Alias,IPAddress,Activate\n"
	for i := 0; i < len(s.Server.Hosts); i++ {
		values += s.Server.Name + "," + s.Server.Hosts[i].ID + "," + s.Server.Hosts[i].Name + "," + s.Server.Hosts[i].Alias + "," + s.Server.Hosts[i].Address + "," + s.Server.Hosts[i].Activate + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the hosts to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the hosts to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
