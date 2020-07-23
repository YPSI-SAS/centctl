package display

import (
	"centctl/contact"
	"centctl/host"
	"centctl/poller"
	"centctl/service"
	"fmt"
)

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
