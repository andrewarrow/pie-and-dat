package main

import "fmt"
import "strings"
import "log"
import "io/ioutil"
import "github.com/andrewarrow/pie-and-dat"

var m = map[string]string{}
var i = 1

func processFileByCategory(name, category string) {
	if category == "17.json" {
		p := readAFileFullOfJson(name)
		pie.FindKeys(p.Data, "channelId", "channelTitle")
		for _, v := range pie.Pair {
			m[v["id"]] = v["name"]
		}
	}
	for k, v := range m {
		fmt.Println(i, k, v)
		i++
	}
}

func assignMeaning(name string, value bool, token, meaning string) {
	if value == false {
		return
	}
	if meaning == "category" {
		processFileByCategory(name, token)
	}
}
func LoadADirFullOfJsonFiles(path string) {
	fmt.Println("load dir", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		name := f.Name()
		tokens := strings.Split(name, "_")
		name = path + "/" + name
		for i, t := range tokens {
			assignMeaning(name, i == 0, t, "year")
			assignMeaning(name, i == 1, t, "month")
			assignMeaning(name, i == 2, t, "day")
			assignMeaning(name, i == 3, t, "hour")
			assignMeaning(name, i == 4, t, "minute")
			assignMeaning(name, i == 6, t, "category")
		}
	}

}
