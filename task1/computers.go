package main

import (
	"fmt"
)

func computers(count int) string {
	if count%100 >= 11 && count%100 <= 19 {
		return fmt.Sprintf("%d компьютеров", count)
	}
	if count%10 >= 2 && count%100 <= 4 {
		return fmt.Sprintf("%d компьютера", count)
	}
	if count%10 == 1 {
		return fmt.Sprintf("%d компьютер", count)
	} else {
		return fmt.Sprintf("%d компьютеров", count)
	}
}
