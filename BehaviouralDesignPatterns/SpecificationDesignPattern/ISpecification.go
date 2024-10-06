package main

const (
	Equal            = "="
	LessThanEqual    = "<="
	GreaterThanEqual = ">="
)

type ISpecification interface {
	IsSatisfiedBy(f *File) bool
}

type NameSpecification struct {
	Name string
}

func (n *NameSpecification) IsSatisfiedBy(f *File) bool {
	return n.Name == f.GetName()
}

type SizeSpecification struct {
	Size     int
	Operator string
}

func (s *SizeSpecification) IsSatisfiedBy(f *File) bool {
	if s.Operator == Equal {
		return s.Size == f.GetSize()
	} else if s.Operator == LessThanEqual {
		return f.GetSize() <= s.Size
	} else if s.Operator == GreaterThanEqual {
		return f.GetSize() >= s.Size
	}
	return false
}

type ExtensionSpecification struct {
	Ext string
}

func (e *ExtensionSpecification) IsSatisfiedBy(f *File) bool {
	return e.Ext == f.GetExtension()
}

type AndSpecification struct {
	Specifications []ISpecification
}

func (a *AndSpecification) IsSatisfiedBy(f *File) bool {
	for _, spec := range a.Specifications {
		if !spec.IsSatisfiedBy(f) {
			return false
		}
	}
	return true
}

type OrSpecification struct {
	Specifications []ISpecification
}

func (o *OrSpecification) IsSatisfiedBy(f *File) bool {
	for _, spec := range o.Specifications {
		if spec.IsSatisfiedBy(f) {
			return true
		}
	}
	return false
}
