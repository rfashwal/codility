package timecomplexity

import "testing"

func TestPermMissingElem(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				A: []int{2, 1, 5, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		if got := PermMissingElem(tt.args.A); got != tt.want {
			t.Errorf("%q. PermMissingElem() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
