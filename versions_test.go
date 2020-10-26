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
		{
			name: "Given 1.0.0 and 1.0.1, should get 1.0.0 is before 1.0.1",
			args: args{
				versionA: "1.0.0",
				versionB: "1.0.1",
			},
			want:    "1.0.0 is before 1.0.1",
			wantErr: false,
		},
		{
			name: "Given 2.0 and 1.0.17, should get 2.0 is after 1.0.17",
			args: args{
				versionA: "2.0",
				versionB: "1.0.17",
			},
			want:    "2.0 is after 1.0.17",
			wantErr: false,
		},
		{
			name: "Given 1.18.0 and 1.18.0, should get 1.18.0 is equal to 1.18.0",
			args: args{
				versionA: "1.18.0",
				versionB: "1.18.0",
			},
			want:    "1.18.0 is equal to 1.18.0",
			wantErr: false,
		},
		{
			name: "Given 1.18.0 and 1.18, should get 1.18.0 is equal to 1.18",
			args: args{
				versionA: "1.18.0",
				versionB: "1.18",
			},
			want:    "1.18.0 is equal to 1.18",
			wantErr: false,
		},
		{
			name: "Given 2.1 and 2.1.0, should get 2.1 is equal to 2.1.0",
			args: args{
				versionA: "2.1",
				versionB: "2.1.0",
			},
			want:    "2.1 is equal to 2.1.0",
			wantErr: false,
		},
		{
			name: "Given 1.18.1 and 1.18, should get 1.18.1 is after 1.18",
			args: args{
				versionA: "1.18.1",
				versionB: "1.18",
			},
			want:    "1.18.1 is after 1.18",
			wantErr: false,
		},
		{
			name: "Given 2.1 and 2.1.1, should get 2.1 is before 2.1.1",
			args: args{
				versionA: "2.1",
				versionB: "2.1.1",
			},
			want:    "2.1 is before 2.1.1",
			wantErr: false,
		},
		{
			name: "Given 1.99.99 and 2.00, should get 1.99.99 is before 2.00",
			args: args{
				versionA: "1.99.99",
				versionB: "2.00",
			},
			want:    "1.99.99 is before 2.00",
			wantErr: false,
		},
		{
			name: "Given pearljam and ween, should error",
			args: args{
				versionA: "pearljam",
				versionB: "ween",
			},
			want:    "",
			wantErr: true,
		},
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
