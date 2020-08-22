package transcript

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ledongthuc/pdf"
)

// Term stores information about a single term
type Term struct {
	Name   string
	Grades []Grade
}

// Grade stores information about a single grade
type Grade struct {
	Subject    string
	CourseCode string
	Grade      int
}

// GradeToGPA converts a grade into its GPA equivalent
func GradeToGPA(grade int) float64 {
	if 90 <= grade && grade <= 100 {
		return 4
	}
	if 85 <= grade && grade < 90 {
		return 3.9
	}
	if 80 <= grade && grade < 85 {
		return 3.7
	}
	if 77 <= grade && grade < 80 {
		return 3.3
	}
	if 73 <= grade && grade < 77 {
		return 3
	}
	if 70 <= grade && grade < 73 {
		return 2.7
	}
	if 67 <= grade && grade < 70 {
		return 2.3
	}
	if 63 <= grade && grade < 67 {
		return 2
	}
	if 60 <= grade && grade < 63 {
		return 1.7
	}
	if 57 <= grade && grade < 60 {
		return 1.3
	}
	if 53 <= grade && grade < 57 {
		return 1
	}
	if 50 <= grade && grade < 53 {
		return 0.7
	}
	if 0 <= grade && grade < 50 {
		return 0
	}
	// TODO: or have error
	return 0
}

// ExtractPdf takes the path of a pdf, and extracts the text into slices of slices of strings
func ExtractPdf(path string) ([][]string, error) {
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return make([][]string, 0), err
	}

	pdfOutput := make([][]string, 1)
	totalPage := r.NumPage()
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			rowOutput := make([]string, 1)
			for _, word := range row.Content {
				rowOutput = append(rowOutput, word.S)
			}
			pdfOutput = append(pdfOutput, rowOutput)

		}
	}
	return pdfOutput, nil
}

// ParsePdf takes the data from the pdf and returns grades by term
func ParsePdf(pdf [][]string) []Term {
	var terms = make([]Term, 0)
	var termRegexp = regexp.MustCompile("(Fall|Winter|Spring) ([0-9][0-9][0-9][0-9])")
	var curTerm Term = Term{}

	for _, row := range pdf {
		if len(row) == 0 {
			continue
		}
		// we remove the first element since its always ""
		row = row[1:]
		if len(row) == 1 && len(strings.Fields(row[0])) == 0 {
			continue
		}

		// this is trimming the text
		for i, col := range row {
			col = strings.Join(strings.Fields(col), " ")
			row[i] = col
		}

		// detects if a new term started
		if len(row) == 1 && termRegexp.MatchString(row[0]) {
			if len(curTerm.Grades) > 0 {
				terms = append(terms, curTerm)
			}
			curTerm = Term{Name: row[0]}
		}

		if len(row) == 6 {
			matched, _ := regexp.MatchString(`^[A-Z]+$`, row[0])
			if !matched {
				continue
			}
			// we ignore these
			if row[0] == "COOP" || row[0] == "PD" || row[0] == "WKRPT" {
				continue
			}
			// these don't count (not sure about WF)
			if row[5] == "WD" || row[5] == "WF" {
				continue
			}
			grade, err := strconv.Atoi(row[5])
			if err != nil {
				fmt.Println(err)
				continue
			}
			// fmt.Printf("%+q\n", row)
			curTerm.Grades = append(curTerm.Grades, Grade{row[0], row[1], grade})
		}
	}
	return terms
}
