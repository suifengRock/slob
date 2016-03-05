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