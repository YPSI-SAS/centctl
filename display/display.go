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
	"centctl/resources/ACL"
	"centctl/resources/LDAP"
	"centctl/resources/broker"
	"centctl/resources/centreonProxy"
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
	"fmt"
)

//CentreonProxy permits to display the CentreonProxy depending on the output
func CentreonProxy(output string, server centreonProxy.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//Vendor permits to display the Vendor depending on the output
func Vendor(output string, server vendor.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailVendor permits to display the Vendor depending on the output
func DetailVendor(output string, server vendor.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//Trap permits to display the Trap depending on the output
func Trap(output string, server trap.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailTrap permits to display the Trap depending on the output
func DetailTrap(output string, server trap.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//TimePeriod permits to display the TimePeriod depending on the output
func TimePeriod(output string, server timePeriod.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailTimePeriod permits to display the DetailTimePeriod depending on the output
func DetailTimePeriod(output string, server timePeriod.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//ResourceCFG permits to display the ResourceCFG depending on the output
func ResourceCFG(output string, server resourceCFG.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailResourceCFG permits to display the ResourceCFG depending on the output
func DetailResourceCFG(output string, server resourceCFG.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//LDAPs permits to display the LDAPs depending on the output
func LDAPs(output string, server LDAP.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailLDAP permits to display the LDAPs depending on the output
func DetailLDAP(output string, server LDAP.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//Dependency permits to display the dependencies depending on the output
func Dependency(output string, server dependency.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailDependency permits to display the dependencies depending on the output
func DetailDependency(output string, server dependency.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//Contact permits to display the contacts depending on the output
func Contact(output string, server contact.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailContact permits to display the contacts depending on the output
func DetailContact(output string, server contact.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//RealtimeService permits to display the services depending on the output
func RealtimeService(output string, server service.RealtimeServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//Service permits to display the services depending on the output
func Service(output string, server service.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailService permits to display the services depending on the output
func DetailService(output string, server service.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailRealtimeService permits to display the services depending on the output
func DetailRealtimeService(output string, server service.DetailRealtimeServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailTimelineService permits to display the service depending on the output
func DetailTimelineService(output string, server service.DetailTimelineServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//RealtimeHost permits to display the hosts depending on the output
func RealtimeHost(output string, server host.RealtimeServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//RealtimeHostV2 permits to display the hosts depending on the output
func RealtimeHostV2(output string, server host.RealtimeServerV2) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//Host permits to display the hosts depending on the output
func Host(output string, server host.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailTimelineHost permits to display the host depending on the output
func DetailTimelineHost(output string, server host.DetailTimelineServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailRealtimeHost permits to display the host depending on the output
func DetailRealtimeHost(output string, server host.DetailRealtimeServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailHost permits to display the host depending on the output
func DetailHost(output string, server host.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//Poller permits to display the pollers depending on the output
func Poller(output string, server poller.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailPoller permits to display the pollers depending on the output
func DetailPoller(output string, server poller.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//TemplateHost permits to display the host templates depending on the output
func TemplateHost(output string, server host.TemplateServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailTemplateHost permits to display the template host depending on the output
func DetailTemplateHost(output string, server host.DetailTemplateServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailTemplateContact permits to display the template contact depending on the output
func DetailTemplateContact(output string, server contact.DetailTemplateServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailTemplateService permits to display the template service depending on the output
func DetailTemplateService(output string, server service.DetailTemplateServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//TemplateService permits to display the service templates depending on the output
func TemplateService(output string, server service.TemplateServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//GroupContact permits to display the contact groups depending on the output
func GroupContact(output string, server contact.GroupServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//GroupHost permits to display the host groups depending on the output
func GroupHost(output string, server host.GroupServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//GroupService permits to display the service groups depending on the output
func GroupService(output string, server service.GroupServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//CategoryHost permits to display the host categories depending on the output
func CategoryHost(output string, server host.CategoryServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//CategoryService permits to display the service categories depending on the output
func CategoryService(output string, server service.CategoryServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailCategoryHost permits to display the host categoriy depending on the output
func DetailCategoryHost(output string, server host.DetailCategoryServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailCategoryService permits to display the service categoriy depending on the output
func DetailCategoryService(output string, server service.DetailCategoryServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//ACLGroup permits to display the ACL group depending on the output
func ACLGroup(output string, server ACL.GroupServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//ACLAction permits to display the ACL action depending on the output
func ACLAction(output string, server ACL.ActionServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//ACLMenu permits to display the ACL menu depending on the output
func ACLMenu(output string, server ACL.MenuServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//ACLResource permits to display the ACL resource depending on the output
func ACLResource(output string, server ACL.ResourceServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailACLGroup permits to display the ACL group depending on the output
func DetailACLGroup(output string, server ACL.DetailGroupServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailACLAction permits to display the ACL action depending on the output
func DetailACLAction(output string, server ACL.DetailActionServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailACLMenu permits to display the ACL menu depending on the output
func DetailACLMenu(output string, server ACL.DetailMenuServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailACLResource permits to display the ACL resource depending on the output
func DetailACLResource(output string, server ACL.DetailResourceServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//Command permits to display the command depending on the output
func Command(output string, server command.Server) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailCommand permits to display the command depending on the output
func DetailCommand(output string, server command.DetailServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//BrokerCFG permits to display the BrokerCFG depending on the output
func BrokerCFG(output string, server broker.ServerCFG) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//BrokerInput permits to display the BrokerInput depending on the output
func BrokerInput(output string, server broker.ServerInput) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//BrokerOutput permits to display the BrokerOutput depending on the output
func BrokerOutput(output string, server broker.ServerOutput) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//BrokerLogger permits to display the BrokerLogger depending on the output
func BrokerLogger(output string, server broker.ServerLogger) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailBrokerCFG permits to display the BrokerCFG depending on the output
func DetailBrokerCFG(output string, server broker.DetailServerCFG) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailBrokerInput permits to display the Broker input depending on the output
func DetailBrokerInput(output string, server broker.DetailServerInput) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailBrokerOutput permits to display the Broker output depending on the output
func DetailBrokerOutput(output string, server broker.DetailServerOutput) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailBrokerLogger permits to display the Broker logger depending on the output
func DetailBrokerLogger(output string, server broker.DetailServerLogger) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//EngineCFG permits to display the EngineCFG depending on the output
func EngineCFG(output string, server engineCFG.ServerEngineCFG) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailEngineCFG permits to display the EngineCFG depending on the output
func DetailEngineCFG(output string, server engineCFG.DetailServerEngineCFG) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailGroupContact permits to display the group contact depending on the output
func DetailGroupContact(output string, server contact.DetailGroupServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailGroupHost permits to display the group host depending on the output
func DetailGroupHost(output string, server host.DetailGroupServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//DetailGroupService permits to display the group service depending on the output
func DetailGroupService(output string, server service.DetailGroupServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}

//TemplateContact permits to display the contact templates depending on the output
func TemplateContact(output string, server contact.TemplateServer) (string, error) {
	switch output {
	case "json":
		return server.StringJSON(), nil
	case "csv":
		return server.StringCSV(), nil
	case "yaml":
		return server.StringYAML(), nil
	case "text":
		return server.StringText(), nil
	default:
		return "", fmt.Errorf("The output is not correct, used : text, csv, json or yaml")
	}
}
