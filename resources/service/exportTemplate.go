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

package service

//ExportTemplateService represents the caracteristics of a template service
type ExportTemplateService struct {
	Alias                      string `json:"alias" yaml:"alias"`
	Description                string `json:"description" yaml:"description"`
	Template                   string `json:"template" yaml:"template"`
	CheckCommand               string `json:"check_command" yaml:"check_command"`
	CheckCommandArguments      string `json:"check_command_arguments" yaml:"check_command_arguments"`
	CheckPeriod                string `json:"check_period" yaml:"check_period"`
	MaxCheckAttempts           string `json:"max_check_attempts" yaml:"max_check_attempts"`
	NormalCheckInterval        string `json:"normal_check_interval" yaml:"normal_check_interval"`
	RetryCheckInterval         string `json:"retry_check_interval" yaml:"retry_check_interval"`
	ActiveChecksEnabled        string `json:"active_checks_enabled" yaml:"active_checks_enabled"`
	PassiveChecksEnabled       string `json:"passive_checks_enabled" yaml:"passive_checks_enabled"`
	IsVolatile                 string `json:"is_volatile" yaml:"is_volatile"`
	NotificationsEnabled       string `json:"notifications_enabled" yaml:"notifications_enabled"`
	ContactAdditiveInheritance string `json:"contact_additive_inheritance" yaml:"contact_additive_inheritance"`
	CgAdditiveInheritance      string `json:"cg_additive_inheritance" yaml:"cg_additive_inheritance"`
	NotificationOptions        string `json:"notification_options" yaml:"notification_options"`
	NotificationInterval       string `json:"notification_interval" yaml:"notification_interval"`
	NotificationPeriod         string `json:"notification_period" yaml:"notification_period"`
	FirstNotificationDelay     string `json:"first_notification_delay" yaml:"first_notification_delay"`
	ObsessOverService          string `json:"obsess_over_service" yaml:"obsess_over_service"`
	CheckFreshness             string `json:"check_freshness" yaml:"check_freshness"`
	FreshnessThreshold         string `json:"freshness_threshold" yaml:"freshness_threshold"`
	FlapDetectionEnabled       string `json:"flap_detection_enabled" yaml:"flap_detection_enabled"`
	RetainStatusInformation    string `json:"retain_status_information" yaml:"retain_status_information"`
	RetainNonstatusInformation string `json:"retain_nonstatus_information" yaml:"retain_nonstatus_information"`
	EventHandlerEnabled        string `json:"event_handler_enabled" yaml:"event_handler_enabled"`
	EventHandler               string `json:"event_handler" yaml:"event_handler"`
	EventHandlerArguments      string `json:"event_handler_arguments" yaml:"event_handler_arguments"`
	ActionURL                  string `json:"action_url" yaml:"action_url"`
	Notes                      string `json:"notes" yaml:"notes"`
	NotesURL                   string `json:"notes_url" yaml:"notes_url"`
	IconImage                  string `json:"icon_image" yaml:"icon_image"`
	IconImageAlt               string `json:"icon_image_alt" yaml:"icon_image_alt"`
	Comment                    string `json:"comment" yaml:"comment"`
	Activate                   string `json:"activate" yaml:"activate"`

	Macros        []ExportServiceTemplateMacro
	ContactGroups []ExportServiceTemplateContactGroup
	Contacts      []ExportServiceTemplateContact
	Traps         []ExportServiceTemplateTrap
	Categories    []ExportServiceTemplateCategory
	HostTemplates []ExportServiceTemplateHostTemplate
}

//ExportServiceTemplateResult represents a template service array send by the API
type ExportServiceTemplateResult struct {
	TemplateServices []ExportTemplateService `json:"result" yaml:"result"`
}

//ExportServiceTemplateMacro represents the caracteristics of a macro
type ExportServiceTemplateMacro struct {
	Name        string `json:"macro name"`
	Value       string `json:"macro value"`
	IsPassword  string `json:"is_password"`
	Description string `json:"description"`
}

//ExportResultServiceTemplateMacro represents a macro array send by the API
type ExportResultServiceTemplateMacro struct {
	Macros []ExportServiceTemplateMacro `json:"result" yaml:"result"`
}

//ExportServiceTemplateContactGroup represents the caracteristics of a contactgroup
type ExportServiceTemplateContactGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultServiceTemplateContactGroup represents a contactgroup array send by the API
type ExportResultServiceTemplateContactGroup struct {
	ContactGroups []ExportServiceTemplateContactGroup `json:"result" yaml:"result"`
}

//ExportServiceTemplateContact represents the caracteristics of a contact
type ExportServiceTemplateContact struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultServiceTemplateContact represents a contact array send by the API
type ExportResultServiceTemplateContact struct {
	Contacts []ExportServiceTemplateContact `json:"result" yaml:"result"`
}

//ExportServiceTemplateTrap represents the caracteristics of a trap
type ExportServiceTemplateTrap struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultServiceTemplateTrap represents a trap array send by the API
type ExportResultServiceTemplateTrap struct {
	Traps []ExportServiceTemplateTrap `json:"result" yaml:"result"`
}

//ExportServiceTemplateCategory represents the caracteristics of a category
type ExportServiceTemplateCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultServiceTemplateCategory represents a category array send by the API
type ExportResultServiceTemplateCategory struct {
	Categories []ExportServiceTemplateCategory `json:"result" yaml:"result"`
}

//ExportServiceTemplateHostTemplate represents the caracteristics of a host template
type ExportServiceTemplateHostTemplate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ExportResultServiceTemplateHostTemplate represents a host template array send by the API
type ExportResultServiceTemplateHostTemplate struct {
	HostTemplates []ExportServiceTemplateHostTemplate `json:"result" yaml:"result"`
}
