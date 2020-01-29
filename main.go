package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		println("Must provide exactly one argument - the MAC address")
		os.Exit(1)
	}
	apiKey := os.Getenv("MACADDRESS_IO_API_KEY")
	macAddress := os.Args[1]

	url := fmt.Sprintf("https://api.macaddress.io/v1?apiKey=%s&output=json&search=%s", apiKey, macAddress)
	resp, err := http.Get(url)
	if err != nil {
		println("There was an error retrieving data from an upstream system")
		os.Exit(2)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	answer := extractCompanyName(body)
	print(answer)
}

func extractCompanyName(body []byte) string {
	var macResponse map[string]interface{}
	json.Unmarshal(body, &macResponse)
	return macResponse["vendorDetails"].(map[string]interface{})["companyName"].(string)
}
