package main

import "io/ioutil"
import "encoding/json"
import "log"
import "strings"

type Pie struct {
	data interface{}
}

func readAFileFullOfJson(name string) *Pie {
	data, _ := ioutil.ReadFile(name)
	dec := json.NewDecoder(strings.NewReader(string(data)))
	pie := Pie{}
	err := dec.Decode(&pie.data)
	if err != nil {
		log.Fatal(err)
	}
	return &pie
}
