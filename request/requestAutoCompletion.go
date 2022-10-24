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

package request

import (
	"centctl/resources/ACL"
	"centctl/resources/LDAP"
	"centctl/resources/broker"
	"centctl/resources/command"
	"centctl/resources/contact"
	"centctl/resources/dependency"
	"centctl/resources/engineCFG"
	"centctl/resources/host"
	"centctl/resources/poller"
	"centctl/resources/resourceCFG"
	"centctl/resources/service"
	"centctl/resources/timePeriod"
	"centctl/resources/trap"
	"centctl/resources/vendor"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type ServerList struct {
	Servers []struct {
		Server   string `yaml:"server"`
		Url      string `yaml:"url"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Version  string `yaml:"version"`
		Default  string `yaml:"default,omitempty"`
		Insecure string `yaml:"insecure,omitempty"`
		Proxy    []struct {
			HttpURL  string `yaml:"httpURL"`
			HttpsURL string `yaml:"httpsURL"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"proxy,omitempty"`
	} `yaml:"servers"`
}

//InitAuthentification permits to init authentification if a server is find for auto completion
func InitAuthentification(cmd *cobra.Command) bool {
	var server string
	if cmd.Root().Flag("server").Value.String() != "" {
		server = cmd.Root().Flag("server").Value.String()
	} else {
		server = getDefaultServer()
	}
	if server != "" {
		authentificationToServer(server)
		return true
	} else {
		return false
	}
}

//getDefaultServer permits to get the default server in a file
func getDefaultServer() string {
	var server string
	if os.Getenv("CENTCTL_CONF") != "" {
		servers := &ServerList{}
		yamlFile, _ := ioutil.ReadFile(os.Getenv("CENTCTL_CONF"))
		_ = yaml.Unmarshal(yamlFile, servers)
		for _, serv := range servers.Servers {
			if serv.Default == "true" {
				server = serv.Server
			}
		}
	}
	return server
}

//authentificationToServer permits to generate the token for the server
func authentificationToServer(name string) {
	servers := &ServerList{}
	var token string
	var versionAPI string
	yamlFile, _ := ioutil.ReadFile(os.Getenv("CENTCTL_CONF"))
	_ = yaml.Unmarshal(yamlFile, servers)
	for _, serv := range servers.Servers {
		if serv.Server == name {
			if serv.Version == "v1" {
				insecure := false
				if serv.Insecure == "true" {
					insecure = true
				}
				os.Setenv("URL", serv.Url)

				token, _ = AuthentificationV1(serv.Url, serv.Login, serv.Password, insecure)
			} else {
				insecure := false
				if serv.Insecure == "true" {
					insecure = true
				}
				os.Setenv("URL", serv.Url)
				token, versionAPI, _ = AuthentificationV2(serv.Url, serv.Login, serv.Password, insecure, "/beta")
				os.Setenv("VERSIONAPI", versionAPI)
			}
		}
	}
	_ = os.Setenv("TOKEN", token)

}

//GetACLActionNames permits to get ACL action names for auto completion
func GetACLActionNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "ACLACTION", "", "list ACL action", false, false, "")
	actions := ACL.ResultAction{}
	json.Unmarshal(body, &actions)

	if len(actions.Actions) != 0 {
		for _, action := range actions.Actions {
			values = append(values, "\""+action.Name+"\"")
		}
	}

	return values
}

//GetACLGroupNames permits to get ACL group names for auto completion
func GetACLGroupNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "ACLGROUP", "", "list ACL group", false, false, "")
	groups := ACL.ResultGroup{}
	json.Unmarshal(body, &groups)

	if len(groups.Groups) != 0 {
		for _, group := range groups.Groups {
			values = append(values, "\""+group.Name+"\"")
		}
	}

	return values
}

//GetACLMenuNames permits to get ACL menu names for auto completion
func GetACLMenuNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "ACLMENU", "", "list ACL menu", false, false, "")
	menus := ACL.ResultMenu{}
	json.Unmarshal(body, &menus)

	if len(menus.Menus) != 0 {
		for _, menu := range menus.Menus {
			values = append(values, "\""+menu.Name+"\"")
		}
	}

	return values
}

//GetACLResourceNames permits to get ACL resource names for auto completion
func GetACLResourceNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "ACLRESOURCE", "", "list ACL resource", false, false, "")
	resources := ACL.ResultResource{}
	json.Unmarshal(body, &resources)

	if len(resources.Resources) != 0 {
		for _, resource := range resources.Resources {
			values = append(values, "\""+resource.Name+"\"")
		}
	}

	return values
}

//GetBrokerCFGNames permits to get broker CFG names for auto completion
func GetBrokerCFGNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "CENTBROKERCFG", "", "list broker CFG", false, false, "")
	brokerCFG := broker.ResultCFG{}
	json.Unmarshal(body, &brokerCFG)

	if len(brokerCFG.BrokerCFGs) != 0 {
		for _, broker := range brokerCFG.BrokerCFGs {
			values = append(values, "\""+broker.Name+"\"")
		}
	}

	return values
}

//GetBrokerInputID permits to get broker input ID for auto completion
func GetBrokerInputID(brokerName string) []string {
	var values []string

	_, body := GeneriqueCommandV1Post("listinput", "CENTBROKERCFG", brokerName, "list broker input", false, false, "")
	brokerInput := broker.ResultInput{}
	json.Unmarshal(body, &brokerInput)

	if len(brokerInput.BrokerInputs) != 0 {
		for _, input := range brokerInput.BrokerInputs {
			values = append(values, "\""+input.ID+"\"")
		}
	}

	return values
}

//GetBrokerLoggerID permits to get broker Logger ID for auto completion
func GetBrokerLoggerID(brokerName string) []string {
	var values []string

	_, body := GeneriqueCommandV1Post("listLogger", "CENTBROKERCFG", brokerName, "list broker logger", false, false, "")
	brokerLogger := broker.ResultLogger{}
	json.Unmarshal(body, &brokerLogger)

	if len(brokerLogger.BrokerLoggers) != 0 {
		for _, logger := range brokerLogger.BrokerLoggers {
			values = append(values, "\""+logger.ID+"\"")
		}
	}

	return values
}

//GetBrokerOutputID permits to get broker Output ID for auto completion
func GetBrokerOutputID(brokerName string) []string {
	var values []string

	_, body := GeneriqueCommandV1Post("listOutput", "CENTBROKERCFG", brokerName, "list broker output", false, false, "")
	brokerOutput := broker.ResultOutput{}
	json.Unmarshal(body, &brokerOutput)

	if len(brokerOutput.BrokerOutputs) != 0 {
		for _, output := range brokerOutput.BrokerOutputs {
			values = append(values, "\""+output.ID+"\"")
		}
	}

	return values
}

//GetCategoryHostNames permits to get host category name for auto completion
func GetCategoryHostNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "HC", "", "list catagory host", false, false, "")
	category := host.ResultCategory{}
	json.Unmarshal(body, &category)

	if len(category.Categories) != 0 {
		for _, cat := range category.Categories {
			values = append(values, "\""+cat.Name+"\"")
		}
	}

	return values
}

//GetCategoryServiceNames permits to get service category name for auto completion
func GetCategoryServiceNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "SC", "", "list catagory Service", false, false, "")
	category := service.ResultCategory{}
	json.Unmarshal(body, &category)

	if len(category.Categories) != 0 {
		for _, cat := range category.Categories {
			values = append(values, "\""+cat.Name+"\"")
		}
	}

	return values
}

//GetGroupContactNames permits to get group contact name for auto completion
func GetGroupContactNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "CG", "", "list contact group", false, false, "")
	groups := contact.ResultGroup{}
	json.Unmarshal(body, &groups)

	if len(groups.Groups) != 0 {
		for _, group := range groups.Groups {
			values = append(values, "\""+group.Name+"\"")
		}
	}

	return values
}

//GetGroupHostNames permits to get group Host name for auto completion
func GetGroupHostNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "HG", "", "list Host group", false, false, "")
	groups := host.ResultGroup{}
	json.Unmarshal(body, &groups)

	if len(groups.Groups) != 0 {
		for _, group := range groups.Groups {
			values = append(values, "\""+group.Name+"\"")
		}
	}

	return values
}

//GetGroupServiceNames permits to get group Service name for auto completion
func GetGroupServiceNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "SG", "", "list Service group", false, false, "")
	groups := service.ResultGroup{}
	json.Unmarshal(body, &groups)

	if len(groups.Groups) != 0 {
		for _, group := range groups.Groups {
			values = append(values, "\""+group.Name+"\"")
		}
	}

	return values
}

//GetTemplateContactAlias permits to get template contact alias for auto completion
func GetTemplateContactAlias() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "CONTACTTPL", "", "list template contact", false, false, "")
	templates := contact.ResultTemplate{}
	json.Unmarshal(body, &templates)

	if len(templates.Templates) != 0 {
		for _, template := range templates.Templates {
			values = append(values, "\""+template.Alias+"\"")
		}
	}

	return values
}

//GetTemplateHostNames permits to get template Host names for auto completion
func GetTemplateHostNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "HTPL", "", "list template Host", false, false, "")
	templates := host.ResultTemplate{}
	json.Unmarshal(body, &templates)

	if len(templates.Templates) != 0 {
		for _, template := range templates.Templates {
			values = append(values, "\""+template.Name+"\"")
		}
	}

	return values
}

//GetTemplateServiceNames permits to get template Service names for auto completion
func GetTemplateServiceNames() []string {
	var values []string
	_, body := GeneriqueCommandV1Post("show", "STPL", "", "list template Service", false, false, "")
	templates := service.ResultTemplate{}
	json.Unmarshal(body, &templates)
	if len(templates.Templates) != 0 {
		for _, template := range templates.Templates {
			values = append(values, "\""+template.Description+"\"")

		}
	}

	return values
}

//GetCommandNames permits to get command names for auto completion
func GetCommandNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "CMD", "", "list command", false, false, "")
	commands := command.Result{}
	json.Unmarshal(body, &commands)

	if len(commands.Commands) != 0 {
		for _, cmd := range commands.Commands {
			values = append(values, "\""+cmd.Name+"\"")
		}
	}

	return values
}

//GetContactAlias permits to get contact alias for auto completion
func GetContactAlias() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "contact", "", "list contact alias", false, false, "")
	contacts := contact.Result{}
	json.Unmarshal(body, &contacts)

	if len(contacts.Contacts) != 0 {
		for _, contact := range contacts.Contacts {
			values = append(values, "\""+contact.Alias+"\"")
		}
	}

	return values
}

//GetDependencyNames permits to get Dependency names for auto completion
func GetDependencyNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "DEP", "", "list Dependency", false, false, "")
	dependencys := dependency.Result{}
	json.Unmarshal(body, &dependencys)

	if len(dependencys.Dependencies) != 0 {
		for _, dep := range dependencys.Dependencies {
			values = append(values, "\""+dep.Name+"\"")
		}
	}

	return values
}

//GetEngineCFGNames permits to get EngineCFG names for auto completion
func GetEngineCFGNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "ENGINECFG", "", "list EngineCFG", false, false, "")
	enginesCFGs := engineCFG.ResultEngineCFG{}
	json.Unmarshal(body, &enginesCFGs)

	if len(enginesCFGs.EngineCFG) != 0 {
		for _, engineCFG := range enginesCFGs.EngineCFG {
			values = append(values, "\""+engineCFG.Name+"\"")
		}
	}

	return values
}

//GetHostNames permits to get Host names for auto completion
func GetHostNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "Host", "", "list host", false, false, "")
	hosts := host.Result{}
	json.Unmarshal(body, &hosts)

	if len(hosts.Hosts) != 0 {
		for _, host := range hosts.Hosts {
			values = append(values, "\""+host.Name+"\"")
		}
	}

	return values
}

//GetLDAPNames permits to get LDAP names for auto completion
func GetLDAPNames() []string {
	var values []string

	_, body := GeneriqueCommandV1Post("show", "LDAP", "", "list LDAP", false, false, "")
	LDAPs := LDAP.Result{}
	json.Unmarshal(body, &LDAPs)

	if len(LDAPs.LDAP) != 0 {
		for _, ldap := range LDAPs.LDAP {
			values = append(values, "\""+ldap.Name+"\"")
		}
	}

	return values
}

//GetPollerNames permits to get Poller names for auto completion
func GetPollerNames() []string {
	var values []string

	urlCentreon := "/configuration/monitoring-servers"
	_, body := GeneriqueCommandV2Get(urlCentreon, "list poller", false)

	var pollerResult poller.RealtimeResultPoller
	json.Unmarshal(body, &pollerResult)

	if len(pollerResult.Pollers) != 0 {
		for _, poller := range pollerResult.Pollers {
			values = append(values, "\""+poller.Name+"\"")
		}
	}

	return values
}

//GetResourceCFGNames permits to get resource CFG names for auto completion
func GetResourceCFGNames() []string {
	var values []string
	_, body := GeneriqueCommandV1Post("show", "RESOURCECFG", "", "list resource CFG", false, false, "")
	resourceCFGs := resourceCFG.Result{}
	json.Unmarshal(body, &resourceCFGs)

	if len(resourceCFGs.ResourceCFG) != 0 {
		for _, resource := range resourceCFGs.ResourceCFG {
			val := resource.Name
			val = strings.ReplaceAll(val, "$", "")
			values = append(values, "\""+val+"\"")
		}
	}

	return values
}

//GetServiceDescriptions permits to get service descriptions for auto completion
func GetServiceDescriptions(hostName string) []string {
	var values []string
	_, body := GeneriqueCommandV1Post("show", "SERVICE", hostName+";", "list resource CFG", false, false, "")
	services := service.Result{}
	json.Unmarshal(body, &services)

	if len(services.Services) != 0 {
		for _, resource := range services.Services {
			if strings.ToLower(resource.HostName) == strings.ToLower(hostName) {
				values = append(values, "\""+resource.Description+"\"")
			}
		}
	}

	return values
}

//GetTrapNames permits to get trap names for auto completion
func GetTrapNames() []string {
	var values []string
	_, body := GeneriqueCommandV1Post("show", "TRAP", "", "list trap", false, false, "")
	traps := trap.Result{}
	json.Unmarshal(body, &traps)

	if len(traps.Traps) != 0 {
		for _, trap := range traps.Traps {
			values = append(values, "\""+trap.Name+"\"")
		}
	}

	return values
}

//GetVendorNames permits to get Vendor names for auto completion
func GetVendorNames() []string {
	var values []string
	_, body := GeneriqueCommandV1Post("show", "Vendor", "", "list Vendor", false, false, "")
	vendors := vendor.Result{}
	json.Unmarshal(body, &vendors)

	if len(vendors.Vendors) != 0 {
		for _, vendor := range vendors.Vendors {
			values = append(values, "\""+vendor.Name+"\"")
		}
	}

	return values
}

//GetTimePeriodNames permits to get time period names for auto completion
func GetTimePeriodNames() []string {
	var values []string
	_, body := GeneriqueCommandV1Post("show", "TP", "", "list Vendor", false, false, "")
	timePeriods := timePeriod.Result{}
	json.Unmarshal(body, &timePeriods)

	if len(timePeriods.TimePeriods) != 0 {
		for _, tp := range timePeriods.TimePeriods {
			values = append(values, "\""+tp.Name+"\"")
		}
	}

	return values
}
