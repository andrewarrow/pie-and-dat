package main

import "io/ioutil"

func readAFileFullOfJson(name string) string {
	data, _ := ioutil.ReadFile(name)
	return string(data)
}
