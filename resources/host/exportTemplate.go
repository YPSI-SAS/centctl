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

package host

//ExportTemplateHost represents the caracteristics of a template host
type ExportTemplateHost struct {
	Name                       string `json:"name" yaml:"name"`
	Alias                      string `json:"alias" yaml:"alias"`
	Address                    string `json:"address" yaml:"address"`
	SnmpCommunity              string `json:"snmp_community" yaml:"snmp_community"`
	SnmpVersion                string `json:"snmp_version" yaml:"snmp_version"`
	Timezone                   string `json:"timezone" yaml:"timezone"`
	CheckCommand               string `json:"check_command" yaml:"check_command"`
	CheckCommandArguments      string `json:"check_command_arguments" yaml:"check_command_arguments"`
	CheckPeriod                string `json:"check_period" yaml:"check_period"`
	MaxCheckAttempts           string `json:"max_check_attempts" yaml:"max_check_attempts"`
	CheckInterval              string `json:"check_interval" yaml:"check_interval"`
	RetryCheckInterval         string `json:"retry_check_interval" yaml:"retry_check_interval"`
	ActiveChecksEnabled        string `json:"active_checks_enabled" yaml:"active_checks_enabled"`
	PassiveChecksEnabled       string `json:"passive_checks_enabled" yaml:"passive_checks_enabled"`
	NotificationsEnabled       string `json:"notifications_enabled" yaml:"notifications_enabled"`
	ContactAdditiveInheritance string `json:"contact_additive_inheritance" yaml:"contact_additive_inheritance"`
	CgAdditiveInheritance      string `json:"cg_additive_inheritance" yaml:"cg_additive_inheritance"`
	NotificationOptions        string `json:"notification_options" yaml:"notification_options"`
	NotificationInterval       string `json:"notification_interval" yaml:"notification_interval"`
	NotificationPeriod         string `json:"notification_period" yaml:"notification_period"`
	FirstNotificationDelay     string `json:"first_notification_delay" yaml:"first_notification_delay"`
	RecoveryNotificationDelay  string `json:"recovery_notification_delay" yaml:"recovery_notification_delay"`
	ObsessOverHost             string `json:"obsess_over_host" yaml:"obsess_over_host"`
	AcknowledgementTimeout     string `json:"acknowledgement_timeout" yaml:"acknowledgement_timeout"`
	CheckFreshness             string `json:"check_freshness" yaml:"check_freshness"`
	FreshnessThreshold         string `json:"freshness_threshold" yaml:"freshness_threshold"`
	FlapDetectionEnabled       string `json:"flap_detection_enabled" yaml:"flap_detection_enabled"`
	LowFlapThreshold           string `json:"low_flap_threshold" yaml:"low_flap_threshold"`
	HighFlapThreshold          string `json:"high_flap_threshold" yaml:"high_flap_threshold"`
	RetainStatusInformation    string `json:"retain_status_information" yaml:"retain_status_information"`
	RetainNonstatusInformation string `json:"retain_nonstatus_information" yaml:"retain_nonstatus_information"`
	StalkingOptions            string `json:"stalking_options" yaml:"stalking_options"`
	EventHandlerEnabled        string `json:"event_handler_enabled" yaml:"event_handler_enabled"`
	EventHandler               string `json:"event_handler" yaml:"event_handler"`
	EventHandlerArguments      string `json:"event_handler_arguments" yaml:"event_handler_arguments"`
	ActionURL                  string `json:"action_url" yaml:"action_url"`
	Notes                      string `json:"notes" yaml:"notes"`
	NotesURL                   string `json:"notes_url" yaml:"notes_url"`
	IconImage                  string `json:"icon_image" yaml:"icon_image"`
	IconImageAlt               string `json:"icon_image_alt" yaml:"icon_image_alt"`
	StatusMapImage             string `json:"statusmap_image" yaml:"statusmap_image"`
	Coords2d                   string `json:"2d_coords" yaml:"2d_coords"`
	Coords3d                   string `json:"3d_coords" yaml:"3d_coords"`
	Comment                    string `json:"comment" yaml:"comment"`
	Activate                   string `json:"activate" yaml:"activate"`

	Macros        []ExportHostTemplateMacro
	ContactGroups []ExportHostTemplateContactGroup
	Contacts      []ExportHostTemplateContact
	Templates     []ExportHostTemplateTemplate
}

//ExportHostTemplateResult represents a template host array send by the API
type ExportHostTemplateResult struct {
	HostTemplates []ExportTemplateHost `json:"result" yaml:"result"`
}

//ExportHostTemplateTemplate represents the caracteristics of a template
type ExportHostTemplateTemplate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultHostTemplateTemplate represents a template array send by the API
type ExportResultHostTemplateTemplate struct {
	Templates []ExportHostTemplateTemplate `json:"result" yaml:"result"`
}

//ExportHostTemplateMacro represents the caracteristics of a macro
type ExportHostTemplateMacro struct {
	Name        string `json:"macro name"`
	Value       string `json:"macro value"`
	IsPassword  string `json:"is_password"`
	Description string `json:"description"`
}

//ExportResultHostTemplateMacro represents a macro array send by the API
type ExportResultHostTemplateMacro struct {
	Macros []ExportHostTemplateMacro `json:"result" yaml:"result"`
}

//ExportHostTemplateContactGroup represents the caracteristics of a contactgroup
type ExportHostTemplateContactGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultHostTemplateContactGroup represents a contactgroup array send by the API
type ExportResultHostTemplateContactGroup struct {
	ContactGroups []ExportHostTemplateContactGroup `json:"result" yaml:"result"`
}

//ExportHostTemplateContact represents the caracteristics of a contact
type ExportHostTemplateContact struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultHostTemplateContact represents a contact array send by the API
type ExportResultHostTemplateContact struct {
	Contacts []ExportHostTemplateContact `json:"result" yaml:"result"`
}
