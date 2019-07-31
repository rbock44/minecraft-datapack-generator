package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type function struct {
	Name      string
	Namespace string
	Content   string
}

func newFunction(name string, namespace string, content string) *function {
	return &function{
		Name:      name,
		Namespace: namespace,
		Content:   content,
	}
}

func (f *function) generate() {
	var namespace string
	if f.Namespace != "" {
		parts := strings.Split(f.Namespace, "/")
		namespace = filepath.Join(parts...)
	}
	functionFolder := filepath.Join(
		"generated", "data",
		dp.FunctionRoot,
		"functions",
		namespace)

	err := os.MkdirAll(functionFolder, 0755)
	if err != nil {
		fatal("Cannot create function directory <%s>", err.Error())
	}

	ioutil.WriteFile(filepath.Join(functionFolder, f.Name+".mcfunction"), []byte(f.Content), 0644)
}
