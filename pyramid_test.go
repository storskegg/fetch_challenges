package main

import (
	"reflect"
	"testing"
)

func TestIsCardinalIncrementalOrder(t *testing.T) {
	type args struct {
		counts []int
	}
	tests := []struct {
		name  string
		args  args
		wantB bool
	}{
		{
			name: "1, 2, 3, 4 should return true",
			args: args{
				counts: []int{1, 2, 3, 4},
			},
			wantB: true,
		},
		{
			name: "1, 1, 2, 3, 4 should return false",
			args: args{
				counts: []int{1, 1, 2, 3, 4},
			},
			wantB: false,
		},
		{
			name: "2, 3, 4, 5 should return false",
			args: args{
				counts: []int{2, 3, 4, 5},
			},
			wantB: false,
		},
		{
			name: "0, 1, 2, 3, 4 should return false",
			args: args{
				counts: []int{0, 1, 2, 3, 4},
			},
			wantB: false,
		},
		{
			name: "1, 2, 3, 3 should return false",
			args: args{
				counts: []int{1, 2, 3, 3},
			},
			wantB: false,
		},
		{
			name: "1 should return true",
			args: args{
				counts: []int{1},
			},
			wantB: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := IsCardinalIncrementalOrder(tt.args.counts); gotB != tt.wantB {
				t.Errorf("IsCardinalIncrementalOrder() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestIsPyramidWord(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name  string
		args  args
		wantB bool
	}{
		{
			name: "banana should be true",
			args: args{
				word: "banana",
			},
			wantB: true,
		},
		{
			name: "banANa should be true",
			args: args{
				word: "banANa",
			},
			wantB: true,
		},
		{
			name: "bandana should be false",
			args: args{
				word: "bandana",
			},
			wantB: false,
		},
		{
			name: "add should be true",
			args: args{
				word: "add",
			},
			wantB: true,
		},
		{
			name: "oh should be false",
			args: args{
				word: "oh",
			},
			wantB: false,
		},
		{
			name: "a should be true",
			args: args{
				word: "a",
			},
			wantB: true,
		},
		{
			name: "add should be true",
			args: args{
				word: "add",
			},
			wantB: true,
		},
		{
			name: "atl should be false",
			args: args{
				word: "atl",
			},
			wantB: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := IsPyramidWord(tt.args.word); gotB != tt.wantB {
				t.Errorf("IsPyramidWord() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestLedgerFromWord(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want Ledger
	}{
		{
			name: "banana to ledger",
			args: args{
				word: "banana",
			},
			want: Ledger{
				'a': 3,
				'b': 1,
				'n': 2,
			},
		},
		{
			name: "banANa to ledger",
			args: args{
				word: "banANa",
			},
			want: Ledger{
				'a': 3,
				'b': 1,
				'n': 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LedgerFromWord(tt.args.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LedgerFromWord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortLedgerCounts(t *testing.T) {
	type args struct {
		ledger    Ledger
		lenLedger int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortLedgerCounts(tt.args.ledger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortLedgerCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
