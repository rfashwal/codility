package timecomplexity

import "testing"

func TestFrogJump(t *testing.T) {
	type args struct {
		X int
		Y int
		D int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				X: 10,
				Y: 85,
				D: 30,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FrogJump(tt.args.X, tt.args.Y, tt.args.D); got != tt.want {
				t.Errorf("FrogJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
