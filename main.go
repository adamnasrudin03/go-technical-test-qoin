package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		N = flag.Int("N", 3, " Jumlah pemain")
		M = flag.Int("M", 4, "Jumlah dadu")
	)

	flag.Parse()
	fmt.Println("N = ", *N, "M = ", *M)

}
