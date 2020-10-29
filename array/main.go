package main

import (
	"fmt"
	"log"
)



const (
	rowsCount    = 5
	columnsCount = 5

	randomLimit = 100
)

func main() {
	randomizer := NewRandomizer()
	randomIntegers, err := randomizer.GetUniqueRandomIntegers(GetUniqueRandomIntegersInput{
		SizeX:       rowsCount,
		SizeY:       columnsCount,
		RandomLimit: randomLimit,
	})
	if err != nil {
		log.Fatal(err)
	}

	array, err := NewIntArray(rowsCount, columnsCount)
	if err != nil {
		log.Fatal(err)
	}

	array.Fill(randomIntegers)

	for rowIndex, rows := range array.Get() {
		fmt.Printf("%d. %v\n", rowIndex+1, rows)
	}
}