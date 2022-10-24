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

package poller

import (
	"centctl/resources"
	"encoding/json"
	"strconv"

	"github.com/jszwec/csvutil"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

//Poller represents the caracteristics of a poller
type Poller struct {
	ID                       int    `json:"id"`
	Name                     string `json:"name"`
	Address                  string `json:"address"`
	IsLocalhost              bool   `json:"is_localhost"`
	IsDefault                bool   `json:"is_default"`
	SSHPort                  int    `json:"ssh_port"`
	LastRestart              string `json:"last_restart"`
	EngineStartCommand       string `json:"engine_start_command"`
	EngineStopCommand        string `json:"engine_stop_command"`
	EngineRestartCommand     string `json:"engine_restart_command"`
	EngineReloadCommand      string `json:"engine_reload_command"`
	NagiosBin                string `json:"nagios_bin"`
	NagiostatsBin            string `json:"nagiostats_bin"`
	BrokerReloadCommand      string `json:"broker_reload_command"`
	CentreonbrokerCfgPath    string `json:"centreonbroker_cfg_path"`
	CentreonbrokerModulePath string `json:"centreonbroker_module_path"`
	CentreonbrokerLogsPath   string `json:"centreonbroker_logs_path"`
	CentreonconnectorPath    string `json:"centreonconnector_path"`
	InitScriptCentreontrapd  string `json:"init_script_centreontrapd"`
	SnmpTrapdPathConf        string `json:"snmp_trapd_path_conf"`
	RemoteID                 string `json:"remote_id"`
	RemoteServerUseAsProxy   bool   `json:"remote_server_use_as_proxy"`
	IsUpdated                bool   `json:"is_updated"`
	IsActivate               bool   `json:"is_activate"`
}

type ResultPoller struct {
	Pollers []Poller `json:"result" yaml:"result"`
}

//Server represents a server with informations
type Server struct {
	Server Informations `json:"server" yaml:"server"`
}

//Informations represents the informations of the server
type Informations struct {
	Name    string   `json:"name" yaml:"name"`
	Pollers []Poller `json:"pollers" yaml:"pollers"`
}

//StringText permits to display the caracteristics of the pollers to text
func (s Server) StringText() string {
	var table pterm.TableData
	table = append(table, []string{"ID", "Name", "Address", "IsLocalhost", "IsDefault", "IsUpdate", "IsActivate"})
	for i := 0; i < len(s.Server.Pollers); i++ {
		table = append(table, []string{strconv.Itoa(s.Server.Pollers[i].ID), s.Server.Pollers[i].Name, s.Server.Pollers[i].Address, strconv.FormatBool(s.Server.Pollers[i].IsLocalhost), strconv.FormatBool(s.Server.Pollers[i].IsActivate), strconv.FormatBool(s.Server.Pollers[i].IsUpdated)})
	}
	values := resources.TableListWithHeader(table)
	return values
}

//StringCSV permits to display the caracteristics of the pollers to csv
func (s Server) StringCSV() string {
	b, _ := csvutil.Marshal(s.Server.Pollers)
	return string(b)
}

//StringJSON permits to display the caracteristics of the pollers to json
func (s Server) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the pollers to yaml
func (s Server) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
