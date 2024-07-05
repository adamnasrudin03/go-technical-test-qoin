package main

import (
	"flag"

	"github.com/adamnasrudin03/go-technical-test-qoin/game"
)

// go run main.go -N=3 -M=4
func main() {
	var (
		N = flag.Int("N", 3, " Jumlah pemain")
		M = flag.Int("M", 4, "Jumlah dadu")
	)

	flag.Parse()
	game.Dice(*N, *M)
}
