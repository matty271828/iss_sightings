package main

import (
	"encoding/json"
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

func contact_api(lat, long, minElevation, hours string) string {
	// Contact API
	response, err := http.Get("https://api.g7vrd.co.uk/v1/satellite-passes/25544/" + lat + "/" + long + ".json?minelevation=" + minElevation + "&hours=" + hours + "")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Retrieve data from body of http response
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parses the JSON-encoded data and stores the result
	m := map[string]interface{}{}
	if err := json.Unmarshal(responseData, &m); err != nil {
		panic(err)
	}

	// Use type assertion to loop over []interface{} and retrieve the value
	for _, pass := range m["passes"].([]interface{}) {
		fmt.Println(pass)
	}

	return (string(responseData))
}
