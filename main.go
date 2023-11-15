package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type LongitudeLatitude struct {
	Longitude   float64
	Latitude    float64
	ExcludedIds []string
}

func main() {
	// You could probably get the filename from request.EntryName and appending .json.tmpl
	fileName := "shops-within-5m-radius.json.tmpl"

	fnMap := template.FuncMap{
		"join": func(excludedIds []string) string {
			jsonIds, err := json.Marshal(excludedIds)
			if err != nil {
				fmt.Println(err) // throw an appropriate error here and do not os.Exit! lol
				os.Exit(1)
			}
			return string(jsonIds)
		},
	}

	// Open template file
	searchQueryTemplate, err := template.New(fileName).Funcs(fnMap).ParseFiles(fileName)
	if err != nil {
		fmt.Printf("cannot parse template file: %v", err)
		os.Exit(1)

	} // Prepare data to push into parsed file, this will be your params
	ll := LongitudeLatitude{
		Longitude:   52.01,
		Latitude:    -37.63,
		ExcludedIds: []string{"one", "two", "three"},
	}

	// Execute template by pushing Longitude Latitude into the text
	parsedQuery := new(strings.Builder)
	err = searchQueryTemplate.Execute(parsedQuery, ll)
	if err != nil {
		fmt.Printf("problem executing temple with given data: %v", err)
		os.Exit(1)
	}

	fmt.Println(parsedQuery)
}
