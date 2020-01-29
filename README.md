# Find Company by MAC Address

This command line utility finds the company name associated with a MAC address and writes it to the console output.
You will need your https://macaddress.io api key to run

## Get module
`go get github.com/kamal2311/macaddressio`

## Run 
`export MACADDRESSIO_API_KEY=<Your api key>`

`go run github.com/kamal2311/macaddressio <Mac Address>`

e.g.

`go run github.com/kamal2311/macaddressio 44:38:39:ff:ef:57`

Sample output

```Cumulus Networks, Inc```

## Test
`go test github.com/kamal2311/macaddressio`