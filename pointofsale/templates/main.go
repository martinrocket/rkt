package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
)

const tpl string = `
<!DOCTYPE html>
<html>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
            <script type="text/javascript">
                $(document).ready(function() {
                    $("#callGo1").on('click', function() {
                        $.ajax({
                            url: "http://localhost:8080/API/1?id=10",
                            method: "GET",
                            success: function(data) {
                                $("#response").html(data);
                            },
                        });
					});
				});
			</script>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Buttons}}<div><button>{{ . }}</button></div>{{else}}<div><strong>no rows</strong></div>{{end}}
		<div><button>{{(index .Buttons 1)}}</button></div>
		<button>{{(index .Buttons 0)}}</button>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

func main() {
	fmt.Println("hello")
	fmt.Printf(tpl)
	createTmpl()
	t, err := template.New("index").Parse(tpl)
	if err != nil {
		log.Printf("Could not parse template, %v with error %v", t, err)
	}
	data := struct {
		Title   string
		Items   []string
		Buttons []string
	}{
		Title: "Rkt POS",
		Items: []string{
			"Start Order",
			"Cash Out",
		},
		Buttons: []string{
			"Button1",
			"button2",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Printf("Could not Execute template index.html, %v with error %v\n", t, err)
	} else {
		log.Printf("Executed template %v\n", t)
		f, _ := os.Create("edit.html")
		err = t.Execute(f, data)

	}

}

func createTmpl() {
	f, err := os.Create("index.html")
	defer f.Close()
	if err != nil {
		log.Printf("Could not build index.html: %v\n", err)
	} else {
		log.Printf("Created index.html err = %v\n", err)
	}

	io.WriteString(f, tpl)

}
