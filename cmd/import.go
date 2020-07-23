/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import the objects contained in the csv file",
	Long:  `Import the objects contained in the csv file. It import contacts, hosts or services.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		debugV, _ := cmd.Flags().GetBool("DEBUG")
		apply, _ := cmd.Flags().GetBool("apply")
		ImportCSV(file, debugV, apply)
	},
}

//ImportCSV permits to import objects contains in a CSV file
func ImportCSV(file string, debugV bool, apply bool) {
	//Verification that the file has a correct extension
	if !strings.Contains(file, ".csv") {
		fmt.Println("The extension of the file must be .csv")
		os.Exit(1)
	}
	if ValidateCSV(file) {
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
			//switch on the object
			switch record[0] {
			case "service":
				//Recovery the values of the service
				hostName := record[1]
				description := record[2]
				template := record[3]

				//Creation of service
				err = AddService(hostName, description, template, debugV, apply)

			case "host":
				//Recovery the values of the host
				name := record[1]
				alias := record[2]
				IPaddress := record[3]
				template := record[4]
				poller := record[5]
				hostGroup := record[6]

				//Creation of host
				err = AddHost(name, alias, IPaddress, template, poller, hostGroup, debugV, apply)

			case "contact":
				//Recovery the values of the contact
				name := record[1]
				alias := record[2]
				email := record[3]
				password := record[4]

				////Transformation of the admin record to bool
				admin := record[5]
				adminV := true
				if admin == "" {
					adminV = false
				}

				//Creation of contact
				err = AddContact(name, alias, email, password, adminV, debugV)
			}
		}
	}
}

//ValidateCSV permits to validate the configuration of the csv
func ValidateCSV(file string) bool {
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
		//switch on the object
		switch record[0] {
		case "service":

			//Verification that all arguments exist
			if len(record) < 4 {
				fmt.Printf("It missing an argument for the service `%v` in the csv file\n", record[2])
				okFile = false
			}

			//Verification that there are not many arguments
			if len(record) > 4 {
				fmt.Printf("There are too many arguments for the service `%v` in the csv file\n", record[2])
				okFile = false
			}

		case "host":

			//Verification that all arguments exist
			if len(record) < 6 {
				fmt.Printf("It missing an argument for the host `%v` in the csv file\n", record[1])
				okFile = false
			}

			//Verification that there are not many arguments
			if len(record) > 7 {
				fmt.Printf("There are too many arguments for the host `%v` in the csv file\n", record[1])
				okFile = false
			}

			//Verification that the comma exists in the end of line if hostGroup record not exists
			if len(record) == 6 {
				fmt.Printf("It missing a comma at the end of line of the host `%v` in the csv file\n", record[1])
				okFile = false
			}

		case "contact":

			//Verification that all arguments exist
			if len(record) < 5 {
				fmt.Printf("It missing an argument for the contact `%v` in the csv file\n", record[1])
				okFile = false
			}

			//Verification that there are not many arguments
			if len(record) > 6 {
				fmt.Printf("There are too many arguments for the contact `%v` in the csv file\n", record[1])
				okFile = false
			}

			//Verification that the comma exists in the end of line if admin record not exists
			if len(record) == 5 {
				fmt.Printf("It missing a comma at the end of line of the contact `%v` in the csv file\n", record[1])
				okFile = false
			}

		default:
			fmt.Printf("The object %v in the csv file is incorrect\n", record[0])
			okFile = false
		}

	}
	return okFile
}

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.Flags().StringP("file", "f", "", "To define the file which contains the objects to be imported")
	importCmd.MarkFlagRequired("file")
	importCmd.Flags().Bool("apply", false, "Export configuration of the poller")
}
