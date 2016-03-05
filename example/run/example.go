package main


import (
	"slob"
)

type Person struct{
	Name string
	Age int
	PhoneNum string
}

func main() {
	obj := new(Person)	

	slob.SetGenParams("src/slob/example/gen","prefix","suffix","go")
	slob.Render("src/slob/example/tpl/test.tpl").
		Read(obj).
		Set("pkgName","main").
		Execute()
}


