package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {

	APIKEY := os.Getenv("MACADDRESSIO_API_KEY")
	if APIKEY == "" {
		log.Fatal("APIKEY is missing - Please set MACADDRESSIO_API_KEY environment variable.")
	}
	if len(os.Args) != 2 {
		log.Fatal("Must provide exactly one argument - the MAC address")
	}
	MACADDRESS := os.Args[1]
	if !isValidMACAddress(MACADDRESS) {
		log.Fatal("MAC address provided was invalid. A sample valid MAC address looks like 44:38:39:ff:ef:57")
	}

	URL := fmt.Sprintf("https://api.macaddress.io/v1?apiKey=%s&output=json&search=%s", APIKEY, MACADDRESS)

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	name, err := extractCompanyName(body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name)
}

func extractCompanyName(body []byte) (companyName string, err error) {
	var macResponse map[string]interface{}
	err = json.Unmarshal([]byte(body), &macResponse)
	if err != nil {
		return "", err
	}

	if macResponse["error"] != nil {
		return "", errors.New(macResponse["error"].(string))
	}

	return macResponse["vendorDetails"].(map[string]interface{})["companyName"].(string), nil
}

func isValidMACAddress(mac string) bool {
	validMAC := regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`)
	return validMAC.Match([]byte(mac))
}
