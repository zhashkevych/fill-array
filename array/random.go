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
	if i.SizeX <= 0  || i.SizeY <= 0 {
		return errors.New("array length can't be negative or 0")
	}

	if i.RandomLimit <= 0 {
		return errors.New("random limit can't be 0")
	}

	if i.RandomLimit < i.SizeX * i.SizeY {
		return errors.New(
			fmt.Sprintf("can't generate unique numbers with limit %d for %dx%d sized array",
				i.RandomLimit, i.SizeX, i.SizeY))
	}

	return nil
}

type Randomizer struct {
	usedValues map[interface{}]bool
}

func NewRandomizer() *Randomizer {
	rand.Seed(time.Now().Unix())

	return &Randomizer{
		usedValues: make(map[interface{}]bool),
	}
}

func (r *Randomizer) GetUniqueRandomIntegers(input GetUniqueRandomIntegersInput) ([]int, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	r.reset()

	numbersCount := input.SizeX * input.SizeY
	result := make([]int, 0)

	for i := 0; len(r.usedValues) < numbersCount; i++ {
		r.usedValues[r.getRandomInt(input.RandomLimit)] = true
	}

	for num := range r.usedValues {
		result = append(result, num.(int))
	}

	return result, nil
}

func (r *Randomizer) reset() {
	r.usedValues = make(map[interface{}]bool)
}

func (r *Randomizer) getRandomInt(limit int) int {
	return rand.Intn(limit)
}
