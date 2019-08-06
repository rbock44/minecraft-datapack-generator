package main

import (
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
	mkdir(tagFolder)
	writeFile(filepath.Join(tagFolder, t.Name+".json"), t.Content)
}
