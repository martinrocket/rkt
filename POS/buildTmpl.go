package main

import (
	"html/template"
	"log"
	"os"
)

// myStruct is an example struct to be written into a template.
type MyButton struct {
	Name      string
	PositionX int
	PositionY int
	Action    string
}

// myData is an example of a string to be written into a template.
const myData = "hello"

// myFooter is an exmple of a string written into a template with template.HTML formatting.
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
	// myDataStruct uses myStuct to build some test data for the template "Index.html"
	Button1 := MyButton{
		Name:      "Pepsi",
		PositionX: 100,
		PositionY: 200,
		Action:    "Sell Item",
	}

	Button2 := MyButton{
		Name:      "Sprite",
		PositionX: 300,
		PositionY: 200,
		Action:    "Sell Item",
	}

	ButtonGrp1 := []MyButton{Button1, Button2}

	// t is a template of index.html
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Printf("Error parsing file %v: %v/n", t.Name(), err)
	}

	// f is a new file in the local directory called edit.html where the t template will be written to for persistance.
	f, _ := os.Create("edit.html")

	t.ExecuteTemplate(os.Stdout, "POSHeader", myData)
	t.ExecuteTemplate(os.Stdout, "POSBody", ButtonGrp1)
	t.ExecuteTemplate(f, "POSFooter", myFooter)

}
