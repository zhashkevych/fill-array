package main

import "testing"

func TestIntArray_Fill(t *testing.T) {
	testTable := []struct {
		name         string
		sizeX, sizeY int
		values       []int
		wantInitErr  bool
		wantFillErr  bool
	}{
		{
			name:  "Ok",
			sizeX: 2,
			sizeY: 2,
			values: []int{
				1, 2, 3, 4,
			},
		},
		{
			name:        "SizeX is 0",
			sizeX:       0,
			wantInitErr: true,
		},
		{
			name:        "SizeX is Negative",
			sizeX:       -1,
			wantInitErr: true,
		},
		{
			name:        "SizeY is 0",
			sizeX:       2,
			sizeY:       0,
			wantInitErr: true,
		},
		{
			name:        "SizeY is Negative",
			sizeX:       2,
			sizeY:       -2,
			wantInitErr: true,
		},
		{
			name:        "Values Array Has Less Elements",
			sizeX:       2,
			sizeY:       2,
			values:      []int{1, 2, 3},
			wantFillErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			intArray, err := NewIntArray(testCase.sizeX, testCase.sizeY)
			if err != nil && !testCase.wantInitErr {
				t.Fatal("Unexpected Initialization Error")
			}

			if err != nil && testCase.wantInitErr {
				t.Skip("Ok")
			}

			if err == nil && testCase.wantInitErr {
				t.Fatal("Initialization Error was expected, but got nil")
			}

			err = intArray.Fill(testCase.values)
			if err != nil && !testCase.wantFillErr {
				t.Fatal("Unexpected Fill Error")
			}

			if err != nil && testCase.wantFillErr {
				t.Skip("Ok")
			}

			if err == nil && testCase.wantFillErr {
				t.Fatal("Initialization Error was expected, but got nil")
			}
		})
	}
}