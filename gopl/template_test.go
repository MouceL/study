package gopl

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

type person struct {
	Student []Student
}

type Student struct {
	Name string
	City string
	Age int
}

var tmpl = `{{range .Student}}
{{ .Name }} comes from {{ .City | printf "%.64s" }},and {{ .Age }}
{{end}}`

func TestTemplate(t *testing.T){

	test,err := template.New("test").Parse(tmpl)
	if err!=nil{
		fmt.Println(err.Error())
	}
	person := &person{}
	array := make([]Student,0,10)
	for i:=0;i<5;i++{
		array = append(array,Student{
			Name: fmt.Sprintf("name_%d",i),
			City: "bjbjbjbj",
			Age:  i,
		})
	}

	person.Student = array
	buffer := &bytes.Buffer{}
	err = test.Execute(buffer,person)
	if err!=nil{
		fmt.Printf(err.Error())
	}
	fmt.Println(buffer.String())



}