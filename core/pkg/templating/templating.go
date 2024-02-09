package templating

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/jgfranco17/readysetgo/core/pkg/models"
)

func CreateTemplatedFile(filename string, data *models.File) (models.File, error) {
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

	fmt.Printf("%s generated successfully.", filename)
	return *data, nil
}
