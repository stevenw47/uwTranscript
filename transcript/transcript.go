package transcript

import (
	// "fmt"
	"strings"
	"strconv"
	"github.com/ledongthuc/pdf"
)

func gradeToGPA(grade int) float64 {
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

// ParsePdf takes the data from the pdf and returns the average grade and GPA
func ParsePdf(pdf [][]string) (float64, float64) {
	var totalGrade float64
	var totalGPA float64
	var count int
	for _, row := range pdf {
		// TODO: also if its an array containing only whitespace strings
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
			col = strings.Join(strings.Fields(col)," ")
			row[i] = col
		}
		// fmt.Printf("%+q\n", row)

		// TODO: temporary hacky way to get grade
		if len(row) == 6 {
			if row[0] == "Program:" {
				continue
			}
			// we ignore these
			if row[0] == "COOP" || row[0] == "PD" || row[0] == "WKRPT" {
				continue
			}
			// these don't count
			if row[5] == "WD" || row[5] == "WF" {
				continue
			}
			// fmt.Printf("%+q\n", row)
			grade, err := strconv.Atoi(row[5])
			if err != nil {
				// fmt.Println(err)
				continue
			}
			totalGrade += float64(grade)
			totalGPA += gradeToGPA(grade)
			count++
		}
	}
	return totalGrade / float64(count), totalGPA / float64(count)
}
