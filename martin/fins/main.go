package main

import (
	"html/template"
	"io"
	"log"
	"os"
)

const tpl = `


<!DOCTYPE html>

{{ define "myHeader" }}
<html>
    <header>
        {{ . }}
		</header>
{{ end }}

{{ define "myFooter" }}
foot
{{ end }}


`

func main() {
	createFile("index.html")
	f, err := os.OpenFile("index.html", os.O_RDWR, 0777)
	io.WriteString(f, tpl)
	f.Close()
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Printf("Error Parsing file %v, error: %v \n", t, err)
	} else {
		log.Printf("Parsed template %v.\n", t.Name())
	}
	myFile, _ := os.Create("after.html")
	t.ExecuteTemplate(myFile, "myHeader", ":)")
	t.ExecuteTemplate(myFile, "myFooter", " ")
	myFile.Close()

}

func createFile(s string) {
	_, err := os.Create(s)
	if err != nil {
		log.Printf("Error creating file %s, error: %v \n", s, err)
	} else {
		log.Printf("Created file %v.\n", s)
	}

}
