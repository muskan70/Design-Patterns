package main

import "fmt"

func main() {
	utl := NewFile("hello.txt", 3)
	utl2 := NewFile("picture.png", 7)
	utl3 := NewFile("response.json", 2)

	spec1 := &NameSpecification{Name: "hello"}
	spec2 := &ExtensionSpecification{Ext: "txt"}
	spec3 := &SizeSpecification{Size: 5, Operator: GreaterThanEqual}
	spec := &OrSpecification{Specifications: []ISpecification{spec3, &AndSpecification{Specifications: []ISpecification{spec1, spec2}}}}

	files := SearchFiles([]*File{utl, utl2, utl3}, spec)

	for i := range files {
		fmt.Println(files[i])
	}
}
