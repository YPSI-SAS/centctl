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

package cmd

import (
	"bufio"
	"centctl/cmd/add"
	"centctl/cmd/add/acl"
	"centctl/cmd/add/broker"
	"centctl/cmd/add/category"
	"centctl/cmd/add/group"
	"centctl/cmd/add/template"
	"centctl/cmd/modify"
	"centctl/colorMessage"
	"centctl/request"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	applyPoller "centctl/cmd/apply"
	mACL "centctl/cmd/modify/acl"
	mBroker "centctl/cmd/modify/broker"
	mCategory "centctl/cmd/modify/category"
	mGroup "centctl/cmd/modify/group"
	mTemplate "centctl/cmd/modify/template"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import the objects contained in the csv file",
	Long:  `Import the objects contained in the csv file`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		detail, _ := cmd.Flags().GetBool("DETAIL")
		ImportCSV(file, debugV, apply, detail)
	},
}

//ImportCSV permits to import objects
func ImportCSV(file string, debugV bool, apply bool, detail bool) {
	//colorRed := "\033[1;31m%s\033[0m"
	if file != "" {
		importCSVFile(file, debugV, detail)
	} else {
		importStdin(debugV, detail)
	}
	if apply {
		exportPoller(debugV)
	}
}

//importStdin permits to read stdin and import objects
func importStdin(debugV bool, detail bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := scanner.Text()
		record := strings.Split(ucl, ",")
		if len(record) != 1 {
			for i, r := range record {
				if len(r) != 0 {
					if r[0:1] == "\"" {
						record[i] = r[1:]
					}
				}
			}
			for i, r := range record {
				if len(r) != 0 {
					if r[len(r)-1:] == "\"" {
						record[i] = r[:len(r)-1]
					}
				}
			}
			if validation(record) {
				executeActionOnObject(record, debugV, detail)
			}
		}
	}

}

//importCSVFile permits to read CSV file and import objects
func importCSVFile(file string, debugV bool, detail bool) {
	//Verification that the file has a correct extension
	if !strings.Contains(file, ".csv") {
		fmt.Println("The extension of the file must be .csv")
		os.Exit(1)
	}

	okFile := true
	//Creation of a reader on the csv
	read, _ := ioutil.ReadFile(file)
	r := csv.NewReader(strings.NewReader(string(read)))
	r.Comma = ','
	r.Comment = '#'
	for {
		//Reading line by line
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		okFile = validation(record)
	}

	if okFile {
		//Creation of a reader on the csv
		read, _ := ioutil.ReadFile(file)
		r := csv.NewReader(strings.NewReader(string(read)))
		r.FieldsPerRecord = -1
		r.Comma = ','
		r.Comment = '#'
		for {
			//Reading line by line
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			executeActionOnObject(record, debugV, detail)
		}

	}
}

//executeActionObject permits to execute add or modify action on the object
func executeActionOnObject(record []string, debugV bool, detail bool) {
	var err error
	colorOrange := colorMessage.GetColorOrange()

	//switch on the object
	switch strings.ToLower(record[0]) {
	case "add":
		switch strings.ToLower(record[1]) {

		case "aclaction":
			//Recovery the values of the ACL Action
			name := record[2]
			alias := record[3]

			//Creation of ACL Action
			err = acl.AddACLAction(name, alias, debugV, true)

		case "aclgroup":
			//Recovery the values of the ACL Group
			name := record[2]
			alias := record[3]

			//Creation of ACL Group
			err = acl.AddACLGroup(name, alias, debugV, true)

		case "aclmenu":
			//Recovery the values of the ACL Menu
			name := record[2]
			alias := record[3]

			//Creation of ACL Menu
			err = acl.AddACLMenu(name, alias, debugV, true)

		case "aclresource":
			//Recovery the values of the ACL Resource
			name := record[2]
			alias := record[3]

			//Creation of ACL Resource
			err = acl.AddACLResource(name, alias, debugV, true)

		case "brokercfg":
			//Recovery the values of the broker CFG
			name := record[2]
			instance := record[3]

			//Creation of broker CFG
			err = broker.AddBrokerCFG(name, instance, debugV, true)

		case "brokerinput":
			//Recovery the values of the broker intput
			name := record[2]
			objectName := record[3]
			objectNature := record[4]

			//Creation of broker intput
			err = broker.AddBrokerInput(name, objectName, objectNature, debugV, true)

		case "brokerlogger":
			//Recovery the values of the broker logger
			name := record[2]
			objectName := record[3]
			objectNature := record[4]

			//Creation of broker logger
			err = broker.AddBrokerLogger(name, objectName, objectNature, debugV, true)

		case "brokeroutput":
			//Recovery the values of the broker output
			name := record[2]
			objectName := record[3]
			objectNature := record[4]

			//Creation of broker output
			err = broker.AddBrokerOutput(name, objectName, objectNature, debugV, true)

		case "categoryhost":
			//Recovery the values of the category host
			name := record[2]
			alias := record[3]

			//Creation of category host
			err = category.AddCategoryHost(name, alias, debugV, true)

		case "categoryservice":
			//Recovery the values of the category service
			name := record[2]
			description := record[3]

			//Creation of category service
			err = category.AddCategoryService(name, description, debugV, true)

		case "centreonproxy":
			//Recovery the values of the centreonproxy
			url := record[2]
			login := record[3]
			password := record[4]
			port := record[5]
			portInt, _ := strconv.Atoi(port)

			//Creation of centreonproxy
			err = add.AddCentreonProxy(url, login, password, portInt, debugV, true)

		case "command":
			//Recovery the values of the command
			name := record[2]
			typeCmd := record[3]
			line := record[4]

			//Creation of command
			err = add.AddCommand(name, typeCmd, line, debugV, true)

		case "contact":
			//Recovery the values of the contact
			name := record[2]
			alias := record[3]
			email := record[4]
			password := record[5]

			//Transformation of the admin record to bool
			admin := record[6]
			adminV := true
			if admin == "" {
				adminV = false
			}

			//Creation of contact
			err = add.AddContact(name, alias, email, password, adminV, debugV, true)

		case "dependency":
			//Recovery the values of the depedency
			name := record[2]
			description := record[3]
			typeD := record[4]
			parentName := record[5]

			//Creation of depedency
			err = add.AddDependencie(name, description, typeD, parentName, debugV, true)

		case "enginecfg":
			//Recovery the values of the engine CFG
			name := record[2]
			instance := record[3]
			comment := record[4]

			//Creation of engine CFG
			err = add.AddEngineCFG(name, instance, comment, debugV, true)

		case "groupcontact":
			//Recovery the values of the group contact
			name := record[2]
			alias := record[3]

			//Creation of group contact
			err = group.AddGroupContact(name, alias, debugV, true)

		case "grouphost":
			//Recovery the values of the group host
			name := record[2]
			alias := record[3]

			//Creation of group host
			err = group.AddGroupHost(name, alias, debugV, true)

		case "groupservice":
			//Recovery the values of the group service
			name := record[2]
			alias := record[3]

			//Creation of group service
			err = group.AddGroupService(name, alias, debugV, true)

		case "host":
			//Recovery the values of the host
			name := record[2]
			alias := record[3]
			IPaddress := record[4]
			template := record[5]
			poller := record[6]
			hostGroup := record[7]

			//Creation of host
			err = add.AddHost(name, alias, IPaddress, template, poller, hostGroup, debugV, false, true)

		case "ldap":
			//Recovery the values of the LDAP
			name := record[2]
			description := record[3]

			//Creation of LDAP
			err = add.AddLDAP(name, description, debugV, true)

		case "poller":
			//Recovery the values of the poller
			name := record[2]
			IPaddress := record[3]
			SSHPort := record[4]
			SSHPortInt, _ := strconv.Atoi(SSHPort)
			connProtocol := record[5]
			portConn := record[6]
			portConnInt, _ := strconv.Atoi(portConn)

			//Creation of poller
			err = add.AddPoller(name, IPaddress, SSHPortInt, connProtocol, portConnInt, debugV, true)

		case "resourcecfg":
			//Recovery the values of the resourceCFG
			name := record[2]
			value := record[3]
			instance := record[4]
			comment := record[5]

			//Creation of resourceCFG
			err = add.AddResourceCFG(name, value, instance, comment, debugV, true)

		case "service":
			//Recovery the values of the service
			hostName := record[2]
			description := record[3]
			template := record[4]

			//Creation of service
			err = add.AddService(hostName, description, template, debugV, false, true)

		case "templatecontact":
			//Recovery the values of the template contact
			name := record[2]
			alias := record[3]

			//Creation of template contact
			err = template.AddTemplateContact(name, alias, debugV, true)

		case "templatehost":
			//Recovery the values of the template host
			name := record[2]
			alias := record[3]
			templateH := record[4]

			//Creation of template host
			err = template.AddTemplateHost(name, alias, templateH, debugV, true)

		case "templateservice":
			//Recovery the values of the template service
			name := record[2]
			alias := record[3]
			templateS := record[4]

			//Creation of template service
			err = template.AddTemplateService(name, alias, templateS, debugV, true)

		case "timeperiod":
			//Recovery the values of the timePeriod
			name := record[2]
			alias := record[3]

			//Creation of timePeriod
			err = add.AddTimePeriod(name, alias, debugV, true)

		case "trap":
			//Recovery the values of the trap
			name := record[2]
			oid := record[3]

			//Creation of trap
			err = add.AddTrap(name, oid, debugV, true)

		case "vendor":
			//Recovery the values of the vendor
			name := record[2]
			alias := record[3]

			//Creation of vendor
			err = add.AddVendor(name, alias, debugV, true)

		}
		if err != nil {
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println(err.Error())
		}
	case "modify":
		var value string
		switch strings.ToLower(record[1]) {
		case "aclaction":
			//Recovery the values of the ACL Action
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of ACL Action
			err = mACL.ModifyACLAction(name, parameter, value, debugV, false, true, detail)

		case "aclgroup":
			//Recovery the values of the ACL Group
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of ACL Group
			err = mACL.ModifyACLGroup(name, parameter, value, debugV, false, true, detail)

		case "aclmenu":
			//Recovery the values of the ACL Menu
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of ACL Menu
			err = mACL.ModifyACLMenu(name, parameter, value, debugV, false, true, detail)

		case "aclresource":
			//Recovery the values of the ACL Resource
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of ACL Resource
			err = mACL.ModifyACLResource(name, parameter, value, debugV, false, true, detail)

		case "brokercfg":
			//Recovery the values of the broker CFG
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of broker CFG
			err = mBroker.ModifyBrokerCFG(name, parameter, value, debugV, true, detail)

		case "brokerinput":
			//Recovery the values of the broker intput
			name := record[2]
			id := record[3]
			idInt, _ := strconv.Atoi(id)
			parameter := record[4]
			value = record[5]

			//Creation of broker intput
			err = mBroker.ModifyBrokerInput(name, idInt, parameter, value, debugV, true, detail)

		case "brokerlogger":
			//Recovery the values of the broker logger
			name := record[2]
			id := record[3]
			idInt, _ := strconv.Atoi(id)
			parameter := record[4]
			value = record[5]

			//Creation of broker logger
			err = mBroker.ModifyBrokerLogger(name, idInt, parameter, value, debugV, true, detail)

		case "brokeroutput":
			//Recovery the values of the broker output
			name := record[2]
			id := record[3]
			idInt, _ := strconv.Atoi(id)
			parameter := record[4]
			value = record[5]

			//Creation of broker output
			err = mBroker.ModifyBrokerOutput(name, idInt, parameter, value, debugV, true, detail)

		case "categoryhost":
			//Recovery the values of the category host
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of category host
			err = mCategory.ModifyCategoryHost(name, parameter, value, debugV, true, detail)

		case "categoryservice":
			//Recovery the values of the category service
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of category service
			err = mCategory.ModifyCategoryService(name, parameter, value, debugV, true, detail)

		case "command":
			//Recovery the values of the command
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of command
			err = modify.ModifyCommand(name, parameter, value, debugV, true, detail)

		case "contact":
			//Recovery the values of the contact
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of command
			err = modify.ModifyContact(name, parameter, value, debugV, true, detail)

		case "dependency":
			//Recovery the values of the dependency
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of dependency
			err = modify.ModifyDependency(name, parameter, value, debugV, true, detail)

		case "enginecfg":
			//Recovery the values of the engine CFG
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of engine CFG
			err = modify.ModifyEngineCFG(name, parameter, value, debugV, true, detail)

		case "groupcontact":
			//Recovery the values of the group contact
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of group contact
			err = mGroup.ModifyGroupContact(name, parameter, value, debugV, true, detail)

		case "grouphost":
			//Recovery the values of the group host
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of group host
			err = mGroup.ModifyGroupHost(name, parameter, value, debugV, true, detail)

		case "groupservice":
			//Recovery the values of the group service
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of group service
			err = mGroup.ModifyGroupService(name, parameter, value, debugV, true, detail)

		case "host":
			//Recovery the values of the host
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of group host
			err = modify.ModifyHost(name, parameter, value, debugV, false, true, detail)

		case "ldap":
			//Recovery the values of the LDAP
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of LDAP
			err = modify.ModifyLDAP(name, parameter, value, debugV, true, detail)

		case "poller":
			//Recovery the values of the poller
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of LDAP
			err = modify.ModifyPoller(name, parameter, value, debugV, false, true, detail)

		case "resourcecfg":
			//Recovery the values of the resourceCFG
			id := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of resourceCFG
			err = modify.ModifyResourceCFG(id, parameter, value, debugV, true, detail)

		case "service":
			//Recovery the values of the service
			hostName := record[2]
			description := record[3]
			parameter := record[4]
			value = record[5]

			//Modification of service
			err = modify.ModifyService(hostName, description, parameter, value, debugV, false, true, detail)

		case "templatecontact":
			//Recovery the values of the template contact
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of template contact
			err = mTemplate.ModifyTemplateContact(name, parameter, value, debugV, true, detail)

		case "templatehost":
			//Recovery the values of the template host
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of template host
			err = mTemplate.ModifyTemplateHost(name, parameter, value, debugV, true, detail)

		case "templateservice":
			//Recovery the values of the template service
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of template service
			err = mTemplate.ModifyTemplateService(name, parameter, value, debugV, true, detail)

		case "timeperiod":
			//Recovery the values of the timePeriod
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of timePeriod
			err = modify.ModifyTimePeriod(name, parameter, value, debugV, true, detail)

		case "trap":
			//Recovery the values of the trap
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of trap
			err = modify.ModifyTrap(name, parameter, value, debugV, true, detail)

		case "vendor":
			//Recovery the values of the vendor
			name := record[2]
			parameter := record[3]
			value = record[4]

			//Modification of vendor
			err = modify.ModifyVendor(name, parameter, value, debugV, true, detail)
		}
		if err != nil {
			if value != "" {
				fmt.Printf(colorRed, "ERROR: ")
				fmt.Println(err.Error())
			} else {
				fmt.Printf(colorOrange, "WARNING: ")
				fmt.Println(err.Error() + "This value is not set.")
			}
		}
	}
}

//validation permits to verify the action line
func validation(record []string) bool {
	colorRed := colorMessage.GetColorRed()
	okFile := true
	//switch on the object
	switch strings.ToLower(record[0]) {
	case "add":
		switch strings.ToLower(record[1]) {
		case "aclaction":
			okFile = ValidateLine(4, 4, record, record[2], "add ACL Action")

		case "aclgroup":
			okFile = ValidateLine(4, 4, record, record[2], "add ACL Group")

		case "aclmenu":
			okFile = ValidateLine(4, 4, record, record[2], "add ACL Menu")

		case "aclresource":
			okFile = ValidateLine(4, 4, record, record[2], "add ACL Resource")

		case "brokercfg":
			okFile = ValidateLine(4, 4, record, record[2], "add broker CFG")

		case "brokerinput":
			okFile = ValidateLine(5, 5, record, record[2], "add broker input")

		case "brokerlogger":
			okFile = ValidateLine(5, 5, record, record[2], "add broker logger")

		case "brokeroutput":
			okFile = ValidateLine(5, 5, record, record[2], "add broker output")

		case "categoryhost":
			okFile = ValidateLine(4, 4, record, record[2], "add category host")

		case "categoryservice":
			okFile = ValidateLine(4, 4, record, record[2], "add category service")

		case "centreonproxy":
			okFile = ValidateLine(6, 6, record, record[2], "add centreonProxy")

		case "command":
			okFile = ValidateLine(5, 5, record, record[2], "add command")

		case "contact":
			okFile = ValidateLine(6, 7, record, record[2], "add contact")

		case "dependency":
			okFile = ValidateLine(6, 6, record, record[2], "add dependency")

		case "enginecfg":
			okFile = ValidateLine(5, 5, record, record[2], "add engine cfg")

		case "groupcontact":
			okFile = ValidateLine(4, 4, record, record[2], "add group contact")

		case "grouphost":
			okFile = ValidateLine(4, 4, record, record[2], "add group host")

		case "groupservice":
			okFile = ValidateLine(4, 4, record, record[2], "add group service")

		case "host":
			okFile = ValidateLine(7, 8, record, record[2], "add host")

		case "ldap":
			okFile = ValidateLine(4, 4, record, record[2], "add LDAP")

		case "poller":
			okFile = ValidateLine(7, 7, record, record[2], "add poller")

		case "resourcecfg":
			okFile = ValidateLine(6, 6, record, record[2], "add resourceCFG")

		case "service":
			okFile = ValidateLine(5, 5, record, record[3], "add service")

		case "templatecontact":
			okFile = ValidateLine(4, 4, record, record[2], "add template contact")

		case "templatehost":
			okFile = ValidateLine(4, 5, record, record[2], "add template host")

		case "templateservice":
			okFile = ValidateLine(4, 5, record, record[2], "add template service")

		case "timeperiod":
			okFile = ValidateLine(4, 4, record, record[2], "add time period")

		case "trap":
			okFile = ValidateLine(4, 4, record, record[2], "add trap")

		case "vendor":
			okFile = ValidateLine(4, 4, record, record[2], "add vendor")

		default:
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Printf("The object %v in the csv file is incorrect\n", record[1])
			okFile = false
		}

	case "modify":
		switch strings.ToLower(record[1]) {

		case "aclaction":
			okFile = ValidateLine(5, 5, record, record[2], "modify ACL Action "+record[3])

		case "aclgroup":
			okFile = ValidateLine(5, 5, record, record[2], "modify ACL Group "+record[3])

		case "aclmenu":
			okFile = ValidateLine(5, 5, record, record[2], "modify ACL Menu "+record[3])

		case "aclresource":
			okFile = ValidateLine(5, 5, record, record[2], "modify ACL Resource "+record[3])

		case "brokercfg":
			okFile = ValidateLine(5, 5, record, record[2], "modify broker CFG "+record[3])

		case "brokerinput":
			okFile = ValidateLine(6, 6, record, record[2], "modify broker input "+record[4])

		case "brokerlogger":
			okFile = ValidateLine(6, 6, record, record[2], "modify broker logger "+record[4])

		case "brokeroutput":
			okFile = ValidateLine(6, 6, record, record[2], "modify broker output "+record[4])

		case "categoryhost":
			okFile = ValidateLine(5, 5, record, record[2], "modify category host "+record[3])

		case "categoryservice":
			okFile = ValidateLine(5, 5, record, record[2], "modify category service "+record[3])

		case "command":
			okFile = ValidateLine(5, 5, record, record[2], "modify command "+record[3])

		case "contact":
			okFile = ValidateLine(5, 5, record, record[2], "modify contact "+record[3])

		case "dependency":
			okFile = ValidateLine(5, 5, record, record[2], "modify dependency "+record[3])

		case "enginecfg":
			okFile = ValidateLine(5, 5, record, record[2], "modify engine cfg "+record[3])

		case "groupcontact":
			okFile = ValidateLine(5, 5, record, record[2], "modify group contact "+record[3])

		case "grouphost":
			okFile = ValidateLine(5, 5, record, record[2], "modify group host "+record[3])

		case "groupservice":
			okFile = ValidateLine(5, 5, record, record[2], "modify group service "+record[3])

		case "host":
			okFile = ValidateLine(5, 5, record, record[2], "modify host "+record[3])

		case "ldap":
			okFile = ValidateLine(5, 5, record, record[2], "modify LDAP "+record[3])

		case "poller":
			okFile = ValidateLine(5, 5, record, record[2], "modify poller "+record[3])

		case "resourcecfg":
			okFile = ValidateLine(5, 5, record, record[2], "modify resourceCFG "+record[3])

		case "service":
			okFile = ValidateLine(6, 6, record, record[3], "modify service "+record[4])

		case "templatecontact":
			okFile = ValidateLine(5, 5, record, record[2], "modify template contact "+record[3])

		case "templatehost":
			okFile = ValidateLine(5, 5, record, record[2], "modify template host "+record[3])

		case "templateservice":
			okFile = ValidateLine(5, 5, record, record[2], "modify template service "+record[3])

		case "timeperiod":
			okFile = ValidateLine(5, 5, record, record[2], "modify time period "+record[3])

		case "trap":
			okFile = ValidateLine(5, 5, record, record[2], "modify trap "+record[3])

		case "vendor":
			okFile = ValidateLine(5, 5, record, record[2], "modify vendor "+record[3])

		default:
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Printf("The object %v in the csv file is incorrect\n", record[1])
			okFile = false
		}
	default:
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Printf("The type %v in the csv file is incorrect (add or modify)\n", record[0])
		okFile = false
	}
	if okFile == false {
		return false
	}

	return okFile
}

//ValidateLine permits to validate if the line is conform with action
func ValidateLine(minNum int, maxNum int, record []string, valName string, name string) bool {
	colorRed := colorMessage.GetColorRed()
	okFile := true
	//Verification that all arguments exist
	if len(record) < minNum {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Printf("It missing an argument for the "+name+" `%v` in the csv file\n", valName)
		okFile = false
	}

	//Verification that there are not many arguments
	if len(record) > maxNum {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Printf("There are too many arguments for the "+name+" `%v` in the csv file\n", valName)
		okFile = false
	}

	if minNum != maxNum {
		//Verification that the comma exists in the end of line
		if len(record) == 3 {
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Printf("It missing a comma or value at the end of line of the "+name+" `%v` in the csv file\n", valName)
			okFile = false
		}
	}

	return okFile
}

type pollerStruct struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type result struct {
	Pollers []pollerStruct `json:"result"`
}

//exportPoller permits to export the poller
func exportPoller(debugV bool) {
	pollers := getPollers(debugV)
	for _, p := range pollers {
		err := applyPoller.Apply(p, debugV)
		if err != nil {
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println(err.Error())
		}
	}

}

//getPollers permits to get all pollers in centreon
func getPollers(debugV bool) []string {
	var pollers []string

	err, body := request.GeneriqueCommandV1Post("show", "instance", "", "import apply", debugV, false, "")
	if err != nil {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var pollerList result
	json.Unmarshal(body, &pollerList)
	for _, p := range pollerList.Pollers {
		if p.Status == "1" {
			pollers = append(pollers, p.Name)
		}
	}
	return pollers
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringP("file", "f", "", "To define the file which contains the objects to be imported")
	importCmd.Flags().Bool("apply", false, "Export configuration of the poller")
	importCmd.Flags().Bool("DETAIL", false, "Details information for all addition or modification")
}
