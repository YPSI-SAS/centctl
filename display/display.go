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
