package iterations

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1041",
			args: args{
				N: 1041,
			},
			want: 5,
		},
		{
			name: "32",
			args: args{
				N: 32,
			},
			want: 0,
		},
		{
			name: "4547",
			args: args{
				N: 4547,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
