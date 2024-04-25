package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func multiplicationTable(lim int) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight)
	defer writer.Flush()

	for str := 0; str <= lim; str++ {
		for col := 0; col <= lim; col++ {
			if str == 0 && col == 0 {
				fmt.Fprint(writer, "  \t")
				continue
			}
			if str == 0 {
				fmt.Fprint(writer, col, "  \t")
				continue
			}
			if col == 0 {
				fmt.Fprint(writer, str, "  \t")
				continue
			}
			fmt.Fprint(writer, str*col, "  \t")
		}
		fmt.Fprint(writer, "\n")
	}

}
