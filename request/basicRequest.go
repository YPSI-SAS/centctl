/*MIT License

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
package request

import (
	"centctl/colorMessage"
	"centctl/debug"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func GeneriqueCommandV1Post(action string, object string, values string, command string, debugV bool, apply bool, pollerName string) (error, []byte) {
	requestBody, err := CreateBodyRequest(action, object, values)
	if err != nil {
		return err, []byte{}
	}

	//Recovery of the response body
	urlCentreon := os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi"
	client := NewClientV1(os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi")
	statusCode, body, err := client.CentreonCLAPI(requestBody)

	//If flag debug, print informations about the request API
	if debugV {
		debug.Show(command, string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err, []byte{}
	}

	if apply {
		//Export the poller configuration
		client = NewClientV1(os.Getenv("URL") + "/api/index.php?action=action&object=centreon_clapi")
		err = client.ExportConf(pollerName, debugV)
		if err != nil {
			return err, []byte{}
		}
	}
	return nil, body
}

func GeneriqueCommandV1Get(urlCentreon string, command string, debugV bool) (error, []byte) {
	client := NewClientV1(urlCentreon)
	statusCode, body, err := client.Get()

	//If flag debug, print informations about the request API
	if debugV {
		debug.Show(command, "", urlCentreon, statusCode, body)
	}
	if err != nil {
		return err, []byte{}
	}

	return nil, body
}

func GeneriqueCommandV2Get(urlCentreon string, command string, debugV bool) (error, []byte) {
	colorRed := colorMessage.GetColorRed()

	client := NewClientV2(urlCentreon)
	statusCode, body, err := client.Get()

	//If flag debug, print informations about the request API
	if debugV {
		debug.Show(command, "", urlCentreon, statusCode, body)
	}
	if err != nil {
		return err, []byte{}
	}
	//Verification with the response body
	if statusCode != 200 && !strings.Contains(command, "show") && !strings.Contains(command, "list") {
		var raw map[string]interface{}
		err = json.Unmarshal(body, &raw)
		if err != nil {
			// handle err
		}
		_, ok := raw["code"]
		if ok {
			message, _ := raw["message"]
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println(message)
			os.Exit(1)
		}
	}

	return nil, body
}

func GeneriqueCommandV2Post(urlCentreon string, requestBody []byte, command string, debugV bool) (error, []byte) {
	colorRed := colorMessage.GetColorRed()

	client := NewClientV2(urlCentreon)
	statusCode, body, err := client.Post(requestBody)

	//If flag debug, print informations about the request API
	if debugV {
		debug.Show(command, string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err, []byte{}
	}
	//Verification with the response body
	if statusCode != 200 {
		var raw map[string]interface{}
		err = json.Unmarshal(body, &raw)
		if err != nil {
			// handle err
		}
		_, ok := raw["code"]
		if ok {
			message, _ := raw["message"]
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println(message)
			os.Exit(1)
		}
	}

	return nil, body
}

func GeneriqueCommandV2Put(urlCentreon string, requestBody []byte, command string, debugV bool) error {
	colorRed := colorMessage.GetColorRed()

	client := NewClientV2(urlCentreon)
	statusCode, body, err := client.Put(requestBody)

	//If flag debug, print informations about the request API
	if debugV {
		debug.Show(command, string(requestBody), urlCentreon, statusCode, body)
	}
	if err != nil {
		return err
	}
	//Verification with the response body
	if statusCode != 200 {
		var raw map[string]interface{}
		err = json.Unmarshal(body, &raw)
		if err != nil {
			// handle err
		}
		_, ok := raw["code"]
		if ok {
			message, _ := raw["message"]
			fmt.Printf(colorRed, "ERROR: ")
			fmt.Println(message)
			os.Exit(1)
		}
	}

	return nil
}

func Add(action string, object string, values string, command string, name string, debugV bool, isImport bool, apply bool, pollerName string, successMessage string) error {
	colorRed := colorMessage.GetColorRed()
	colorGreen := colorMessage.GetColorGreen()

	err, body := GeneriqueCommandV1Post(action, object, values, command, debugV, apply, pollerName)
	if err != nil {
		return err
	}

	//Verification with the response body that the object was created out
	if string(body) != "{\"result\":[]}" && !isImport {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(string(body))
		os.Exit(1)
	} else if string(body) != "{\"result\":[]}" && isImport {
		return fmt.Errorf("%s. "+object+"'s name %v.", string(body), name)
	}

	fmt.Printf(colorGreen, "INFO: ")
	if successMessage == "" {
		fmt.Printf("The "+object+" %v is created\n", name)
	} else {
		fmt.Println(successMessage)
	}

	return nil
}

func Delete(action string, object string, values string, command string, name string, debugV bool, apply bool, pollerName string) error {
	colorRed := colorMessage.GetColorRed()
	colorGreen := colorMessage.GetColorGreen()

	err, body := GeneriqueCommandV1Post(action, object, values, command, debugV, apply, pollerName)
	if err != nil {
		return err
	}

	//Verification with the response body that the object was deleted out
	if string(body) != "{\"result\":[]}" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(string(body))
		os.Exit(1)
	}

	fmt.Printf(colorGreen, "INFO: ")
	fmt.Printf("The "+object+" %v is deleted\n", name)

	return nil
}

func Modify(action string, object string, values string, command string, name string, parameter string, detail bool, debugV bool, apply bool, pollerName string, isImport bool) error {
	colorRed := colorMessage.GetColorRed()
	colorGreen := colorMessage.GetColorGreen()

	err, body := GeneriqueCommandV1Post(action, object, values, command, debugV, apply, pollerName)
	if err != nil {
		return err
	}
	//Verification with the response body that the object was modified out
	if string(body) != "{\"result\":[]}" && !isImport {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(string(body))
		os.Exit(1)
	} else if string(body) != "{\"result\":[]}" && isImport {
		return fmt.Errorf("%s. "+object+"'s name %v modified %v.", string(body), name, parameter)
	}

	if detail {
		fmt.Printf(colorGreen, "INFO: ")
		fmt.Printf("The parameter %v of the %v %v is modified\n", parameter, object, name)
	}

	return nil
}
