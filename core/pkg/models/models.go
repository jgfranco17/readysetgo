package models

type Project struct {
	Name   string
	Author string
	Type   string
}

type Subdirectory struct {
	Name  string
	Files []string
}

type TextFile struct {
	Title    string
	Contents string
}
