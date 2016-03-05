package slob

import (
	"fmt"
	"os/exec"
	"text/template"
)

type render struct {
	TmplPath    string
	GenFileName string
	Data        template.FuncMap
	input
}

func Render(tplPath string) *render {
	r := new(render)
	r.TmplPath = tplPath
	r.Data = template.FuncMap{}
	r.input = slob_input
	return r
}

func (r *render) SetRender(inp input) *render {
	r.input = inp
	return r
}

func (r *render) SetFileName(name string) *render {
	r.GenFileName = name
	return r
}

func (r *render) Set(name string, param interface{}) *render {
	r.Data[name] = param
	return r
}

func (r *render) execute() {
	tmpl, err := template.ParseGlob(r.TmplPath)
	if err != nil {
		panic(err)
	}
	fd, err := GenFileHandle(r.GenFileName)
	defer fd.Close()
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(fd, r.Data)
	if err != nil {
		panic(err)
	}
}

func (r *render) Read(obj ...interface{}) *render {
	r.input.Read(obj...)
	return r
}

func (r *render) Execute() {
	r.input.Render(r)
	gofmt()
}

func gofmt() {
	cmd := exec.Command("gofmt", "-w", fmt.Sprintf("%s/", GetGenDir()))
	err := cmd.Run()
	if err != nil {
		fmt.Println("gofmt err , please run : gofmt -w ", fmt.Sprintf("%s/", GetGenDir()))
	}
}
