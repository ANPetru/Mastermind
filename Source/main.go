package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randNumToCalc string
var randomNumber string
var guessNumber string

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber = getStringFromInt(rand.Intn(9999))
	guessNumber = ""
	fmt.Println("A 4 digit number has been randomly generated. Guess it!")
	for true {
		fmt.Println("Enter a 4 digits number.")
		fmt.Scanf("%s\n", &guessNumber)
		if len(guessNumber) != 4{
			fmt.Println("The number must have 4 digits.")
		} else {
			randNumToCalc = randomNumber
			samePos := getNumberInSamePosition()
			if samePos == 4 {
				fmt.Println("You won")
				break
			}
			diffPos := getNumeberInDifferentPosition()
			fmt.Printf("Numbers at same position: %d\n", samePos)
			fmt.Printf("Numbers at different position: %d\n", diffPos)
		}

	}
	fmt.Println("Press enter to exit-")
	fmt.Scanf("%s\n", &guessNumber)
}

func getNumeberInDifferentPosition() int{
	diffPos := 0
	for i := 0; i < 4; i++ {
		if guessNumber[i] != '-' {
			for j := 0; j < 4; j++ {
				if guessNumber[i] == randNumToCalc[j] {
					diffPos++
					guessNumber = getStringExceptPos(guessNumber, i)
					randNumToCalc = getStringExceptPos(randNumToCalc,i)

					break

				}
			}
		}
	}
	return diffPos
}

func getNumberInSamePosition() int{
	samePos := 0
	for i := 0; i < 4; i++ {
		if guessNumber[i] == randNumToCalc[i] {
			guessNumber = getStringExceptPos(guessNumber, i)
			randNumToCalc = getStringExceptPos(randNumToCalc,i)
			samePos++
		}
	}
	return samePos
}


func getStringExceptPos(s string, pos int) string {
	return s[:pos] + "-" + s[pos+1:]
}

func getStringFromInt(num int) string {
	return fmt.Sprintf("%04d", num)
}
