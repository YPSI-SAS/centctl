package request

import (
	"bytes"
	"centctl/colorMessage"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

//Token represents the generate token by the authentification
type Token struct {
	Token string `json:"authToken"`
}

//AuthentificationV1 allow the authentification at the server specified with API v1
func AuthentificationV1(urlServer string, login string, password string, insecure bool) (string, error) {
	colorRed := colorMessage.GetColorRed()
	urlCentreon := urlServer + "/api/index.php?action=authenticate"
	if insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	resp, err := http.PostForm(urlCentreon,
		url.Values{"username": {login}, "password": {password}})
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 401 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(resp.Status)
		os.Exit(1)
	} else if resp.StatusCode == 407 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(resp.Status)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if string(body) == "\"Bad credentials\"" {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println("Login or password incorrect")
		os.Exit(1)
	}
	t := &Token{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return "", err
	}
	return t.Token, err
}

//AuthentificationV2 allow the authentification at the server specified with API v2
func AuthentificationV2(urlServer string, login string, password string, insecure bool) (string, error) {
	colorRed := colorMessage.GetColorRed()
	request := make(map[string]interface{})
	request["security"] = map[string]interface{}{
		"credentials": map[string]interface{}{
			"login":    login,
			"password": password,
		},
	}

	// Marshal the map into a JSON string.
	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	urlCentreon := urlServer + "/api/beta/login"

	if insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	resp, err := http.Post(urlCentreon, "application/json",
		bytes.NewBuffer(requestBody))

	if err != nil {
		return "", err
	}

	if resp.StatusCode == 401 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(resp.Status)
		os.Exit(1)
	} else if resp.StatusCode == 407 {
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(resp.Status)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	var raw map[string]interface{}
	err = json.Unmarshal(body, &raw)
	if err != nil {
		return "", err
	}
	_, ok := raw["code"]
	if ok {
		message, _ := raw["message"]
		fmt.Printf(colorRed, "ERROR: ")
		fmt.Println(message)
		os.Exit(1)
	}
	token, _ := raw["security"]
	token = token.(interface{}).(map[string]interface{})["token"]
	tokenVal := fmt.Sprintf("%v", token)

	return tokenVal, err
}
