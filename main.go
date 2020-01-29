package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	APIKEY := os.Getenv("MACADDRESSIO_API_KEY")
	if APIKEY == "" {
		println("APIKEY is missing - report an issue")
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Println("Must provide exactly one argument - the MAC address")
		os.Exit(1)
	}
	MACADDRESS := os.Args[1]

	URL := fmt.Sprintf("https://api.macaddress.io/v1?apiKey=%s&output=json&search=%s", APIKEY, MACADDRESS)

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("There was an error retrieving data from an upstream system")
		os.Exit(2)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	name := extractCompanyName(body)
	fmt.Println(name)
}

func extractCompanyName(body []byte) (companyName string) {
	var macResponse map[string]interface{}
	err := json.Unmarshal([]byte(body), &macResponse)
	if err != nil {
		fmt.Println("API response could not be parsed")
		return ""
	}
	return macResponse["vendorDetails"].(map[string]interface{})["companyName"].(string)
}
