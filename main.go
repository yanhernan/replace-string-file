package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	filename := args[0]
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalln("file not exist, " + filename)
	}

	search := args[1]
	replace := args[1]
	data, e := ioutil.ReadFile(filename)
	if e != nil {
		log.Fatalln(e)
	}
	res := strings.ReplaceAll(string(data), search, replace)
	err := ioutil.WriteFile(filename, []byte(res), 0)
	if err != nil {
		panic(err)
	}
}
