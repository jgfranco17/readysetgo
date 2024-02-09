package models

import "fmt"

type Project struct {
	Name   string
	Author string
	Type   string
}

type Subdirectory struct {
	Name  string
	Files []File
}

type File struct {
	Name      string
	Extension string
}

func (f *File) FullName() string {
	return fmt.Sprintf("%s.%s", f.Name, f.Extension)
}
