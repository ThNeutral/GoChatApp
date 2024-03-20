package main

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

func serveHTML(w io.Writer, data interface{}, path string, filenames ...string) {
	log.Printf("Connection registered to %s\n", path)

	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		fmt.Fprintf(w, "Internal error")
		log.Printf("Failed to parse files in %s handler. Error: %s\n", path, err.Error())
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Registered error while handling %s route. Error: %s\n", path, err.Error())
	}
}
