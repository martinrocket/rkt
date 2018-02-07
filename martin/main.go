package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {

	files := []string{
		"myfile.html", "footer.html", "dilly.html",
	}

	t, err := template.ParseFiles(files[0], files[1], files[2])
	s := files[0] + " " + files[1] + " " + files[2]
	fmt.Println(s)

	if err != nil {
		log.Printf("Erro parting template files: %v \n", err)
	} else {
		log.Printf("Parsed templates %v\n", t)
	}

	allButtons := []Button{
		Button{"Item1", 500, 300},
		Button{"Item2", 500, 300},
	}

	ButtonGroup1 := ButtonGroup{
		GridButtons: allButtons,
	}

	htmlSections := []string{
		"header",
		"body",
		"altfooter",
		"footer.html",
		"dilly.html",
	}

	for i := range htmlSections {
		t.ExecuteTemplate(os.Stdout, htmlSections[i], ButtonGroup1)
	}

}

// Button is a data type that describes all aspects of a POS button
type Button struct {
	Name   string
	Width  int
	Height int
}

// ButtonGroup is a colletion of buttons
type ButtonGroup struct {
	GridButtons []Button
}
