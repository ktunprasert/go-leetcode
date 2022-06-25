package pow

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"self", args{2.0, 1}, 2.0},
		{"self", args{2.0, 2}, 4.0},
		{"self", args{-2.0, 2}, 4.0},
		{"e1", args{2.0, 10}, 1024.0},
		// {"e2", args{2.1, 3}, 9.261},
		{"e3", args{2.0, -2}, 0.25},
		// {"gotcha1", args{1.0, -2147483648}, 1},
		// {"gotcha2", args{2.0, -2147483648}, 0},
		// {"gotcha3", args{2.0, 2147483648}, 10000},
		// {"gotcha4", args{3.76050, -8}, 3e-05},
		{"gotcha5", args{-1.0, 2147483648}, 1.0},
		{"gotcha6", args{-1.0, 2147483647}, -1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
