package main

import "strings"

type File struct {
	Name        string
	SizeInBytes int
	Extension   string
}

func NewFile(name string, bytes int) *File {
	fd := strings.Split(name, ".")
	ext := ""
	if len(fd) > 1 {
		ext = fd[1]
	}
	return &File{
		Name: fd[0], SizeInBytes: bytes, Extension: ext,
	}
}

func (f *File) GetName() string {
	return f.Name
}

func (f *File) GetSize() int {
	return f.SizeInBytes
}

func (f *File) GetExtension() string {
	return f.Extension
}

func SearchFiles(files []*File, criteria ISpecification) []*File {
	i := 0
	var res []*File
	for i < len(files) {
		if criteria.IsSatisfiedBy(files[i]) {
			res = append(res, files[i])
		}
		i++
	}
	return res
}
