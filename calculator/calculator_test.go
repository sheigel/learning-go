package calculator

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1 and 2", args: args{a: 1, b: 2}, want: 3},
		{name: "0 and 0", args: args{a: 0, b: 0}, want: 0},
		{name: "-1 and -2", args: args{a: -1, b: -2}, want: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
