package game

import (
	"fmt"
	"strconv"
	"strings"
)

func printPlayer(playerPoint map[int]int, playerPicked map[int][]int) {
	for i := 0; i < len(playerPoint); i++ {
		// Print and Format value picked array
		valuesText := []string{}
		for key := range playerPicked[i] {
			text := strconv.Itoa(playerPicked[i][key])
			valuesText = append(valuesText, text)
		}
		textPicked := strings.Join(valuesText, ",")
		fmt.Printf("\t Pemain #%v (%v): %v \n", i+1, playerPoint[i], textPicked)
	}
}

func initPlayerPoint(players int) map[int]int {
	var playerPoint = map[int]int{}
	for i := 0; i < players; i++ {
		playerPoint[i] = 0
	}
	return playerPoint
}

func deleteInSlice(param int, arr []int) (res []int) {
	for _, value := range arr {
		if value != param {
			res = append(res, value)
		}
	}
	return res
}

func cekValueInSlice(param int, arr []int) bool {
	for _, value := range arr {
		if value == param {
			return true
		}
	}
	return false
}
