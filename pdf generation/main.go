package main

import (
	"fmt"
	"log"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	//wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
		return
	}
	htmlStr := `<html><body><h1 style="color:red;">This is an html
from pdf to test color<h1><img src="http://api.qrserver.com/v1/create-qr-
code/?data=HelloWorld" alt="img" height="42" width="42"></img></body></html>`

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile("./pdf/test.pdf")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done")
}
