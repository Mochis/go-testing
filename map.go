package main

import "fmt"

func main() {
	var countryMap map[string]string
	countryMap = make(map[string]string) // create hashmap
	countryMap["SP"] = "Spain"
	countryMap["IT"] = "Italy"

	// check key exists
	_, exists := countryMap["PT"]
	if !exists {
		fmt.Println("Key PT does not exists")
	}
	if countryMap["GE"] == "" {
		fmt.Println("Key GE does not exists")
	}

	// delete key

	for key, value := range countryMap {
		fmt.Println(key, " -> ", value)
	}
}