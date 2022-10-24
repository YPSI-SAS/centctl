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
	"fmt"
	"strconv"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailPoller represents the caracteristics of a poller
type DetailPoller struct {
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

type ResultDetailPoller struct {
	Pollers []DetailPoller `json:"result" yaml:"result"`
}

//DetailServer represents a server with informations
type DetailServer struct {
	Server DetailInformations `json:"server" yaml:"server"`
}

//DetailInformations represents the informations of the server
type DetailInformations struct {
	Name   string        `json:"name" yaml:"name"`
	Poller *DetailPoller `json:"poller" yaml:"poller"`
}

//StringText permits to display the caracteristics of the pollers to text
func (s DetailServer) StringText() string {
	var values string
	poller := s.Server.Poller
	if poller != nil {
		elements := [][]string{{"0", "Poller:"}}
		elements = append(elements, []string{"1", "ID: " + strconv.Itoa((*poller).ID)})
		elements = append(elements, []string{"1", "Name: " + (*poller).Name})
		elements = append(elements, []string{"1", "Is localhost: " + strconv.FormatBool((*poller).IsLocalhost)})
		elements = append(elements, []string{"1", "Is default: " + strconv.FormatBool((*poller).IsDefault)})
		elements = append(elements, []string{"1", "Last restart: " + (*poller).LastRestart})
		elements = append(elements, []string{"1", "Address: " + (*poller).Address})
		elements = append(elements, []string{"1", "Is activate: " + strconv.FormatBool((*poller).IsActivate)})
		elements = append(elements, []string{"1", "Engine start command: " + (*poller).EngineStartCommand})
		elements = append(elements, []string{"1", "Engine stop command: " + (*poller).EngineStopCommand})
		elements = append(elements, []string{"1", "Engine restart command: " + (*poller).EngineRestartCommand})
		elements = append(elements, []string{"1", "Engine reload command: " + (*poller).EngineReloadCommand})
		elements = append(elements, []string{"1", "Nagios bin: " + (*poller).NagiosBin})
		elements = append(elements, []string{"1", "Nagios stats bin: " + (*poller).NagiostatsBin})
		elements = append(elements, []string{"1", "Broker reload command: " + (*poller).BrokerReloadCommand})
		elements = append(elements, []string{"1", "Centreon broker cfg path: " + (*poller).CentreonbrokerCfgPath})
		elements = append(elements, []string{"1", "Centreon broker module path: " + (*poller).CentreonbrokerModulePath})
		elements = append(elements, []string{"1", "Centreon connector path: " + (*poller).CentreonconnectorPath})
		elements = append(elements, []string{"1", "SSH port: " + strconv.Itoa((*poller).SSHPort)})
		elements = append(elements, []string{"1", "Init script centreon trapd: " + (*poller).InitScriptCentreontrapd})
		elements = append(elements, []string{"1", "SNMP trapd path conf: " + (*poller).SnmpTrapdPathConf})
		elements = append(elements, []string{"1", "Centreon broker logs path: " + (*poller).CentreonbrokerLogsPath})
		elements = append(elements, []string{"1", "Remote ID: " + (*poller).RemoteID})
		elements = append(elements, []string{"1", "Remote server use as proxy: " + strconv.FormatBool((*poller).RemoteServerUseAsProxy)})
		elements = append(elements, []string{"1", "Is update: " + strconv.FormatBool((*poller).IsUpdated)})
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "poller: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the pollers to csv
func (s DetailServer) StringCSV() string {
	var p []DetailPoller
	if s.Server.Poller != nil {
		p = append(p, *s.Server.Poller)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the pollers to json
func (s DetailServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the pollers to yaml
func (s DetailServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
