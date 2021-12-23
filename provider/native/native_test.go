package native

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		in   int32
		want int32
	}{
		{name: "5", in: 5, want: 5},
		{name: "10", in: 10, want: 55},
		{name: "20", in: 20, want: 6765},
		{name: "30", in: 30, want: 832040},
		{name: "40", in: 40, want: 102334155},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.in); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestHTTP(t *testing.T) {
	tests := []struct {
		name string
		want int32
	}{
		{
			name: "default",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequestHTTP(); got != tt.want {
				t.Errorf("RequestHTTP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileIO(t *testing.T) {
	tests := []struct {
		name string
		want int32
	}{
		{
			name: "default",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileIO(); got != tt.want {
				t.Errorf("FileIO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiThreads(t *testing.T) {
	type args struct {
		num int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "default",
			args: args{num: 4},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiThreads(tt.args.num); got != tt.want {
				t.Errorf("MultiThreads() = %v, want %v", got, tt.want)
			}
		})
	}
}
