package main

import (
	"reflect"
	"testing"
)

//Test functions here

func TestFormatGames(t *testing.T) {
	type test struct {
		input []GameInfo
		output string
	}
	testGames := make([]GameInfo, 1)
	for i := range testGames {
		testGames[i].Name = "FunGame"
		testGames[i].Price = "$40.00"
		testGames[i].ReleaseDate = "Jan 01, 2022"
	}
	//remove any character in the following string to see the test fail
	expectedOutput := "Title: *FunGame* Price: *$40.00* Release: *Jan 01, 2022*\n"
	tests := []test{
		{input: testGames, output: expectedOutput},
	}
	for _, tc := range tests {
		got := formatGames(tc.input)
		if !reflect.DeepEqual(tc.output, got) {
			t.Fatalf("expected %v, got: %v", tc.output, got)
		}
	}
}

func TestCreateJson(t *testing.T) {
	type test struct {
		input []GameInfo
		fileName string
		output error
	}
	testGames := make([]GameInfo, 1)
	for i := range testGames {
		testGames[i].Name = "FunGame"
		testGames[i].Price = "$40.00"
		testGames[i].ReleaseDate = "Jan 01, 2022"
	}
	tests := []test{
		{input: testGames, fileName: "output.json", output: nil},
	}
	for _, tc := range tests {
		got := createJson(tc.input, tc.fileName)
		if got != nil {
			t.Fatalf("Error encountered: %v", got)
		}
	}
}

func BenchmarkFormatGames(b *testing.B) {
	gamesForBenchmarking := make([]GameInfo, 10)
	for i := range gamesForBenchmarking {
		gamesForBenchmarking[i].Name = "gamename"
		gamesForBenchmarking[i].Price = "gameprice"
		gamesForBenchmarking[i].ReleaseDate = "gamerelease"
}
	for i := 0; i < b.N; i++ {
		formatGames(gamesForBenchmarking)
	}
}