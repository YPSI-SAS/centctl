package display

import (
	"centctl/contact"
	"centctl/host"
	"centctl/poller"
	"centctl/service"
	"encoding/json"
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
		ServiceID:    "10",
		Description:  "realtime service test",
		HostID:       "1",
		HostName:     "hostTEST",
		State:        "1",
		Output:       "output warning",
		Acknowledged: "0",
		Activate:     "1",
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
		ServiceID:    "10",
		Description:  "realtime service test",
		HostID:       "1",
		HostName:     "hostTEST",
		State:        "1",
		Output:       "output warning",
		Acknowledged: "0",
		Activate:     "1",
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
		ServiceID:    "10",
		Description:  "realtime service test",
		HostID:       "1",
		HostName:     "hostTEST",
		State:        "1",
		Output:       "output warning",
		Acknowledged: "0",
		Activate:     "1",
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
	expected := "Server,ID,Description,HostID,HostName,State,Output,Acknowledged,Activate\n"
	expected += server.Server.Name + "," + service1.ServiceID + "," + service1.Description + "," + service1.HostID + "," + service1.HostName + ",Warning," + service1.Output + ",no," + service1.Activate + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayService)
}

func TestDisplayRealtimeServiceText(t *testing.T) {
	service1 := service.RealtimeService{
		ServiceID:    "10",
		Description:  "realtime service test",
		HostID:       "1",
		HostName:     "hostTEST",
		State:        "1",
		Output:       "output warning",
		Acknowledged: "0",
		Activate:     "1",
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
	expected += "ID: " + service1.ServiceID + "\t"
	expected += "Description: " + service1.Description + "\t"
	expected += "Host ID: " + service1.HostID + "\t"
	expected += "Host name: " + service1.HostName + "\t"
	expected += "State: Warning\t"
	expected += "Output: " + service1.Output + "\t"
	expected += "Acknowledged: no\t"
	expected += "Activate: " + service1.Activate + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayService)
}

func TestDisplayRealtimeServiceIncorrectOutput(t *testing.T) {
	service1 := service.RealtimeService{
		ServiceID:    "10",
		Description:  "realtime service test",
		HostID:       "1",
		HostName:     "hostTEST",
		State:        "1",
		Output:       "output warning",
		Acknowledged: "0",
		Activate:     "1",
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
	service1 := service.DetailService{
		ServiceID:              "10",
		Description:            "detail service test",
		HostID:                 "1",
		HostName:               "hostTEST",
		State:                  "2",
		StateType:              "1",
		Output:                 "CRITICAL: Number of current processes running: 0\n",
		Perfdata:               "'nbproc'=0;;1:1;0;",
		MaxCheckAttempts:       "3",
		CurrentAttempt:         "3",
		NextCheck:              "1589451778",
		LastUpdate:             "1589451480",
		LastCheck:              "1589451478",
		LastStateChange:        "1587653158",
		LastHardStateChange:    "1587653278",
		Acknowledged:           "0",
		Activate:               "1",
		PollerName:             "Poller",
		Criticality:            "",
		PassiveChecks:          "0",
		Notify:                 "1",
		ScheduledDowntimeDepth: "1",
	}
	services := []service.DetailService{}
	services = append(services, service1)
	server := service.DetailServer{
		Server: service.DetailInformations{
			Name:    "serverTEST",
			Service: service1,
		},
	}
	displayService, err := DetailService("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayDetailServiceYAML(t *testing.T) {
	service1 := service.DetailService{
		ServiceID:              "10",
		Description:            "detail service test",
		HostID:                 "1",
		HostName:               "hostTEST",
		State:                  "2",
		StateType:              "1",
		Output:                 "CRITICAL: Number of current processes running: 0\n",
		Perfdata:               "'nbproc'=0;;1:1;0;",
		MaxCheckAttempts:       "3",
		CurrentAttempt:         "3",
		NextCheck:              "1589451778",
		LastUpdate:             "1589451480",
		LastCheck:              "1589451478",
		LastStateChange:        "1587653158",
		LastHardStateChange:    "1587653278",
		Acknowledged:           "0",
		Activate:               "1",
		PollerName:             "Poller",
		Criticality:            "",
		PassiveChecks:          "0",
		Notify:                 "1",
		ScheduledDowntimeDepth: "1",
	}
	services := []service.DetailService{}
	services = append(services, service1)
	server := service.DetailServer{
		Server: service.DetailInformations{
			Name:    "serverTEST",
			Service: service1,
		},
	}
	displayService, err := DetailService("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayDetailServiceCSV(t *testing.T) {
	service1 := service.DetailService{
		ServiceID:              "10",
		Description:            "detail service test",
		HostID:                 "1",
		HostName:               "hostTEST",
		State:                  "2",
		StateType:              "1",
		Output:                 "CRITICAL: Number of current processes running: 0\n",
		Perfdata:               "'nbproc'=0;;1:1;0;",
		MaxCheckAttempts:       "3",
		CurrentAttempt:         "3",
		NextCheck:              "1589451778",
		LastUpdate:             "1589451480",
		LastCheck:              "1589451478",
		LastStateChange:        "1587653158",
		LastHardStateChange:    "1587653278",
		Acknowledged:           "0",
		Activate:               "1",
		PollerName:             "Poller",
		Criticality:            "",
		PassiveChecks:          "0",
		Notify:                 "1",
		ScheduledDowntimeDepth: "1",
	}
	services := []service.DetailService{}
	services = append(services, service1)
	server := service.DetailServer{
		Server: service.DetailInformations{
			Name:    "serverTEST",
			Service: service1,
		},
	}
	displayService, err := DetailService("csv", server)
	expected := "Server,ID,Description,HostID,HostName,State,StateType,Output,Perfdata,MaxCheckAttempts,CheckAttempt,CurrentAttempt,NextCheck,LastUpdate,LastCheck,LastStateChange,LastHardStateChange,Acknowledged,Activate,PollerName,Criticality,PassiveChecks,Notify,ScheduledDowntimeDepth\n"
	expected += server.Server.Name + ","
	expected += service1.ServiceID + ","
	expected += service1.Description + ","
	expected += service1.HostID + ","
	expected += service1.HostName + ","
	expected += "Critical,"
	expected += "HARD,"
	expected += service1.Output + ","
	expected += service1.Perfdata + ","
	expected += service1.MaxCheckAttempts + ","
	expected += service1.CurrentAttempt + ","
	expected += service1.NextCheck + ","
	expected += service1.LastUpdate + ","
	expected += service1.LastCheck + ","
	expected += service1.LastStateChange + ","
	expected += service1.LastHardStateChange + ","
	expected += "no,"
	expected += service1.Activate + ","
	expected += service1.PollerName + ","
	expected += service1.Criticality + ","
	expected += service1.PassiveChecks + ","
	expected += service1.Notify + ","
	expected += service1.ScheduledDowntimeDepth + "\n"

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayDetailServiceText(t *testing.T) {
	service1 := service.DetailService{
		ServiceID:              "10",
		Description:            "detail service test",
		HostID:                 "1",
		HostName:               "hostTEST",
		State:                  "2",
		StateType:              "1",
		Output:                 "CRITICAL: Number of current processes running: 0\n",
		Perfdata:               "'nbproc'=0;;1:1;0;",
		MaxCheckAttempts:       "3",
		CurrentAttempt:         "3",
		NextCheck:              "1589451778",
		LastUpdate:             "1589451480",
		LastCheck:              "1589451478",
		LastStateChange:        "1587653158",
		LastHardStateChange:    "1587653278",
		Acknowledged:           "0",
		Activate:               "1",
		PollerName:             "Poller",
		Criticality:            "",
		PassiveChecks:          "0",
		Notify:                 "1",
		ScheduledDowntimeDepth: "1",
	}
	services := []service.DetailService{}
	services = append(services, service1)
	server := service.DetailServer{
		Server: service.DetailInformations{
			Name:    "serverTEST",
			Service: service1,
		},
	}
	displayService, err := DetailService("text", server)
	expected := "Service detail for server " + server.Server.Name + ": \n"

	expected += "ID: " + service1.ServiceID + "\t"
	expected += "Description: " + service1.Description + "\t"
	expected += "Host ID: " + service1.HostID + "\t"
	expected += "Host name: " + service1.HostName + "\t"
	expected += "State: Critical\t"
	expected += "State type: HARD\t"
	expected += "Output: " + service1.Output + "\t"
	expected += "Perfdata: " + service1.Perfdata + "\t"
	expected += "Max check attempts: " + service1.MaxCheckAttempts + "\t"
	expected += "Current attempt: " + service1.CurrentAttempt + "\t"
	expected += "Next check: " + service1.NextCheck + "\t"
	expected += "Last update: " + service1.LastUpdate + "\t"
	expected += "Last check: " + service1.LastCheck + "\t"
	expected += "Last state change: " + service1.LastStateChange + "\t"
	expected += "Last hard state change: " + service1.LastHardStateChange + "\t"
	expected += "Acknowledged: no\t"
	expected += "Activate: " + service1.Activate + "\t"
	expected += "Poller name: " + service1.PollerName + "\t"
	expected += "Criticality: " + service1.Criticality + "\t"
	expected += "Passive checks: " + service1.PassiveChecks + "\t"
	expected += "Notify: " + service1.Notify + "\t"
	expected += "Scheduled downtime depth: " + service1.ScheduledDowntimeDepth + "\n"

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayService)
}

func TestDisplayDetailServiceIncorrectOutput(t *testing.T) {
	service1 := service.DetailService{
		ServiceID:              "10",
		Description:            "detail service test",
		HostID:                 "1",
		HostName:               "hostTEST",
		State:                  "2",
		StateType:              "1",
		Output:                 "CRITICAL: Number of current processes running: 0\n",
		Perfdata:               "'nbproc'=0;;1:1;0;",
		MaxCheckAttempts:       "3",
		CurrentAttempt:         "3",
		NextCheck:              "1589451778",
		LastUpdate:             "1589451480",
		LastCheck:              "1589451478",
		LastStateChange:        "1587653158",
		LastHardStateChange:    "1587653278",
		Acknowledged:           "0",
		Activate:               "1",
		PollerName:             "Poller",
		Criticality:            "",
		PassiveChecks:          "0",
		Notify:                 "1",
		ScheduledDowntimeDepth: "1",
	}
	services := []service.DetailService{}
	services = append(services, service1)
	server := service.DetailServer{
		Server: service.DetailInformations{
			Name:    "serverTEST",
			Service: service1,
		},
	}
	_, err := DetailService("incorrect", server)

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
	host1 := host.DetailHost{
		ID:                  "1",
		Name:                "hostTEST",
		Alias:               "hostTEST",
		Address:             "127.0.0.1",
		State:               "1",
		StateType:           "1",
		Output:              "UNKNOWN: Need to specify --status option.\n",
		MaxCheckAttempts:    "3",
		CheckAttempt:        "1",
		LastCheck:           "1589446838",
		LastStateChange:     "1589207728",
		LastHardStateChange: "1589207793",
		Acknowledged:        "0",
		Activate:            "1",
		PollerName:          "pollerTEST",
		Criticality:         "",
		PassiveChecks:       "1",
		Notify:              "1",
	}
	server := host.DetailServer{
		Server: host.DetailInformations{
			Name: "serverTEST",
			Host: host1,
		},
	}
	displayHost, err := DetailHost("json", server)
	expected, _ := json.MarshalIndent(server, "", " ")

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayDetailHostYAML(t *testing.T) {
	host1 := host.DetailHost{
		ID:                  "1",
		Name:                "hostTEST",
		Alias:               "hostTEST",
		Address:             "127.0.0.1",
		State:               "1",
		StateType:           "1",
		Output:              "UNKNOWN: Need to specify --status option.\n",
		MaxCheckAttempts:    "3",
		CheckAttempt:        "1",
		LastCheck:           "1589446838",
		LastStateChange:     "1589207728",
		LastHardStateChange: "1589207793",
		Acknowledged:        "0",
		Activate:            "1",
		PollerName:          "pollerTEST",
		Criticality:         "",
		PassiveChecks:       "1",
		Notify:              "1",
	}
	server := host.DetailServer{
		Server: host.DetailInformations{
			Name: "serverTEST",
			Host: host1,
		},
	}
	displayHost, err := DetailHost("yaml", server)
	expected, _ := yaml.Marshal(server)

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayDetailHostCSV(t *testing.T) {
	host1 := host.DetailHost{
		ID:                  "1",
		Name:                "hostTEST",
		Alias:               "hostTEST",
		Address:             "127.0.0.1",
		State:               "1",
		StateType:           "1",
		Output:              "UNKNOWN: Need to specify --status option.\n",
		MaxCheckAttempts:    "3",
		CheckAttempt:        "1",
		LastCheck:           "1589446838",
		LastStateChange:     "1589207728",
		LastHardStateChange: "1589207793",
		Acknowledged:        "1",
		Activate:            "1",
		PollerName:          "pollerTEST",
		Criticality:         "",
		PassiveChecks:       "1",
		Notify:              "1",
	}
	server := host.DetailServer{
		Server: host.DetailInformations{
			Name: "serverTEST",
			Host: host1,
		},
	}
	displayHost, err := DetailHost("csv", server)
	expected := "Server,ID,Name,Alias,IPAddress,State,StateType,Output,MaxCheckAttempts,CheckAttempt,LastCheck,LastStateChange,LastHardStateChange,Acknowledged,Activate,PollerName,Criticality,PassiveChecks,Notify\n"
	expected += server.Server.Name + ","
	expected += host1.ID + ","
	expected += host1.Name + ","
	expected += host1.Alias + ","
	expected += host1.Address + ","
	expected += "DOWN,"
	expected += "HARD,"
	expected += host1.Output + ","
	expected += host1.MaxCheckAttempts + ","
	expected += host1.CheckAttempt + ","
	expected += host1.LastCheck + ","
	expected += host1.LastStateChange + ","
	expected += host1.LastHardStateChange + ","
	expected += "yes,"
	expected += host1.Activate + ","
	expected += host1.PollerName + ","
	expected += host1.Criticality + ","
	expected += host1.PassiveChecks + ","
	expected += host1.Notify + "\n"

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayDetailHostText(t *testing.T) {
	host1 := host.DetailHost{
		ID:                  "1",
		Name:                "hostTEST",
		Alias:               "hostTEST",
		Address:             "127.0.0.1",
		State:               "1",
		StateType:           "1",
		Output:              "UNKNOWN: Need to specify --status option.\n",
		MaxCheckAttempts:    "3",
		CheckAttempt:        "1",
		LastCheck:           "1589446838",
		LastStateChange:     "1589207728",
		LastHardStateChange: "1589207793",
		Acknowledged:        "1",
		Activate:            "1",
		PollerName:          "pollerTEST",
		Criticality:         "",
		PassiveChecks:       "1",
		Notify:              "1",
	}
	server := host.DetailServer{
		Server: host.DetailInformations{
			Name: "serverTEST",
			Host: host1,
		},
	}
	displayHost, err := DetailHost("text", server)
	expected := "Host detail for server " + server.Server.Name + ": \n"
	expected += "ID: " + host1.ID + "\t"
	expected += "Name: " + host1.Name + "\t"
	expected += "Alias: " + host1.Alias + "\t"
	expected += "IP address: " + host1.Address + "\t"
	expected += "State: DOWN\t"
	expected += "State type: HARD\t"
	expected += "Output: " + host1.Output + "\t"
	expected += "Max check attempts: " + host1.MaxCheckAttempts + "\t"
	expected += "Check attempt: " + host1.CheckAttempt + "\t"
	expected += "Last check: " + host1.LastCheck + "\t"
	expected += "Last state change: " + host1.LastStateChange + "\t"
	expected += "Last hard state change: " + host1.LastHardStateChange + "\t"
	expected += "Acknowledged: yes\t"
	expected += "Activate: " + host1.Activate + "\t"
	expected += "Poller name: " + host1.PollerName + "\t"
	expected += "Criticality: " + host1.Criticality + "\t"
	expected += "Passive checks: " + host1.PassiveChecks + "\t"
	expected += "Notify: " + host1.Notify + "\n"

	assert.NoError(t, err)
	assert.Equal(t, string(expected), displayHost)
}

func TestDisplayDetailHostIncorrectOutput(t *testing.T) {
	host1 := host.DetailHost{
		ID:                  "1",
		Name:                "hostTEST",
		Alias:               "hostTEST",
		Address:             "127.0.0.1",
		State:               "1",
		StateType:           "1",
		Output:              "UNKNOWN: Need to specify --status option.\n",
		MaxCheckAttempts:    "3",
		CheckAttempt:        "1",
		LastCheck:           "1589446838",
		LastStateChange:     "1589207728",
		LastHardStateChange: "1589207793",
		Acknowledged:        "1",
		Activate:            "1",
		PollerName:          "pollerTEST",
		Criticality:         "",
		PassiveChecks:       "1",
		Notify:              "1",
	}
	server := host.DetailServer{
		Server: host.DetailInformations{
			Name: "serverTEST",
			Host: host1,
		},
	}
	_, err := DetailHost("incorrect", server)

	assert.EqualError(t, err, "The output is not correct, used : text, csv, json or yaml")
}

func TestDisplayPollerJSON(t *testing.T) {
	poller1 := poller.Poller{
		ID:        "1",
		Name:      "PollerTEST",
		IPAddress: "127.0.0.1",
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
		ID:        "1",
		Name:      "PollerTEST",
		IPAddress: "127.0.0.1",
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
		ID:        "1",
		Name:      "PollerTEST",
		IPAddress: "127.0.0.1",
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
	expected := "Server,ID,Name,IPAddress\n"
	expected += server.Server.Name + "," + poller1.ID + "," + poller1.Name + "," + poller1.IPAddress + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayPoller)
}

func TestDisplayPollerText(t *testing.T) {
	poller1 := poller.Poller{
		ID:        "1",
		Name:      "PollerTEST",
		IPAddress: "127.0.0.1",
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
	expected += "ID: " + poller1.ID + "\t"
	expected += "Name: " + poller1.Name + "\t"
	expected += "IP Address: " + poller1.IPAddress + "\n"

	assert.NoError(t, err)
	assert.Equal(t, expected, displayPoller)
}

func TestDisplayPollerIncorrectOutput(t *testing.T) {
	poller1 := poller.Poller{
		ID:        "1",
		Name:      "PollerTEST",
		IPAddress: "127.0.0.1",
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
