package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

type function struct {
	Name      functionName
	Namespace namespace
	Content   *functionContent
}

type functionContent struct {
	Content bytes.Buffer
}

func newFunctionContent() *functionContent {
	return &functionContent{Content: bytes.Buffer{}}
}

func newFunction(n functionName, ns namespace) *function {
	return &function{
		Name:      n,
		Namespace: ns,
		Content:   &functionContent{},
	}
}

func (f *function) generate() {
	var namespace string
	if f.Namespace != "" {
		parts := strings.Split(string(f.Namespace), "/")
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

	writeFile(filepath.Join(functionFolder, string(f.Name)+".mcfunction"), f.Content.Content.String())
}
