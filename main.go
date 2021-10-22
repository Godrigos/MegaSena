package main

import (
	"fmt"
)

func main() {
	resultado := getContent()
	for i := range resultado {
		fmt.Printf("%02d - %03d - %.2f%%\n", resultado[i].Ten,
			resultado[i].Freq, resultado[i].Perc)
	}
}
