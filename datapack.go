package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type datapack struct {
	Name         string
	FunctionRoot string
	Version      int
	Description  string
	TargetPath   string

	Functions []*function
	Tags      []*tag
}

type mcmeta struct {
	Template string
}

func (d *datapack) generate() {
	d.removeFolder()
	d.createMCMeta()
	for _, f := range d.Functions {
		info("generate function <%s>", f.Name)
		f.generate()
	}
	for _, t := range d.Tags {
		info("generate tag <%s>", t.Name)
		t.generate()
	}

	d.pack()
}

func (d *datapack) targetPath() string {
	return "generated"
}

func (d *datapack) functionsPath() string {
	return filepath.Join(d.targetPath(), d.FunctionRoot)
}

func (d *datapack) addFunction(name, namespace, content string) {
	f := newFunction(name, namespace, content)
	d.Functions = append(d.Functions, f)
}

func (d *datapack) addTag(name, content string) {
	t := newTag(name, content)
	d.Tags = append(d.Tags, t)
}

func (d *datapack) removeFolder() {
	err := os.RemoveAll("generated")
	if err != nil {
		fatal("Cannot clean datapack folder <%s>", err.Error())
	}
}

func (d *datapack) createMCMeta() {
	data := `{
  "pack": {
    "pack_format": {{ .Version }},
    "description": "{{ .Description }}"
  }
}`
	err := os.Mkdir("generated", 0755)
	if err != nil {
		fatal("Cannot create generated folder <%s>", err.Error())
	}

	file, err := os.Create(filepath.Join("generated", "pack.mcmeta"))
	if err != nil {
		fatal("Cannot create pack.mcmeta <%s>", err.Error())
	}
	defer file.Close()
	t, err := template.New("pack.mcmeta").Parse(data)
	if err != nil {
		fatal("Cannot parse template pack.mcmeta <%s>", err.Error())
	}
	err = t.Execute(file, d)
	if err != nil {
		fatal("Cannot execute template pack.mcmeta <%s>", err.Error())
	}
	fmt.Printf("pack.mcmeta generated\n")
}

func (d *datapack) pack() {
	recursiveZip(filepath.Join("generated"), fmt.Sprintf("%s-v%d.zip", d.Name, d.Version))
}
