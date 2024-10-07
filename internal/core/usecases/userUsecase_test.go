package usecases

import "testing"

func Test_lenID(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test lenID",
			args: args{
				ID: "123",
			},
			want: 3,
		},
		{
			name: "Test lenID 10",
			args: args{
				ID: "1234567890",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenID(tt.args.ID); got != tt.want {
				t.Errorf("lenID() = %v, want %v", got, tt.want)
			}
		})
	}
}
