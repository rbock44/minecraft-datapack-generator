package main

import (
	"io/ioutil"
	"os"
)

func mkdir(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fatal("Cannot create folder <%s>", err.Error())
	}
}

func writeFile(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fatal("Cannot write file <%s> error <%s>", path, err.Error())
	}
}
