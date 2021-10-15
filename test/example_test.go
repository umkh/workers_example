package test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestSum ...
func TestSum(t *testing.T) {
	type testCase struct {
		a, b, c int
	}

	cases := []testCase{
		{
			5, 3, 8,
		},
		{
			17, 5, 22,
		},
	}

	for _, c := range cases {
		name := fmt.Sprintf("%d + %d", c.a, c.b)

		t.Run(name, func(t *testing.T) {
			res := Sum(c.a, c.b)
			result := c.c

			require.Equal(t, res, result)
		})
	}
}

// BenchmarkSum ...
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(4, 6)
	}
}

// ExampleSum ...
func ExampleSum() {
	fmt.Println(Sum(4, 6))
}
