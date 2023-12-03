package getdataex

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

func Getppr(week int) (string, string) {
	var (
		sheet   string
		month   = int(time.Now().Month())
		result1 string
		result2 string
	)
	sheet = "ОЗХ ЗГПН"
	answer := []string{}

	f, err := excelize.OpenFile("ППР.xlsx")
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(answer)

	}

	x := cols[0]
	for i, res := range cols[(month*4)-week] {
		if res != "" && i < 88 {
			stell, _ := excelize.CoordinatesToCellName((month*4)-week+1, i)
			st, _ := f.GetCellStyle(sheet, stell)
			fmt.Println(st, "xто тут", stell)

			result1 += "\n <b>" + x[i] + ": </b>" + " " + res + " " + fmt.Sprint(st)
		} else if res != "" && i >= 88 {
			result2 += "\n <b>" + x[i] + ": </b>" + " " + res
		}
	}
	return result1, result2
}
