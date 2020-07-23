package contact

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

//Contact represents the caracteristics of a contact
type Contact struct {
	ID    string `json:"id"`    //Contact ID
	Name  string `json:"name"`  //Contact Name
	Alias string `json:"alias"` //Contact Alias
	Email string `json:"email"` //Contact Email
}

//Result represents a contact array send by the API
type Result struct {
	Contacts []Contact `json:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name     string    `json:"name"`
	Contacts []Contact `json:"contacts"`
}

//StringText permits to display the caracteristics of the contacts to text
func (s Server) StringText() string {
	var values string = "Contact list for server " + s.Server.Name + ": \n"
	for i := 0; i < len(s.Server.Contacts); i++ {
		values += "ID: " + s.Server.Contacts[i].ID + "\t"
		values += "Name: " + s.Server.Contacts[i].Name + "\t"
		values += "Alias: " + s.Server.Contacts[i].Alias + "\t"
		values += "Email: " + s.Server.Contacts[i].Email + "\n"
	}
	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the contacts to csv
func (s Server) StringCSV() string {
	var values string = "Server,ID,Name,Alias,Email\n"
	for i := 0; i < len(s.Server.Contacts); i++ {
		values += s.Server.Name + "," + s.Server.Contacts[i].ID + "," + s.Server.Contacts[i].Name + "," + s.Server.Contacts[i].Alias + "," + s.Server.Contacts[i].Email + "\n"
	}
	return fmt.Sprintf(values)
}

//StringJSON permits to display the caracteristics of the contacts to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the contacts to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
