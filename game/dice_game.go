package game

import (
	"fmt"
	"math/rand"
	"sort"
)

func Dice(players, totalDice int) {
	var (
		playerPoint  = initPlayerPoint(players)
		playerPicked = map[int][]int{}
		playerStop   = map[int]bool{}
		dice         = []int{1, 2, 3, 4, 5, 6}
		throw        = 1
	)
	fmt.Println("Pemain = ", players, ", Dadu = ", totalDice)
	fmt.Println("======================")
	fmt.Printf("Giliran %v lempar dadu:	\n", throw)

	// first throw
	for i := 0; i < len(playerPoint); i++ {
		// Pick dice for each player, and increment player point if 6 is picked
		var picked = []int{}
		for j := 0; j < totalDice; j++ {
			// Pick a random dice
			randomIndex := rand.Intn(len(dice))
			pick := dice[randomIndex]
			picked = append(picked, pick)

			// If 6 is picked, increment player point
			if pick == 6 {
				playerPoint[i] = playerPoint[i] + 1
			}
		}

		playerPicked[i] = picked
	}

	printPlayer(playerPoint, playerPicked)

	// Evaluate
	fmt.Println("Setelah evaluasi:")
	for i := 0; i < len(playerPoint); i++ {
		// Check if player stop
		if playerStop[i] {
			continue
		}

		nextIndex := i
		if nextIndex+1 != len(playerPicked) {
			nextIndex = nextIndex + 1
		}

		for playerStop[nextIndex] {
			nextIndex++
			if nextIndex == len(playerPicked) {
				nextIndex = 0
			}
		}

		// Check if the current player has a 1 and the next player doesn't
		// If true, move the 1 to the next player and remove it from the current player
		if cekValueInSlice(1, playerPicked[i]) && (!cekValueInSlice(1, playerPicked[nextIndex])) {
			// Move the 1 to the next player
			playerPicked[nextIndex] = append(playerPicked[nextIndex], 1)

			// Remove the 1 from the current player
			playerPicked[i] = deleteInSlice(1, playerPicked[i])
		}

		// delete value 6
		playerPicked[i] = deleteInSlice(6, playerPicked[i])
		// sort descending
		sort.Slice(playerPicked[i], func(x, y int) bool {
			return playerPicked[i][x] > playerPicked[i][y]
		})

		// check if player stop
		if len(playerPicked[i]) == 0 {
			playerStop[i] = true
		}
	}

	printPlayer(playerPoint, playerPicked)

	fmt.Println("======================")

	// 2nd throw and so on
	for {
		throw++
		fmt.Printf("Giliran %v lempar dadu:	\n", throw)

		// Picked
		for i := 0; i < len(playerPoint); i++ {
			// Check if player stop
			if playerStop[i] {
				continue
			}

			// Pick dice for each player
			var picked = []int{}
			for j := 0; j < len(playerPicked[i]); j++ {
				randomIndex := rand.Intn(len(dice))
				pick := dice[randomIndex]
				picked = append(picked, pick)

				if pick == 6 {
					playerPoint[i] = playerPoint[i] + 1
				}
			}

			playerPicked[i] = picked
		}

		printPlayer(playerPoint, playerPicked)

		// Evaluate
		fmt.Println("Setelah evaluasi:")
		for i := 0; i < len(playerPoint); i++ {
			// Check if player stop
			if playerStop[i] {
				continue
			}

			nextIndex := i
			if nextIndex+1 != len(playerPicked) {
				nextIndex = nextIndex + 1
			}

			for playerStop[nextIndex] {
				nextIndex++
				if nextIndex == len(playerPicked) {
					nextIndex = 0
				}
			}

			// Check if the current player has a 1 and the next player doesn't
			// If true, move the 1 to the next player and remove it from the current player
			if cekValueInSlice(1, playerPicked[i]) && (!cekValueInSlice(1, playerPicked[nextIndex])) {
				// Move the 1 to the next player
				playerPicked[nextIndex] = append(playerPicked[nextIndex], 1)

				// Remove the 1 from the current player
				playerPicked[i] = deleteInSlice(1, playerPicked[i])
			}

			// delete value 6
			playerPicked[i] = deleteInSlice(6, playerPicked[i])

			// sort descending
			sort.Slice(playerPicked[i], func(x, y int) bool {
				return playerPicked[i][x] > playerPicked[i][y]
			})

			// check if player stop
			if len(playerPicked[i]) == 0 {
				playerStop[i] = true
			}
		}

		printPlayer(playerPoint, playerPicked)

		fmt.Println("======================")
		var players = []int{}
		for i, v := range playerPicked {
			if len(v) > 0 {
				players = append(players, i)
			}
		}

		if len(players) < 2 {
			fmt.Printf("Game berakhir karena hanya pemain #%v yang memiliki dadu. \n", players[0]+1)
			break
		}
	}

	// get player win
	max, playerWin := 0, 0
	for index, value := range playerPoint {
		if value > max {
			max = value
			playerWin = index + 1
		}
	}

	fmt.Printf("Game dimenangkan oleh pemain #%v karena memiliki poin lebih banyak dari pemain lainnya. \n", playerWin)

}
