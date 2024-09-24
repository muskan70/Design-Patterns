package main

import "fmt"

type File struct {
	Name string
}

func (f *File) Clone() Prototype {
	return &File{Name: f.Name + "_clone"}
}

func (f *File) Print(indent string) {
	fmt.Println(indent + f.Name)
}

type Folder struct {
	Name     string
	Children []Prototype
}

func (f *Folder) Clone() Prototype {
	cf := &Folder{Name: f.Name + "_clone"}

	var tempChildren []Prototype
	for _, i := range f.Children {
		tempChildren = append(tempChildren, i.Clone())
	}
	cf.Children = tempChildren
	return cf
}

func (f *Folder) Print(indent string) {
	fmt.Println(indent + f.Name)
	for _, i := range f.Children {
		i.Print(indent + indent)
	}
}
