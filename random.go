package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type GetUniqueRandomIntegersInput struct {
	SizeX       int
	SizeY       int
	RandomLimit int
}

func (i GetUniqueRandomIntegersInput) Validate() error {
	if i.SizeX <= 0 || i.SizeY <= 0 {
		return errors.New("array length can't be negative or 0")
	}

	if i.RandomLimit <= 0 {
		return errors.New("random limit can't be 0")
	}

	if i.RandomLimit < i.SizeX*i.SizeY {
		return errors.New(
			fmt.Sprintf("can't generate unique numbers with limit %d for %dx%d sized array",
				i.RandomLimit, i.SizeX, i.SizeY))
	}

	return nil
}

type Randomizer struct {
	uniqueIntegers map[int]bool
}

func NewRandomizer() *Randomizer {
	rand.Seed(time.Now().Unix())

	return &Randomizer{
		uniqueIntegers: make(map[int]bool),
	}
}

func (r *Randomizer) GetUniqueRandomIntegers(input GetUniqueRandomIntegersInput) ([]int, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	r.reset()

	numbersCount := input.SizeX * input.SizeY
	result := make([]int, 0)

	for len(r.uniqueIntegers) < numbersCount {
		r.uniqueIntegers[r.getRandomInt(input.RandomLimit)] = true
	}

	for num := range r.uniqueIntegers {
		result = append(result, num)
	}

	return result, nil
}

func (r *Randomizer) reset() {
	r.uniqueIntegers = make(map[int]bool)
}

func (r *Randomizer) getRandomInt(limit int) int {
	return rand.Intn(limit)
}
