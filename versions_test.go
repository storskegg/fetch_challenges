package main

import "testing"

func TestNormalizeVersionStrings(t *testing.T) {
	type args struct {
		versionA string
		versionB string
	}
	tests := []struct {
		name            string
		args            args
		wantNormalizedA string
		wantNormalizedB string
	}{
		{
			name: "1.0.1 and 2.0.8",
			args: args{
				versionA: "1.0.1",
				versionB: "2.0.8",
			},
			wantNormalizedA: "1.0.1",
			wantNormalizedB: "2.0.8",
		},
		{
			name: "1.0.1 and 1.0",
			args: args{
				versionA: "1.0.1",
				versionB: "1.0",
			},
			wantNormalizedA: "1.0.1",
			wantNormalizedB: "1.0.0",
		},
		{
			name: "1.0 and 1.0.1",
			args: args{
				versionA: "1.0",
				versionB: "1.0.1",
			},
			wantNormalizedA: "1.0.0",
			wantNormalizedB: "1.0.1",
		},
		{
			name: "1 and 3.1.14.159.2.65.3.59",
			args: args{
				versionA: "1",
				versionB: "3.1.14.159.2.65.3.59",
			},
			wantNormalizedA: "1.0.0.0.0.0.0.0",
			wantNormalizedB: "3.1.14.159.2.65.3.59",
		},
		{
			name: "1 and 2",
			args: args{
				versionA: "1",
				versionB: "2",
			},
			wantNormalizedA: "1",
			wantNormalizedB: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNormalizedA, gotNormalizedB := NormalizeVersionStrings(tt.args.versionA, tt.args.versionB)
			if gotNormalizedA != tt.wantNormalizedA {
				t.Errorf("NormalizeVersionStrings() gotNormalizedA = %v, want %v", gotNormalizedA, tt.wantNormalizedA)
			}
			if gotNormalizedB != tt.wantNormalizedB {
				t.Errorf("NormalizeVersionStrings() gotNormalizedB = %v, want %v", gotNormalizedB, tt.wantNormalizedB)
			}
		})
	}
}

func TestVersionStringCompare(t *testing.T) {
	type args struct {
		versionA string
		versionB string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VersionStringCompare(tt.args.versionA, tt.args.versionB)
			if (err != nil) != tt.wantErr {
				t.Errorf("VersionStringCompare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VersionStringCompare() got = %v, want %v", got, tt.want)
			}
		})
	}
}
