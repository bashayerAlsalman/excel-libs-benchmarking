package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//https://xuri.me/excelize/en/base/installation.html#chart
func main() {
	f, err := excelize.OpenFile("../example_big.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range f.GetSheetMap() {
		// Get all the rows in the Sheet1.
		rows, _ := f.GetRows(name)
		if err == nil {
			for rowID, row := range rows {
				for _, colCell := range row {
					fmt.Printf("row %+v cell %+v \n", rowID, colCell)
				}
			}
		}
	}

}
