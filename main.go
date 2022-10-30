package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// Parameters
	var lat string = "51.45"
	var long string = "-2.5833"
	var minElevation string = "70"
	var hours string = "72"

	// Retrieve upcoming passes
	contact_api(lat, long, minElevation, hours)
}

func contact_api(lat, long, minElevation, hours string) {
	response, err := http.Get("https://api.g7vrd.co.uk/v1/satellite-passes/25544/" + lat + "/" + long + ".json?minelevation=" + minElevation + "&hours=" + hours + "")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
