package main

import (
	"log"
)

const (
	rowsCount    = 5
	columnsCount = 5
	randomLimit  = 100
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

	if err := array.Fill(randomIntegers); err != nil {
		log.Fatal(err)
	}

	if err := array.Print(); err != nil {
		log.Fatal(err)
	}
}
