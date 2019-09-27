package controller

import (
	"html/template"
	"io/ioutil"
	"os"
)

func PopulateTemplates() map[string]*template.Template{
	const basePath = "view"
	result := make(map[string]*template.Template)

	layout :=template.Must(template.ParseFiles(basePath+"/_base.html"))
	dir,err :=os.Open(basePath+"/content")
	if err !=nil{
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis,err :=dir.Readdir(-1)
	if err !=nil{
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _,file :=range fis{
		f,err :=os.Open(basePath+"/content/"+file.Name())
		if err !=nil{
			panic("Failed to open template'"+file.Name()+"'")
		}

		content,err :=ioutil.ReadAll(f)

		if err !=nil{
			panic("Failed to read content from file'"+file.Name()+"'")
		}

		f.Close()
		templ :=template.Must(layout.Clone())

		_,err = templ.Parse(string(content))

		if err !=nil{
			panic("Faield to parse contents of'"+file.Name()+"' as template")
		}

		result[file.Name()] = templ

	}

	return result
}
