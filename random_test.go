package main

import (
	"testing"
)

func TestRandomizer_GetUniqueRandomIntegers(t *testing.T) {
	randomizer := NewRandomizer()

	testTable := []struct {
		name         string
		sizeX, sizeY int
		expectedLen  int
		randomLimit  int
		wantErr      bool
	}{
		{
			name: "Ok",
			sizeX:       5,
			sizeY:       5,
			expectedLen: 5 * 5,
			randomLimit: 100,
		},
		{
			name: "SizeX is 0",
			sizeX:   0,
			wantErr: true,
		},
		{
			name: "SizeX is Negative",
			sizeX:   -1,
			wantErr: true,
		},
		{
			name: "SizeY is 0",
			sizeX:   5,
			sizeY:   0,
			wantErr: true,
		},
		{
			name: "SizeY is Negative",
			sizeX:   5,
			sizeY:   -5,
			wantErr: true,
		},
		{
			name: "Random Limit is 0",
			sizeX:       5,
			sizeY:       5,
			randomLimit: 0,
			wantErr:     true,
		},
		{
			name: "Random Limit is Negative",
			sizeX:       5,
			sizeY:       5,
			randomLimit: -1,
			wantErr:     true,
		},
		{
			name: "Random Limit is lower than array size",
			sizeX:       5,
			sizeY:       5,
			randomLimit: 10,
			wantErr:     true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := randomizer.GetUniqueRandomIntegers(GetUniqueRandomIntegersInput{
				testCase.sizeX, testCase.sizeY, testCase.randomLimit})
			if err != nil && !testCase.wantErr {
				t.Fatal(err)
			}

			if err == nil && testCase.wantErr {
				t.Fatal("Should fail with error")
			}

			if len(result) != testCase.expectedLen {
				t.Fatalf("Incorrect result length. Want %d, got %d", testCase.expectedLen, len(result))
			}
		})
	}
}

func TestGetUniqueRandomIntegersInput_Validate(t *testing.T) {
	testTable := []struct{
		name string
		input GetUniqueRandomIntegersInput
		wantErr bool
	} {
		{
			name: "Ok",
			input: GetUniqueRandomIntegersInput{
				SizeX: 5,
				SizeY: 5,
				RandomLimit: 25,
			},
		},
		{
			name: "SizeX is 0",
			input: GetUniqueRandomIntegersInput{
				SizeX: 0,
				SizeY: 5,
				RandomLimit: 25,
			},
			wantErr: true,
		},
		{
			name: "SizeX is Negative",
			input: GetUniqueRandomIntegersInput{
				SizeX: -1,
				SizeY: 5,
				RandomLimit: 25,
			},
			wantErr: true,
		},
		{
			name: "SizeY is 0",
			input: GetUniqueRandomIntegersInput{
				SizeX: 5,
				SizeY: 0,
				RandomLimit: 25,
			},
			wantErr: true,
		},
		{
			name: "SizeY is Negative",
			input: GetUniqueRandomIntegersInput{
				SizeX: 5,
				SizeY: -1,
				RandomLimit: 25,
			},
			wantErr: true,
		},
		{
			name: "Random Limit is 0",
			input: GetUniqueRandomIntegersInput{
				SizeX: 5,
				SizeY: 5,
				RandomLimit: 0,
			},
			wantErr: true,
		},
		{
			name: "Random Limit is Negative",
			input: GetUniqueRandomIntegersInput{
				SizeX: 5,
				SizeY: 5,
				RandomLimit: -1,
			},
			wantErr: true,
		},
		{
			name: "Random Limit is lower than array size",
			input: GetUniqueRandomIntegersInput{
				SizeX: 5,
				SizeY: 5,
				RandomLimit: 24,
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.input.Validate()
			if err != nil && !testCase.wantErr {
				t.Fatal("Got unexpected error")
			}
		})
	}
}