package main

import (
	"reflect"
	"testing"
)

func TestIsCardinalIncrementalOrder(t *testing.T) {
	type args struct {
		counts  []int
		lastIdx int
	}
	tests := []struct {
		name  string
		args  args
		wantB bool
	}{
		{
			name: "1, 2, 3, 4 should return true",
			args: args{
				counts:  []int{1, 2, 3, 4},
				lastIdx: 3,
			},
			wantB: true,
		},
		{
			name: "1, 1, 2, 3, 4 should return false",
			args: args{
				counts:  []int{1, 1, 2, 3, 4},
				lastIdx: 4,
			},
			wantB: false,
		},
		{
			name: "2, 3, 4, 5 should return false",
			args: args{
				counts:  []int{2, 3, 4, 5},
				lastIdx: 3,
			},
			wantB: false,
		},
		{
			name: "0, 1, 2, 3, 4 should return false",
			args: args{
				counts:  []int{0, 1, 2, 3, 4},
				lastIdx: 4,
			},
			wantB: false,
		},
		{
			name: "1, 2, 3, 3 should return false",
			args: args{
				counts:  []int{1, 2, 3, 3},
				lastIdx: 3,
			},
			wantB: false,
		},
		{
			name: "1 should return true",
			args: args{
				counts:  []int{1},
				lastIdx: 0,
			},
			wantB: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := IsCardinalIncrementalOrder(tt.args.counts, tt.args.lastIdx); gotB != tt.wantB {
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
		// TODO: Add test cases.
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
		name  string
		args  args
		want  *Ledger
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LedgerFromWord(tt.args.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LedgerFromWord() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LedgerFromWord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSortLedgerCounts(t *testing.T) {
	type args struct {
		ledger    *Ledger
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
			if got := SortLedgerCounts(tt.args.ledger, tt.args.lenLedger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortLedgerCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
