package jsgoja

import "testing"

func TestJS(t *testing.T) {
	t.Run("algorithm", func(t *testing.T) {
		tests := []struct {
			name string
			in   int32
			want int32
		}{
			{name: "5", in: 5, want: 5},
			{name: "10", in: 10, want: 55},
			{name: "20", in: 20, want: 6765},
			{name: "30", in: 30, want: 832040},
		}

		f := NewFibonacci()
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		t.Skip("不支持http请求")

		f := NewRequestHTTP()
		f("xxx", "yyy")
	})
}
