package getdataex

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func Getdataex(queston string, sheet string) []string {
	answer := []string{}
	f, err := excelize.OpenFile("930-0910.xlsx")
	if err != nil {
		fmt.Println(err)
		return answer
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	cols, err := f.GetCols(sheet)
	if err != nil {
		fmt.Println(err)
		return answer
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return answer
	}

	for x, col := range cols[2] {
		if strings.Contains(col, queston) {
			var resalt string
			for i, row := range rows[x] {
				resalt += "<b>" + rows[11][i] + " - </b>" + row + "\n"
				//fmt.Println(rows[11][i], row)
			}
			answer = append(answer, resalt)
			resalt = ""
		}
	}

	return answer
}
