package main

import "io/ioutil"
import "encoding/json"
import "log"
import "strings"
import "github.com/andrewarrow/pie-and-dat"

func readAFileFullOfJson(name string) *pie.Pie {
	data, _ := ioutil.ReadFile(name)
	dec := json.NewDecoder(strings.NewReader(string(data)))
	p := pie.Pie{}
	err := dec.Decode(&p.Data)
	if err != nil {
		log.Fatal(err)
	}
	return &p
}
