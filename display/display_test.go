/*
MIT License

Copyright (c)  2020-2021 YPSI SAS


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

package display

import (
	"centctl/resources/contact"
	"centctl/resources/host"
	"centctl/resources/poller"
	"centctl/resources/service"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestDisplayContactJSON(t *testing.T) {
	contact1 := contact.Contact{
		ID:    "10",
		Name:  "nameContact",
		Alias: "aliasContact",
		Email: "contact@mail.com",
	}
	contacts := []contact.Contact{}
	contacts = append(contacts, contact1)
	server := contact.Server{
		Server: contact.Informations{
			Name:     "serverTEST",
			Contacts: contacts,
		},
	}
	displayContact, err := Contact("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayContact)
}

func TestDisplayContactYAML(t *testing.T) {
	contact1 := contact.Contact{
		ID:    "10",
		Name:  "nameContact",
		Alias: "aliasContact",
		Email: "contact@mail.com",
	}
	contacts := []contact.Contact{}
	contacts = append(contacts, contact1)
	server := contact.Server{
		Server: contact.Informations{
			Name:     "serverTEST",
			Contacts: contacts,
		},
	}
	displayContact, err := Contact("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayContact)
}

func TestDisplayContactCSV(t *testing.T) {
	contact1 := contact.Contact{
		ID:    "10",
		Name:  "nameContact",
		Alias: "aliasContact",
		Email: "contact@mail.com",
	}
	contacts := []contact.Contact{}
	contacts = append(contacts, contact1)
	server := contact.Server{
		Server: contact.Informations{
			Name:     "serverTEST",
			Contacts: contacts,
		},
	}
	displayContact, err := Contact("csv", server)
	expected := "Server,ID,Name,Alias,Email\n" + server.Server.Name + "," + contact1.ID + "," + contact1.Name + "," + contact1.Alias + "," + contact1.Email + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayContact)
}

func TestDisplayContactText(t *testing.T) {
	contact1 := contact.Contact{
		ID:    "10",
		Name:  "nameContact",
		Alias: "aliasContact",
		Email: "contact@mail.com",
	}
	contacts := []contact.Contact{}
	contacts = append(contacts, contact1)
	server := contact.Server{
		Server: contact.Informations{
			Name:     "serverTEST",
			Contacts: contacts,
		},
	}
	displayContact, err := Contact("text", server)
	expected := "Contact list for server " + server.Server.Name + ": \n"
	expected += "ID: " + contact1.ID + "\t"
	expected += "Name: " + contact1.Name + "\t"
	expected += "Alias: " + contact1.Alias + "\t"
	expected += "Email: " + contact1.Email + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayContact)
}

func TestDisplayContactIncorrectOutput(t *testing.T) {
	contact1 := contact.Contact{
		ID:    "10",
		Name:  "nameContact",
		Alias: "aliasContact",
		Email: "contact@mail.com",
	}
	contacts := []contact.Contact{}
	contacts = append(contacts, contact1)
	server := contact.Server{
		Server: contact.Informations{
			Name:     "serverTEST",
			Contacts: contacts,
		},
	}
	_, err := Contact("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}

func TestDisplayRealtimeServiceJSON(t *testing.T) {
	service1 := service.RealtimeService{
		ServiceID: 10,
		Name:      "realtime service test",
		Parent: service.RealtimeParent{
			ID:       12,
			Name:     "host",
			Address:  "127.0.0.1",
			PollerID: 1,
		},
		Status: service.RealtimeStatus{
			Code:         0,
			Name:         "OK",
			SeverityCode: 1,
		},
		Information:  "output warning",
		Acknowledged: false,
		ActiveCheck:  true,
	}
	services := []service.RealtimeService{}
	services = append(services, service1)
	server := service.RealtimeServer{
		Server: service.RealtimeInformations{
			Name:     "serverTEST",
			Services: services,
		},
	}
	displayService, err := RealtimeService("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayRealtimeServiceYAML(t *testing.T) {
	service1 := service.RealtimeService{
		ServiceID: 10,
		Name:      "realtime service test",
		Parent: service.RealtimeParent{
			ID:       12,
			Name:     "host",
			Address:  "127.0.0.1",
			PollerID: 1,
		},
		Status: service.RealtimeStatus{
			Code:         0,
			Name:         "OK",
			SeverityCode: 1,
		},
		Information:  "output warning",
		Acknowledged: false,
		ActiveCheck:  true,
	}
	services := []service.RealtimeService{}
	services = append(services, service1)
	server := service.RealtimeServer{
		Server: service.RealtimeInformations{
			Name:     "serverTEST",
			Services: services,
		},
	}
	displayService, err := RealtimeService("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayRealtimeServiceCSV(t *testing.T) {
	service1 := service.RealtimeService{
		ServiceID: 10,
		Name:      "realtime service test",
		Parent: service.RealtimeParent{
			ID:       12,
			Name:     "host",
			Address:  "127.0.0.1",
			PollerID: 1,
		},
		Status: service.RealtimeStatus{
			Code:         0,
			Name:         "OK",
			SeverityCode: 1,
		},
		Information:  "output warning",
		Acknowledged: false,
		ActiveCheck:  true,
	}
	services := []service.RealtimeService{}
	services = append(services, service1)
	server := service.RealtimeServer{
		Server: service.RealtimeInformations{
			Name:     "serverTEST",
			Services: services,
		},
	}
	displayService, err := RealtimeService("csv", server)
	expected := "Server,ID,Name,ParentID,ParentName,ParentPollerID,ParentAddress,StatusCode,StatusName,Information,Acknowledged,Activate\n"
	expected += server.Server.Name + "," + strconv.Itoa(service1.ServiceID) + "," + service1.Name + "," + strconv.Itoa(service1.Parent.PollerID) + "," + strconv.Itoa(service1.Parent.ID) + "," + service1.Parent.Name + "," + service1.Parent.Address + "," + strconv.Itoa(service1.Status.Code) + "," + service1.Status.Name + "," + service1.Information + "," + strconv.FormatBool(service1.Acknowledged) + "," + strconv.FormatBool(service1.ActiveCheck) + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayService)
}

func TestDisplayRealtimeServiceText(t *testing.T) {
	service1 := service.RealtimeService{
		ServiceID: 10,
		Name:      "realtime service test",
		Parent: service.RealtimeParent{
			ID:       12,
			Name:     "host",
			Address:  "127.0.0.1",
			PollerID: 1,
		},
		Status: service.RealtimeStatus{
			Code:         0,
			Name:         "OK",
			SeverityCode: 1,
		},
		Information:  "output warning",
		Acknowledged: false,
		ActiveCheck:  true,
	}
	services := []service.RealtimeService{}
	services = append(services, service1)
	server := service.RealtimeServer{
		Server: service.RealtimeInformations{
			Name:     "serverTEST",
			Services: services,
		},
	}
	displayService, err := RealtimeService("text", server)
	expected := "Service list for server=" + server.Server.Name + ": \n"
	expected += "ID: " + strconv.Itoa(service1.ServiceID) + "\t"
	expected += "Name: " + service1.Name + "\t"
	expected += "Parent ID: " + strconv.Itoa(service1.Parent.ID) + "\t"
	expected += "Parent name: " + service1.Parent.Name + "\t"
	expected += "Parent address: " + service1.Parent.Address + "\t"
	expected += "Parent pollerID: " + strconv.Itoa(service1.Parent.PollerID) + "\t"
	expected += "Status code: " + strconv.Itoa(service1.Status.Code) + "\t"
	expected += "Status name: " + service1.Status.Name + "\t"
	expected += "Information: " + service1.Information + "\t"
	expected += "Acknowledged: " + strconv.FormatBool(service1.Acknowledged) + "\t"
	expected += "ActiveCheck: " + strconv.FormatBool(service1.ActiveCheck) + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayService)
}

func TestDisplayRealtimeServiceIncorrectOutput(t *testing.T) {
	service1 := service.RealtimeService{
		ServiceID: 10,
		Name:      "realtime service test",
		Parent: service.RealtimeParent{
			ID:       12,
			Name:     "host",
			Address:  "127.0.0.1",
			PollerID: 1,
		},
		Status: service.RealtimeStatus{
			Code:         0,
			Name:         "OK",
			SeverityCode: 1,
		},
		Information:  "output warning",
		Acknowledged: false,
		ActiveCheck:  true,
	}
	services := []service.RealtimeService{}
	services = append(services, service1)
	server := service.RealtimeServer{
		Server: service.RealtimeInformations{
			Name:     "serverTEST",
			Services: services,
		},
	}
	_, err := RealtimeService("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}

func TestDisplayServiceJSON(t *testing.T) {
	service1 := service.Service{
		ServiceID:   "10",
		Description: "realtime service test",
		HostID:      "1",
		HostName:    "hostTEST",
		Activate:    "1",
	}
	services := []service.Service{}
	services = append(services, service1)
	server := service.Server{
		Server: service.Informations{
			Name:     "serverTEST",
			Services: services,
		},
	}
	displayService, err := Service("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayServiceYAML(t *testing.T) {
	service1 := service.Service{
		ServiceID:   "10",
		Description: "realtime service test",
		HostID:      "1",
		HostName:    "hostTEST",
		Activate:    "1",
	}
	services := []service.Service{}
	services = append(services, service1)
	server := service.Server{
		Server: service.Informations{
			Name:     "serverTEST",
			Services: services,
		},
	}
	displayService, err := Service("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayServiceCSV(t *testing.T) {
	service1 := service.Service{
		ServiceID:   "10",
		Description: "realtime service test",
		HostID:      "1",
		HostName:    "hostTEST",
		Activate:    "1",
	}
	services := []service.Service{}
	services = append(services, service1)
	server := service.Server{
		Server: service.Informations{
			Name:     "serverTEST",
			Services: services,
		},
	}
	displayService, err := Service("csv", server)
	expected := "Server,ID,Description,HostID,HostName,Activate\n"
	expected += server.Server.Name + "," + service1.ServiceID + "," + service1.Description + "," + service1.HostID + "," + service1.HostName + "," + service1.Activate + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayService)
}

func TestDisplayServiceText(t *testing.T) {
	service1 := service.Service{
		ServiceID:   "10",
		Description: "realtime service test",
		HostID:      "1",
		HostName:    "hostTEST",
		Activate:    "1",
	}
	services := []service.Service{}
	services = append(services, service1)
	server := service.Server{
		Server: service.Informations{
			Name:     "serverTEST",
			Services: services,
		},
	}
	displayService, err := Service("text", server)
	expected := "Service list for server" + server.Server.Name + ": \n"
	expected += "ID: " + service1.ServiceID + "\t"
	expected += "Description: " + service1.Description + "\t"
	expected += "Host ID: " + service1.HostID + "\t"
	expected += "Host name: " + service1.HostName + "\t"
	expected += "Activate: " + service1.Activate + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayService)
}

func TestDisplayServiceIncorrectOutput(t *testing.T) {
	service1 := service.Service{
		ServiceID:   "10",
		Description: "realtime service test",
		HostID:      "1",
		HostName:    "hostTEST",
		Activate:    "1",
	}
	services := []service.Service{}
	services = append(services, service1)
	server := service.Server{
		Server: service.Informations{
			Name:     "serverTEST",
			Services: services,
		},
	}
	_, err := Service("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}

func TestDisplayDetailServiceJSON(t *testing.T) {
	service1 := service.DetailRealtimeService{
		ID:          10,
		Description: "detail service test",
		State:       1,
		Status: service.DetailRealtimeServiceStatus{
			Code:         1,
			Name:         "critical",
			SeverityCode: 1,
		},
		StateType:              1,
		Output:                 "CRITICAL: Number of current processes running: 0\n",
		MaxCheckAttempts:       3,
		NextCheck:              "1589451778",
		LastUpdate:             "1589451480",
		LastCheck:              "1589451478",
		LastStateChange:        "1587653158",
		LastHardStateChange:    "1587653278",
		Acknowledged:           false,
		Activate:               true,
		Checked:                true,
		ScheduledDowntimeDepth: 1,
		Acknowledgement: &service.DetailRealtimeServiceAcknowledgement{
			AuthorID:          1,
			AuthorName:        "admin",
			Comment:           "ack by cli",
			EntryTime:         "2021-04-21T16:13:42+02:00",
			NotifyContact:     false,
			PersistentComment: true,
			Sticky:            false,
			HostID:            12,
			PollerID:          1,
		},
	}
	services := []service.DetailRealtimeService{}
	services = append(services, service1)
	server := service.DetailRealtimeServer{
		Server: service.DetailRealtimeInformations{
			Name:    "serverTEST",
			Service: &service1,
		},
	}
	displayService, err := DetailRealtimeService("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayDetailServiceYAML(t *testing.T) {
	service1 := service.DetailRealtimeService{
		ID:          10,
		Description: "detail service test",
		State:       1,
		Status: service.DetailRealtimeServiceStatus{
			Code:         1,
			Name:         "critical",
			SeverityCode: 1,
		},
		StateType:              1,
		Output:                 "CRITICAL: Number of current processes running: 0\n",
		MaxCheckAttempts:       3,
		NextCheck:              "1589451778",
		LastUpdate:             "1589451480",
		LastCheck:              "1589451478",
		LastStateChange:        "1587653158",
		LastHardStateChange:    "1587653278",
		Acknowledged:           false,
		Activate:               true,
		Checked:                true,
		ScheduledDowntimeDepth: 1,
		Acknowledgement: &service.DetailRealtimeServiceAcknowledgement{
			AuthorID:          1,
			AuthorName:        "admin",
			Comment:           "ack by cli",
			EntryTime:         "2021-04-21T16:13:42+02:00",
			NotifyContact:     false,
			PersistentComment: true,
			Sticky:            false,
			HostID:            12,
			PollerID:          1,
		},
	}
	services := []service.DetailRealtimeService{}
	services = append(services, service1)
	server := service.DetailRealtimeServer{
		Server: service.DetailRealtimeInformations{
			Name:    "serverTEST",
			Service: &service1,
		},
	}
	displayService, err := DetailRealtimeService("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayDetailServiceCSV(t *testing.T) {
	service1 := service.DetailRealtimeService{
		ID:          10,
		Description: "detail service test",
		State:       1,
		Status: service.DetailRealtimeServiceStatus{
			Code:         1,
			Name:         "critical",
			SeverityCode: 1,
		},
		StateType:              1,
		Output:                 "CRITICAL: Number of current processes running: 0\n",
		MaxCheckAttempts:       3,
		NextCheck:              "1589451778",
		LastUpdate:             "1589451480",
		LastCheck:              "1589451478",
		LastStateChange:        "1587653158",
		LastHardStateChange:    "1587653278",
		Acknowledged:           false,
		Activate:               true,
		Checked:                true,
		ScheduledDowntimeDepth: 1,
		Acknowledgement: &service.DetailRealtimeServiceAcknowledgement{
			AuthorID:          1,
			AuthorName:        "admin",
			Comment:           "ack by cli",
			EntryTime:         "2021-04-21T16:13:42+02:00",
			NotifyContact:     false,
			PersistentComment: true,
			Sticky:            false,
			HostID:            12,
			PollerID:          1,
		},
	}
	services := []service.DetailRealtimeService{}
	services = append(services, service1)
	server := service.DetailRealtimeServer{
		Server: service.DetailRealtimeInformations{
			Name:    "serverTEST",
			Service: &service1,
		},
	}
	displayService, err := DetailRealtimeService("csv", server)
	expected := "Server,ID,Description,State,StatusCode,StatusName,StateType,Output,MaxCheckAttempts,NextCheck,LastUpdate,LastCheck,LastStateChange,LastHardStateChange,Acknowledged,Activate,Checked,ScheduledDowntimeDepth\n"
	expected += server.Server.Name + ","
	expected += strconv.Itoa(service1.ID) + ","
	expected += service1.Description + ","
	expected += strconv.Itoa(service1.State) + ","
	expected += strconv.Itoa(service1.Status.Code) + ","
	expected += service1.Status.Name + ","
	expected += strconv.Itoa(service1.StateType) + ","
	expected += service1.Output + ","
	expected += strconv.Itoa(service1.MaxCheckAttempts) + ","
	expected += service1.NextCheck + ","
	expected += service1.LastUpdate + ","
	expected += service1.LastCheck + ","
	expected += service1.LastStateChange + ","
	expected += service1.LastHardStateChange + ","
	expected += strconv.FormatBool(service1.Acknowledged) + ","
	expected += strconv.FormatBool(service1.Activate) + ","
	expected += strconv.FormatBool(service1.Checked) + ","
	expected += strconv.Itoa(service1.ScheduledDowntimeDepth) + "\n"

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayDetailServiceText(t *testing.T) {
	service1 := service.DetailRealtimeService{
		ID:          10,
		Description: "detail service test",
		State:       1,
		Status: service.DetailRealtimeServiceStatus{
			Code:         1,
			Name:         "critical",
			SeverityCode: 1,
		},
		StateType:              1,
		Output:                 "CRITICAL: Number of current processes running: 0\n",
		MaxCheckAttempts:       3,
		NextCheck:              "1589451778",
		LastUpdate:             "1589451480",
		LastCheck:              "1589451478",
		LastStateChange:        "1587653158",
		LastHardStateChange:    "1587653278",
		Acknowledged:           false,
		Activate:               true,
		Checked:                true,
		ScheduledDowntimeDepth: 1,
		Acknowledgement: &service.DetailRealtimeServiceAcknowledgement{
			AuthorID:          1,
			AuthorName:        "admin",
			Comment:           "ack by cli",
			EntryTime:         "2021-04-21T16:13:42+02:00",
			NotifyContact:     false,
			PersistentComment: true,
			Sticky:            false,
			HostID:            12,
			PollerID:          1,
		},
	}
	services := []service.DetailRealtimeService{}
	services = append(services, service1)
	server := service.DetailRealtimeServer{
		Server: service.DetailRealtimeInformations{
			Name:    "serverTEST",
			Service: &service1,
		},
	}
	displayService, err := DetailRealtimeService("text", server)
	expected := "Service detail for server " + server.Server.Name + ": \n"

	expected += "ID: " + strconv.Itoa(service1.ID) + "\t"
	expected += "Description: " + service1.Description + "\t"
	expected += "State: " + strconv.Itoa(service1.State) + "\t"
	expected += "Status code: " + strconv.Itoa(service1.Status.Code) + "\t"
	expected += "Status name: " + service1.Status.Name + "\t"
	expected += "State type: " + strconv.Itoa(service1.StateType) + "\t"
	expected += "Output: " + service1.Output + "\t"
	expected += "Max check attempts: " + strconv.Itoa(service1.MaxCheckAttempts) + "\t"
	expected += "Next check: " + service1.NextCheck + "\t"
	expected += "Last update: " + service1.LastUpdate + "\t"
	expected += "Last check: " + service1.LastCheck + "\t"
	expected += "Last state change: " + service1.LastStateChange + "\t"
	expected += "Last hard state change: " + service1.LastHardStateChange + "\t"
	expected += "Acknowledged: " + strconv.FormatBool(service1.Acknowledged) + "\t"
	expected += "Activate: " + strconv.FormatBool(service1.Activate) + "\t"
	expected += "Checked: " + strconv.FormatBool(service1.Checked) + "\t"
	expected += "Scheduled downtime depth: " + strconv.Itoa(service1.ScheduledDowntimeDepth) + "\n"

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayDetailServiceIncorrectOutput(t *testing.T) {
	service1 := service.DetailRealtimeService{
		ID:          10,
		Description: "detail service test",
		State:       1,
		Status: service.DetailRealtimeServiceStatus{
			Code:         1,
			Name:         "critical",
			SeverityCode: 1,
		},
		StateType:              1,
		Output:                 "CRITICAL: Number of current processes running: 0\n",
		MaxCheckAttempts:       3,
		NextCheck:              "1589451778",
		LastUpdate:             "1589451480",
		LastCheck:              "1589451478",
		LastStateChange:        "1587653158",
		LastHardStateChange:    "1587653278",
		Acknowledged:           false,
		Activate:               true,
		Checked:                true,
		ScheduledDowntimeDepth: 1,
		Acknowledgement: &service.DetailRealtimeServiceAcknowledgement{
			AuthorID:          1,
			AuthorName:        "admin",
			Comment:           "ack by cli",
			EntryTime:         "2021-04-21T16:13:42+02:00",
			NotifyContact:     false,
			PersistentComment: true,
			Sticky:            false,
			HostID:            12,
			PollerID:          1,
		},
	}
	services := []service.DetailRealtimeService{}
	services = append(services, service1)
	server := service.DetailRealtimeServer{
		Server: service.DetailRealtimeInformations{
			Name:    "serverTEST",
			Service: &service1,
		},
	}
	_, err := DetailRealtimeService("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}

func TestDisplayRealtimeHostJSON(t *testing.T) {
	host1 := host.RealtimeHost{
		ID:           "1",
		Name:         "hostTEST",
		Alias:        "hostTEST",
		Address:      "127.0.0.1",
		State:        "1",
		Acknowledged: "0",
		Activate:     "1",
		PollerName:   "pollerTEST",
	}
	hosts := []host.RealtimeHost{}
	hosts = append(hosts, host1)
	server := host.RealtimeServer{
		Server: host.RealtimeInformations{
			Name:  "serverTEST",
			Hosts: hosts,
		},
	}
	displayHost, err := RealtimeHost("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayRealtimeHostYAML(t *testing.T) {
	host1 := host.RealtimeHost{
		ID:           "1",
		Name:         "hostTEST",
		Alias:        "hostTEST",
		Address:      "127.0.0.1",
		State:        "1",
		Acknowledged: "0",
		Activate:     "1",
		PollerName:   "pollerTEST",
	}
	hosts := []host.RealtimeHost{}
	hosts = append(hosts, host1)
	server := host.RealtimeServer{
		Server: host.RealtimeInformations{
			Name:  "serverTEST",
			Hosts: hosts,
		},
	}
	displayHost, err := RealtimeHost("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayRealtimeHostCSV(t *testing.T) {
	host1 := host.RealtimeHost{
		ID:           "1",
		Name:         "hostTEST",
		Alias:        "hostTEST",
		Address:      "127.0.0.1",
		State:        "1",
		Acknowledged: "0",
		Activate:     "1",
		PollerName:   "pollerTEST",
	}
	hosts := []host.RealtimeHost{}
	hosts = append(hosts, host1)
	server := host.RealtimeServer{
		Server: host.RealtimeInformations{
			Name:  "serverTEST",
			Hosts: hosts,
		},
	}
	displayHost, err := RealtimeHost("csv", server)
	expected := "Server,ID,Name,Alias,IPAddress,State,Acknowledged,Activate,PollerName\n"
	expected += server.Server.Name + "," + host1.ID + "," + host1.Name + "," + host1.Alias + "," + host1.Address + ",DOWN,no," + host1.Activate + "," + host1.PollerName + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayHost)
}

func TestDisplayRealtimeHostText(t *testing.T) {
	host1 := host.RealtimeHost{
		ID:           "1",
		Name:         "hostTEST",
		Alias:        "hostTEST",
		Address:      "127.0.0.1",
		State:        "1",
		Acknowledged: "0",
		Activate:     "1",
		PollerName:   "pollerTEST",
	}
	hosts := []host.RealtimeHost{}
	hosts = append(hosts, host1)
	server := host.RealtimeServer{
		Server: host.RealtimeInformations{
			Name:  "serverTEST",
			Hosts: hosts,
		},
	}
	displayHost, err := RealtimeHost("text", server)
	expected := "Host list for server " + server.Server.Name + ": \n"
	expected += "ID: " + host1.ID + "\t"
	expected += "Name: " + host1.Name + "\t"
	expected += "Alias: " + host1.Alias + "\t"
	expected += "IP address: " + host1.Address + "\t"
	expected += "State: DOWN\t"
	expected += "Acknowledged: no\t"
	expected += "Activate: " + host1.Activate + "\t"
	expected += "Poller name: " + host1.PollerName + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayHost)
}

func TestDisplayRealtimeHostIncorrectOutput(t *testing.T) {
	host1 := host.RealtimeHost{
		ID:           "1",
		Name:         "hostTEST",
		Alias:        "hostTEST",
		Address:      "127.0.0.1",
		State:        "1",
		Acknowledged: "0",
		Activate:     "1",
		PollerName:   "pollerTEST",
	}
	hosts := []host.RealtimeHost{}
	hosts = append(hosts, host1)
	server := host.RealtimeServer{
		Server: host.RealtimeInformations{
			Name:  "serverTEST",
			Hosts: hosts,
		},
	}
	_, err := RealtimeHost("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}

func TestDisplayHostJSON(t *testing.T) {
	host1 := host.Host{
		ID:       "1",
		Name:     "hostTEST",
		Alias:    "hostTEST",
		Address:  "127.0.0.1",
		Activate: "1",
	}
	hosts := []host.Host{}
	hosts = append(hosts, host1)
	server := host.Server{
		Server: host.Informations{
			Name:  "serverTEST",
			Hosts: hosts,
		},
	}
	displayHost, err := Host("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayHostYAML(t *testing.T) {
	host1 := host.Host{
		ID:       "1",
		Name:     "hostTEST",
		Alias:    "hostTEST",
		Address:  "127.0.0.1",
		Activate: "1",
	}
	hosts := []host.Host{}
	hosts = append(hosts, host1)
	server := host.Server{
		Server: host.Informations{
			Name:  "serverTEST",
			Hosts: hosts,
		},
	}
	displayHost, err := Host("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayHostCSV(t *testing.T) {
	host1 := host.Host{
		ID:       "1",
		Name:     "hostTEST",
		Alias:    "hostTEST",
		Address:  "127.0.0.1",
		Activate: "1",
	}
	hosts := []host.Host{}
	hosts = append(hosts, host1)
	server := host.Server{
		Server: host.Informations{
			Name:  "serverTEST",
			Hosts: hosts,
		},
	}
	displayHost, err := Host("csv", server)
	expected := "Server,ID,Name,Alias,IPAddress,Activate\n"
	expected += server.Server.Name + "," + host1.ID + "," + host1.Name + "," + host1.Alias + "," + host1.Address + "," + host1.Activate + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayHost)
}

func TestDisplayHostText(t *testing.T) {
	host1 := host.Host{
		ID:       "1",
		Name:     "hostTEST",
		Alias:    "hostTEST",
		Address:  "127.0.0.1",
		Activate: "1",
	}
	hosts := []host.Host{}
	hosts = append(hosts, host1)
	server := host.Server{
		Server: host.Informations{
			Name:  "serverTEST",
			Hosts: hosts,
		},
	}
	displayHost, err := Host("text", server)
	expected := "Host list for server " + server.Server.Name + ": \n"
	expected += "ID: " + host1.ID + "\t"
	expected += "Name: " + host1.Name + "\t"
	expected += "Alias: " + host1.Alias + "\t"
	expected += "IP address: " + host1.Address + "\t"
	expected += "Activate: " + host1.Activate + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayHost)
}

func TestDisplayHostIncorrectOutput(t *testing.T) {
	host1 := host.Host{
		ID:       "1",
		Name:     "hostTEST",
		Alias:    "hostTEST",
		Address:  "127.0.0.1",
		Activate: "1",
	}
	hosts := []host.Host{}
	hosts = append(hosts, host1)
	server := host.Server{
		Server: host.Informations{
			Name:  "serverTEST",
			Hosts: hosts,
		},
	}
	_, err := Host("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}

func TestDisplayDetailHostJSON(t *testing.T) {
	host1 := host.DetailRealtimeHost{
		ID:                  1,
		Name:                "hostTEST",
		Alias:               "hostTEST",
		Address:             "127.0.0.1",
		State:               1,
		StateType:           1,
		Output:              "UNKNOWN: Need to specify --status option.\n",
		CheckCommand:        "App-Monitoring-Centreon-Host-Dummy",
		MaxCheckAttempts:    3,
		CheckAttempt:        1,
		LastCheck:           "1589446838",
		LastStateChange:     "1589207728",
		LastHardStateChange: "1589207793",
		Acknowledged:        false,
		Activate:            true,
		PollerName:          "pollerTEST",
		PassiveChecks:       true,
		Notify:              true,
		Acknowledgement: &host.DetailRealtimeHostAcknowledgement{
			AuthorID:          1,
			AuthorName:        "admin",
			Comment:           "ack by cli",
			EntryTime:         "2021-04-21T16:13:42+02:00",
			NotifyContact:     false,
			PersistentComment: true,
			Sticky:            false,
		},
	}
	server := host.DetailRealtimeServer{
		Server: host.DetailRealtimeInformations{
			Name: "serverTEST",
			Host: &host1,
		},
	}
	displayHost, err := DetailRealtimeHost("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayDetailHostYAML(t *testing.T) {
	host1 := host.DetailRealtimeHost{
		ID:                  1,
		Name:                "hostTEST",
		Alias:               "hostTEST",
		Address:             "127.0.0.1",
		State:               1,
		StateType:           1,
		Output:              "UNKNOWN: Need to specify --status option.\n",
		CheckCommand:        "App-Monitoring-Centreon-Host-Dummy",
		MaxCheckAttempts:    3,
		CheckAttempt:        1,
		LastCheck:           "1589446838",
		LastStateChange:     "1589207728",
		LastHardStateChange: "1589207793",
		Acknowledged:        false,
		Activate:            true,
		PollerName:          "pollerTEST",
		PassiveChecks:       true,
		Notify:              true,
		Acknowledgement: &host.DetailRealtimeHostAcknowledgement{
			AuthorID:          1,
			AuthorName:        "admin",
			Comment:           "ack by cli",
			EntryTime:         "2021-04-21T16:13:42+02:00",
			NotifyContact:     false,
			PersistentComment: true,
			Sticky:            false,
		},
	}
	server := host.DetailRealtimeServer{
		Server: host.DetailRealtimeInformations{
			Name: "serverTEST",
			Host: &host1,
		},
	}
	displayHost, err := DetailRealtimeHost("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayDetailHostCSV(t *testing.T) {
	host1 := host.DetailRealtimeHost{
		ID:                  1,
		Name:                "hostTEST",
		Alias:               "hostTEST",
		Address:             "127.0.0.1",
		State:               1,
		StateType:           1,
		Output:              "UNKNOWN: Need to specify --status option.\n",
		CheckCommand:        "App-Monitoring-Centreon-Host-Dummy",
		MaxCheckAttempts:    3,
		CheckAttempt:        1,
		LastCheck:           "1589446838",
		LastStateChange:     "1589207728",
		LastHardStateChange: "1589207793",
		Acknowledged:        false,
		Activate:            true,
		PollerName:          "pollerTEST",
		PassiveChecks:       true,
		Notify:              true,
		Acknowledgement: &host.DetailRealtimeHostAcknowledgement{
			AuthorID:          1,
			AuthorName:        "admin",
			Comment:           "ack by cli",
			EntryTime:         "2021-04-21T16:13:42+02:00",
			NotifyContact:     false,
			PersistentComment: true,
			Sticky:            false,
		},
	}
	server := host.DetailRealtimeServer{
		Server: host.DetailRealtimeInformations{
			Name: "serverTEST",
			Host: &host1,
		},
	}
	displayHost, err := DetailRealtimeHost("csv", server)
	expected := "Server,ID,Name,Alias,IPAddress,State,StateType,Output,CheckCommand,MaxCheckAttempts,CheckAttempt,LastCheck,LastStateChange,LastHardStateChange,Acknowledged,Activate,PollerName,PollerID,PassiveChecks,Notify\n"
	expected += server.Server.Name + ","
	expected += strconv.Itoa(host1.ID) + ","
	expected += host1.Name + ","
	expected += host1.Alias + ","
	expected += host1.Address + ","
	expected += strconv.Itoa(host1.State) + ","
	expected += strconv.Itoa(host1.StateType) + ","
	expected += host1.Output + ","
	expected += host1.CheckCommand + ","
	expected += strconv.Itoa(host1.MaxCheckAttempts) + ","
	expected += strconv.Itoa(host1.CheckAttempt) + ","
	expected += host1.LastCheck + ","
	expected += host1.LastStateChange + ","
	expected += host1.LastHardStateChange + ","
	expected += strconv.FormatBool(host1.Acknowledged) + ","
	expected += strconv.FormatBool(host1.Activate) + ","
	expected += host1.PollerName + ","
	expected += strconv.Itoa(host1.PollerID) + ","
	expected += strconv.FormatBool(host1.PassiveChecks) + ","
	expected += strconv.FormatBool(host1.Notify) + "\n"

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayDetailHostText(t *testing.T) {
	host1 := host.DetailRealtimeHost{
		ID:                  1,
		Name:                "hostTEST",
		Alias:               "hostTEST",
		Address:             "127.0.0.1",
		State:               1,
		StateType:           1,
		Output:              "UNKNOWN: Need to specify --status option.\n",
		CheckCommand:        "App-Monitoring-Centreon-Host-Dummy",
		MaxCheckAttempts:    3,
		CheckAttempt:        1,
		LastCheck:           "1589446838",
		LastStateChange:     "1589207728",
		LastHardStateChange: "1589207793",
		Acknowledged:        false,
		Activate:            true,
		PollerName:          "pollerTEST",
		PassiveChecks:       true,
		Notify:              true,
		Acknowledgement: &host.DetailRealtimeHostAcknowledgement{
			AuthorID:          1,
			AuthorName:        "admin",
			Comment:           "ack by cli",
			EntryTime:         "2021-04-21T16:13:42+02:00",
			NotifyContact:     false,
			PersistentComment: true,
			Sticky:            false,
		},
	}
	server := host.DetailRealtimeServer{
		Server: host.DetailRealtimeInformations{
			Name: "serverTEST",
			Host: &host1,
		},
	}
	displayHost, err := DetailRealtimeHost("text", server)
	expected := "Host detail for server " + server.Server.Name + ": \n"
	expected += "ID: " + strconv.Itoa(host1.ID) + "\t"
	expected += "Name: " + host1.Name + "\t"
	expected += "Alias: " + host1.Alias + "\t"
	expected += "IP address: " + host1.Address + "\t"
	expected += "State: " + strconv.Itoa(host1.State) + "\t"
	expected += "State type: " + strconv.Itoa(host1.StateType) + "\t"
	expected += "Output: " + host1.Output + "\t"
	expected += "Check command: " + host1.CheckCommand + "\t"
	expected += "Max check attempts: " + strconv.Itoa(host1.MaxCheckAttempts) + "\t"
	expected += "Check attempt: " + strconv.Itoa(host1.CheckAttempt) + "\t"
	expected += "Last check: " + host1.LastCheck + "\t"
	expected += "Last state change: " + host1.LastStateChange + "\t"
	expected += "Last hard state change: " + host1.LastHardStateChange + "\t"
	expected += "Acknowledged: " + strconv.FormatBool(host1.Acknowledged) + "\t"
	expected += "Activate: " + strconv.FormatBool(host1.Activate) + "\t"
	expected += "Poller name: " + host1.PollerName + "\t"
	expected += "Poller id: " + strconv.Itoa(host1.PollerID) + "\t"
	expected += "Passive checks: " + strconv.FormatBool(host1.PassiveChecks) + "\t"
	expected += "Notify: " + strconv.FormatBool(host1.Notify) + "\t"

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayDetailHostIncorrectOutput(t *testing.T) {
	host1 := host.DetailRealtimeHost{
		ID:                  1,
		Name:                "hostTEST",
		Alias:               "hostTEST",
		Address:             "127.0.0.1",
		State:               1,
		StateType:           1,
		Output:              "UNKNOWN: Need to specify --status option.\n",
		CheckCommand:        "App-Monitoring-Centreon-Host-Dummy",
		MaxCheckAttempts:    3,
		CheckAttempt:        1,
		LastCheck:           "1589446838",
		LastStateChange:     "1589207728",
		LastHardStateChange: "1589207793",
		Acknowledged:        false,
		Activate:            true,
		PollerName:          "pollerTEST",
		PassiveChecks:       true,
		Notify:              true,
		Acknowledgement: &host.DetailRealtimeHostAcknowledgement{
			AuthorID:          1,
			AuthorName:        "admin",
			Comment:           "ack by cli",
			EntryTime:         "2021-04-21T16:13:42+02:00",
			NotifyContact:     false,
			PersistentComment: true,
			Sticky:            false,
		},
	}
	server := host.DetailRealtimeServer{
		Server: host.DetailRealtimeInformations{
			Name: "serverTEST",
			Host: &host1,
		},
	}
	_, err := DetailRealtimeHost("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}

func TestDisplayPollerJSON(t *testing.T) {
	poller1 := poller.Poller{
		Type:  "poller",
		Label: "PollerTEST",
		Metadata: poller.Metadata{
			CentreonID: "1",
			HostName:   "host",
			Address:    "127.0.0.1",
		},
	}
	pollers := []poller.Poller{}
	pollers = append(pollers, poller1)
	server := poller.Server{
		Server: poller.Informations{
			Name:    "serverTEST",
			Pollers: pollers,
		},
	}
	displayPoller, err := Poller("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayPoller)
}

func TestDisplayPollerYAML(t *testing.T) {
	poller1 := poller.Poller{
		Type:  "poller",
		Label: "PollerTEST",
		Metadata: poller.Metadata{
			CentreonID: "1",
			HostName:   "host",
			Address:    "127.0.0.1",
		},
	}
	pollers := []poller.Poller{}
	pollers = append(pollers, poller1)
	server := poller.Server{
		Server: poller.Informations{
			Name:    "serverTEST",
			Pollers: pollers,
		},
	}
	displayPoller, err := Poller("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayPoller)
}

func TestDisplayPollerCSV(t *testing.T) {
	poller1 := poller.Poller{
		Type:  "poller",
		Label: "PollerTEST",
		Metadata: poller.Metadata{
			CentreonID: "1",
			HostName:   "host",
			Address:    "127.0.0.1",
		},
	}
	pollers := []poller.Poller{}
	pollers = append(pollers, poller1)
	server := poller.Server{
		Server: poller.Informations{
			Name:    "serverTEST",
			Pollers: pollers,
		},
	}
	displayPoller, err := Poller("csv", server)
	expected := "Server,Type,Label,CentreonID,Hostname,Address\n"
	expected += server.Server.Name + "," + poller1.Type + "," + poller1.Label + "," + poller1.Metadata.CentreonID + "," + poller1.Metadata.HostName + "," + poller1.Metadata.Address + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayPoller)
}

func TestDisplayPollerText(t *testing.T) {
	poller1 := poller.Poller{
		Type:  "poller",
		Label: "PollerTEST",
		Metadata: poller.Metadata{
			CentreonID: "1",
			HostName:   "host",
			Address:    "127.0.0.1",
		},
	}
	pollers := []poller.Poller{}
	pollers = append(pollers, poller1)
	server := poller.Server{
		Server: poller.Informations{
			Name:    "serverTEST",
			Pollers: pollers,
		},
	}
	displayPoller, err := Poller("text", server)
	expected := "Poller list for server" + server.Server.Name + ": \n"
	expected += "Type: " + poller1.Type + "\t"
	expected += "Label: " + poller1.Label + "\t"
	expected += "CentreonID: " + poller1.Metadata.CentreonID + "\t"
	expected += "Hosname: " + poller1.Metadata.HostName + "\t"
	expected += "Address: " + poller1.Metadata.Address + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayPoller)
}

func TestDisplayPollerIncorrectOutput(t *testing.T) {
	poller1 := poller.Poller{
		Type:  "poller",
		Label: "PollerTEST",
		Metadata: poller.Metadata{
			CentreonID: "1",
			HostName:   "host",
			Address:    "127.0.0.1",
		},
	}
	pollers := []poller.Poller{}
	pollers = append(pollers, poller1)
	server := poller.Server{
		Server: poller.Informations{
			Name:    "serverTEST",
			Pollers: pollers,
		},
	}
	_, err := Poller("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}

func TestDisplayTemplateHostJSON(t *testing.T) {
	template1 := host.Template{
		ID:   "1",
		Name: "Template-Host-TEST",
	}
	templates := []host.Template{}
	templates = append(templates, template1)
	server := host.TemplateServer{
		Server: host.TemplateInformations{
			Name:      "serverTEST",
			Templates: templates,
		},
	}
	displayTemplateHost, err := TemplateHost("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayTemplateHost)
}

func TestDisplayTemplateHostYAML(t *testing.T) {
	template1 := host.Template{
		ID:   "1",
		Name: "Template-Host-TEST",
	}
	templates := []host.Template{}
	templates = append(templates, template1)
	server := host.TemplateServer{
		Server: host.TemplateInformations{
			Name:      "serverTEST",
			Templates: templates,
		},
	}
	displayTemplateHost, err := TemplateHost("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayTemplateHost)
}

func TestDisplayTemplateHostCSV(t *testing.T) {
	template1 := host.Template{
		ID:   "1",
		Name: "Template-Host-TEST",
	}
	templates := []host.Template{}
	templates = append(templates, template1)
	server := host.TemplateServer{
		Server: host.TemplateInformations{
			Name:      "serverTEST",
			Templates: templates,
		},
	}
	displayTemplateHost, err := TemplateHost("csv", server)
	expected := "Server,Name\n" + server.Server.Name + "," + template1.Name + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayTemplateHost)
}

func TestDisplayTemplateHostText(t *testing.T) {
	template1 := host.Template{
		ID:   "1",
		Name: "Template-Host-TEST",
	}
	templates := []host.Template{}
	templates = append(templates, template1)
	server := host.TemplateServer{
		Server: host.TemplateInformations{
			Name:      "serverTEST",
			Templates: templates,
		},
	}
	displayTemplateHost, err := TemplateHost("text", server)
	expected := "Host template list for server " + server.Server.Name + ": \n" + template1.Name + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayTemplateHost)
}

func TestDisplayTemplateHostInputIncorrect(t *testing.T) {
	template1 := host.Template{
		ID:   "1",
		Name: "Template-Host-TEST",
	}
	templates := []host.Template{}
	templates = append(templates, template1)
	server := host.TemplateServer{
		Server: host.TemplateInformations{
			Name:      "serverTEST",
			Templates: templates,
		},
	}
	_, err := TemplateHost("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}

func TestDisplayTemplateServiceJSON(t *testing.T) {
	template1 := service.Template{
		Description: "Template-Service-TEST",
	}
	templates := []service.Template{}
	templates = append(templates, template1)
	server := service.TemplateServer{
		Server: service.TemplateInformations{
			Name:      "serverTEST",
			Templates: templates,
		},
	}
	displayTemplateService, err := TemplateService("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayTemplateService)
}

func TestDisplayTemplateServiceYAML(t *testing.T) {
	template1 := service.Template{
		Description: "Template-Service-TEST",
	}
	templates := []service.Template{}
	templates = append(templates, template1)
	server := service.TemplateServer{
		Server: service.TemplateInformations{
			Name:      "serverTEST",
			Templates: templates,
		},
	}
	displayTemplateService, err := TemplateService("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayTemplateService)
}

func TestDisplayTemplateServiceCSV(t *testing.T) {
	template1 := service.Template{
		Description: "Template-Service-TEST",
	}
	templates := []service.Template{}
	templates = append(templates, template1)
	server := service.TemplateServer{
		Server: service.TemplateInformations{
			Name:      "serverTEST",
			Templates: templates,
		},
	}
	displayTemplateService, err := TemplateService("csv", server)
	expected := "Server,Description\n" + server.Server.Name + "," + template1.Description + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayTemplateService)
}

func TestDisplayTemplateServiceText(t *testing.T) {
	template1 := service.Template{
		Description: "Template-Service-TEST",
	}
	templates := []service.Template{}
	templates = append(templates, template1)
	server := service.TemplateServer{
		Server: service.TemplateInformations{
			Name:      "serverTEST",
			Templates: templates,
		},
	}
	displayTemplateService, err := TemplateService("text", server)
	expected := "Service template list for server " + server.Server.Name + ": \n" + template1.Description + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayTemplateService)
}

func TestDisplayTemplateServiceIncorrectOutput(t *testing.T) {
	template1 := service.Template{
		Description: "Template-Service-TEST",
	}
	templates := []service.Template{}
	templates = append(templates, template1)
	server := service.TemplateServer{
		Server: service.TemplateInformations{
			Name:      "serverTEST",
			Templates: templates,
		},
	}
	_, err := TemplateService("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}
