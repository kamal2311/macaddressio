# Find Company by MAC Address

This command line utility finds the company name associated with a MAC address and writes it to the console output.

*Important*: You will need an API key obtained from https://macaddress.io to run

## Simply run the utility - Run from an existing [Dockerhub image](https://hub.docker.com/repository/docker/kamal2311/macaddressio)

`docker run -eMACADDRESSIO_API_KEY=<your api key> kamal2311/macaddressio:v3 44:38:39:ff:ef:57`

## Get module
`go get github.com/kamal2311/macaddressio`

## Run 
`export MACADDRESSIO_API_KEY=<Your api key>`

`go run github.com/kamal2311/macaddressio <Mac Address>`

e.g.

`go run github.com/kamal2311/macaddressio 44:38:39:ff:ef:57`

Sample output

`Cumulus Networks, Inc`

## Test
`go test github.com/kamal2311/macaddressio`

## Build your own docker image

`docker build -t <your_tag_name>:<version> .`

## Run the docker image

`docker run -eMACADDRESSIO_API_KEY=<your api key> <your_tag_name>:<version> <MAC_ADDRESS_TO_SEARCH>`

## For smaller size docker image 800MB reduced to 8MB!

Build the application locally with the following command. This will compile the go application with statically linked libraries.

`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo github.com/kamal2311/macaddressio .`

Then build the docker image using the file Dockerfile.optimized

`docker build -t <your_tag_name>:<version> -f Dockerfile.optimized .`

Run the image as before

`docker run -eMACADDRESSIO_API_KEY=<your api key> <your_tag_name>:<version> <MAC_ADDRESS_TO_SEARCH>`







