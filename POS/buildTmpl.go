package main

import (
	"html/template"
	"log"
	"os"
)

type myStruct struct {
	Name      string
	PositionX int
	PositionY int
}

const myData = "hello"

const myFooter template.HTML = `
		<div><p>

			<ul>
			<li>copyright</li>
			<li>sitemap</li>
			<li>contact</li>
			<li>to top</li>
			</ul>

		</p>
		</div>
`

func main() {

	myDataStruct := myStruct{
		Name:      "Martin",
		PositionX: 100,
		PositionY: 200,
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Printf("Error parsing file %v: %v/n", t.Name(), err)
	}

	f, _ := os.Create("edit.html")
	t.ExecuteTemplate(os.Stdout, "POSHeader", myData)
	t.ExecuteTemplate(os.Stdout, "POSBody", myDataStruct)
	t.ExecuteTemplate(f, "POSFooter", myFooter)

}
