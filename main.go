package main

import (
	"fmt"

	"github.com/stevenw47/uwTranscript/transcript"
)

func main() {

	pdf, err := transcript.ExtractPdf("transcript.pdf")
	if err != nil {
		panic(err)
	}

	grade, gpa := transcript.ParsePdf(pdf)
	fmt.Printf("Grade Average: %v\n", grade)
	fmt.Printf("GPA: %v\n", gpa)
}
