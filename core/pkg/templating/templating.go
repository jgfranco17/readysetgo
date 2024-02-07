package templating

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/jgfranco17/readysetgo/core/pkg/models"
)

var data = models.TextFile{
	Title:    "My Go Project",
	Contents: "This is a README generated using Go templating with a .tpl file.",
}

func CreateTemplatedFile(filename string) {
	// Parse the README template from a .tpl file
	templateFile := fmt.Sprintf("%s.tpl", filename)
	templateFilepath := filepath.Join("resources", templateFile)
	tmpl, err := template.ParseFiles(templateFilepath)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Execute the template and write the result to the file
	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}

	println("README.md generated successfully.")
}
