package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	filename := args[1]

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalln("file not exist, " + filename)
	}
	search := args[2]
	replacePart := args[3:]
	replace := strings.Join(replacePart, " ")
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
