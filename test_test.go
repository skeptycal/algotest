package algotest

import (
	"testing"
)

func TestQuestionsMarks(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"acc?7??sss?3rr1??????5"}, "true"},
		{"2", args{"aa6?9"}, "false"},
		{"2", args{"5??aaaaaaaaaaaaaaaaaaa?5?5"}, "false"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuestionsMarks(tt.args.str); got != tt.want {
				t.Errorf("QuestionsMarks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstReverse(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		// TODO: Add test cases.
		{"1", "12345", "54321"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstReverse(tt.str); got != tt.want {
				t.Errorf("FirstReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBracketCombinations(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		// TODO: Add test cases.
		{"1", 2, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BracketCombinations(tt.num); got != tt.want {
				t.Errorf("BracketCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
