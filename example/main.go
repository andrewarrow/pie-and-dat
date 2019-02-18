package main

import "fmt"
import "os"

func main() {
	args := os.Args
	//pwd, _ := os.Getwd()
	if len(args) == 1 {
		fmt.Println("usage: piedat <option>")
		return
	}
	option := args[1]

	if option == "dir" {
		LoadADirFullOfJsonFiles(args[2])
	} else if option == "http" {
	} else if option == "file" {
	}
}
