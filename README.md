# slob
golang template generate tool

## example
first you need create own tpl 

	// test.tpl
	package {{.pkgName}}
	
	import(
		"fmt"
	)
	
	func main(){
		fmt.Println("{{.struct.Name}}")
	
		{{range .struct.Fields}}
			fmt.Println("the ","{{.Name}}"," is ", "{{.Type}}")	
		{{end}}
	}

run this main.go, you can gen the file

	// main.go
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
		
		// set gen into dir, gen_file prefix suffix and gen_file's type
		slob.SetGenParams("src/slob/example/gen","prefix","suffix","go")

		// set tpl and some params 
		slob.Render("src/slob/example/tpl/test.tpl").
			Read(obj).
			Set("pkgName","main").
			Execute()
	}
	
the gen file

	// prefix_person_suffix.go  
	package main
	
	import (
		"fmt"
	)
	
	func main() {
		fmt.Println("Person")
	
		fmt.Println("the ", "Name", " is ", "string")
	
		fmt.Println("the ", "Age", " is ", "int")
	
		fmt.Println("the ", "PhoneNum", " is ", "string")
	
	}

## Support
If you do have a contribution for the package feel free to put up a Pull Request or open Issue.
