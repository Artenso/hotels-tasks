package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func multiplicationTable(lim int) {
	// create writer with columns formatting
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight)
	// print resolt in th end of working
	defer writer.Flush()
	// filling table string by string
	for str := 0; str <= lim; str++ {
		for col := 0; col <= lim; col++ {
			// in upper right corner we need empty cell
			if str == 0 && col == 0 {
				fmt.Fprint(writer, "  \t")
				continue
			}
			// first string must be from one to lim
			if str == 0 {
				fmt.Fprint(writer, col, "  \t")
				continue
			}
			// first collumn must be from one to lim
			if col == 0 {
				fmt.Fprint(writer, str, "  \t")
				continue
			}
			// fill table cells with values
			fmt.Fprint(writer, str*col, "  \t")
		}
		// jump to new string
		fmt.Fprint(writer, "\n")
	}

}
