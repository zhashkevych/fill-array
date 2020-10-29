package main

import (
	"errors"
	"fmt"
)

type IntArray struct {
	sizeX int
	sizeY int
	arr   [][]int
}

func NewIntArray(sizeX int, sizeY int) (*IntArray, error) {
	if sizeX <= 0 || sizeY <= 0 {
		return nil, errors.New("array size can't be negative or 0")
	}

	return &IntArray{sizeX: sizeX, sizeY: sizeY}, nil
}

func (a *IntArray) Fill(values []int) error {
	if len(values) != a.sizeX*a.sizeY {
		return errors.New("values array has less elements than target array")
	}

	rows := make([][]int, a.sizeX)
	counter := 0

	for row := 0; row < a.sizeX; row++ {
		columns := make([]int, a.sizeY)
		rows[row] = columns

		for col := 0; col < a.sizeY; col++ {
			rows[row][col] = values[counter]
			counter++
		}
	}

	a.arr = rows
	return nil
}

func (a *IntArray) Get() ([][]int, error) {
	if a.arr == nil {
		return nil, errors.New("array is not filled with values")
	}

	return a.arr, nil
}

func (a *IntArray) Print() error {
	if a.arr == nil {
		return errors.New("array is not filled with values")
	}

	for rowIndex, rows := range a.arr {
		fmt.Printf("%d. %v\n", rowIndex+1, rows)
	}

	return nil
}