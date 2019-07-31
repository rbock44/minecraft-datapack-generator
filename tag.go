package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type tag struct {
	Name    string
	Content string
}

func newTag(name, content string) *tag {
	return &tag{Name: name, Content: content}
}

func (t *tag) generate() {
	tagFolder := filepath.Join("generated", "data", "minecraft", "tags", "functions")
	err := os.MkdirAll(tagFolder, 0755)
	if err != nil {
		fatal("Cannot create tag folder <%s>", err.Error())
	}
	err = ioutil.WriteFile(filepath.Join(tagFolder, t.Name+".json"), []byte(t.Content), 0644)
	if err != nil {
		fatal("Cannot write tag file <%s>", err.Error())
	}
}
