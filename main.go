package main

import (
	"fmt"
	"strings"
	"text/template"
)

type LongitudeLatitude struct {
	Longitude float64
	Latitude  float64
}

func main() {
	// You could probably get the filename from request.EntryName and appending .json.tmpl
	fileName := "shops-within-5m-radius.json.tmpl"

	// Open template file
	searchQueryTemplate, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Errorf("cannot parse template file: %v", err)

	} // Prepare data to push into parsed file, this will be your params
	ll := LongitudeLatitude{
		Longitude: 52.01,
		Latitude:  -37.63,
	}

	// Execute template by pushing LongitudeLatitude into the text
	parsedQuery := new(strings.Builder)
	err = searchQueryTemplate.Execute(parsedQuery, ll)
	if err != nil {
		fmt.Errorf("problem executing temple with given data: %v", err)
	}

	fmt.Println(parsedQuery)
}
