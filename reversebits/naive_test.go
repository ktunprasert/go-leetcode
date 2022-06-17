package reversebits

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// {"test1", args{0b10}, 1},
		{"e1", args{0b00000010100101000001111010011100}, 964176192},
		{"e2", args{0b11111111111111111111111111111101}, 3221225471},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("\n got %32b\nwant %32b", got, tt.want)
			}
		})
	}
}
