package main

import (
	"fmt"
	"math"
	"os"

	"github.com/stevenw47/uwTranscript/transcript"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide the path to the transcript as an argument.")
		return
	}
	pdf, err := transcript.ExtractPdf(os.Args[1])
	if err != nil {
		panic(err)
	}

	terms := transcript.ParsePdf(pdf)
	var totalGrade float64
	var totalGPA float64
	var totalCount int
	for _, term := range terms {
		var termGrade float64
		var termGPA float64
		var termCount int

		for _, grade := range term.Grades {
			termGrade += float64(grade.Grade)
			termGPA += transcript.GradeToGPA(grade.Grade)
			termCount++
		}
		totalGrade += termGrade
		totalGPA += termGPA
		totalCount += termCount
		fmt.Printf("%v (%v)\n", term.Name, termCount)
		fmt.Printf("Term Average: %v\n", termGrade/float64(termCount))
		fmt.Printf("Term GPA: %v\n", math.Round(100*termGPA/float64(termCount))/100)
		fmt.Println("")
	}
	fmt.Printf("Overall (%v)\n", totalCount)
	fmt.Printf("Average: %v\n", totalGrade/float64(totalCount))
	fmt.Printf("GPA: %v\n", math.Round(100*totalGPA/float64(totalCount))/100)
}
