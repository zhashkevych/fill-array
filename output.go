package main

import "fmt"

func Print(intArray *IntArray) error {
	arr, err := intArray.Get()
	if err != nil {
		return err
	}

	for rowIndex, rows := range arr {
		fmt.Printf("%d. %v\n", rowIndex+1, rows)
	}

	return nil
}
