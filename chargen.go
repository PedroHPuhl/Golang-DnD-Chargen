package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func statdice() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	var dicea = (rand.Intn(max-min+1) + min)
	var diceb = (rand.Intn(max-min+1) + min)
	var dicec = (rand.Intn(max-min+1) + min)
	var diced = (rand.Intn(max-min+1) + min)

	var diceArray = [4]int{dicea, diceb, dicec, diced}

	sort.Slice(diceArray[:], func(i, j int) bool {
		return diceArray[:][i] > diceArray[:][j]
	})

	var newDiceArray = diceArray[:len(diceArray)-1]

	var diceSum int

	for i := 0; i < len(newDiceArray); i++ {
		diceSum = diceSum + newDiceArray[i]
	}

	return diceSum
}

func main() {

	minSleep := 1
	maxSleep := 5
	randomSleep := rand.Intn(maxSleep-minSleep+1) + minSleep

	sleepDuration := time.Duration(randomSleep)

	str := statdice()
	time.Sleep(sleepDuration)
	dex := statdice()
	time.Sleep(sleepDuration)
	con := statdice()
	time.Sleep(sleepDuration)
	wis := statdice()
	time.Sleep(sleepDuration)
	inl := statdice()
	time.Sleep(sleepDuration)
	cha := statdice()

	if str+dex+con+wis+inl+cha >= 80 {
		fmt.Print("Personagem muito OP, role novamente.")
	} else {
		fmt.Println("STR:", str)
		fmt.Println("DEX:", dex)
		fmt.Println("CON:", con)
		fmt.Println("WIS:", wis)
		fmt.Println("INT:", inl)
		fmt.Println("CHA:", cha)
	}
}
