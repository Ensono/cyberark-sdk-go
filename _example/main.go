package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	os.LookupEnv("")
	url := "https://pam-int-master.nonprod.idam.cloud.iairgroup.com/PasswordVault/API/Auth/CyberArk/Logon"
	method := "POST"

	payload := strings.NewReader(`{
	"username": "Administrator",
	"password": "CmHpWaMCHRrj9!c",
	//"newPassword": "<optional>",
	"concurrentSession": "true"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
