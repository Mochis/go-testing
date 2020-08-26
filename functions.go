package main

import "io/ioutil"

/**
 Function for checking errors
 */
func check(e error) {
	panic(e)
}

func main() {
	_, err := ioutil.ReadFile("not_existing_file.log")
	if err != nil {
		check(err)
	}
}
