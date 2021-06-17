package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// CheckError quits the execution if there is an error and logs the error message
func CheckError(err error, message string) {
	if err != nil {
		log.Fatalln(fmt.Sprintf("ERROR - %s: %v", message, err))
	}
}

// FloatToString transforms the given float64 input to string
func FloatToString(input float64) string {
	return fmt.Sprintf("%.1f", input)
}

// ResourceToStringMap transforms a given resource struct into a string map
func ResourceToStringMap(SecurityContext interface{}) map[string]interface{} {
	var ResourceMap map[string]interface{}
	ResourceJson, err := json.Marshal(SecurityContext)
	CheckError(err, "Could not encode resource struct to JSON")

	err = json.Unmarshal(ResourceJson, &ResourceMap)
	CheckError(err, "Could not unmarshal resource JSON into map")

	return ResourceMap
}
