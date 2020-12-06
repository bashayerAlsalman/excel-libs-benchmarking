package main

import (
	"fmt"

	"github.com/plandem/xlsx"
)

func main() {
	xl, err := xlsx.Open("../example_big.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	var sheet xlsx.Sheet
	// Iterate sheets via iterator
	for sheets := xl.Sheets(); sheets.HasNext(); {
		_, sheet = sheets.Next()

		// Iterate rows via iterator
		for rows := sheet.Rows(); rows.HasNext(); {
			index, row := rows.Next()
			// Iterate row's cells via iterator
			for cells := row.Cells(); cells.HasNext(); {
				_, _, cell := cells.Next()
				fmt.Printf(" row %+v cell %+v \n", index, cell.Value())
			}
		}
	}

}
