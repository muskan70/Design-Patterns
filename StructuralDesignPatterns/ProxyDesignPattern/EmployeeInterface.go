package main

type IEmployee interface {
	Create(user string, name string)
	Get(user string, id int)
	Delete(user string, id int)
}
