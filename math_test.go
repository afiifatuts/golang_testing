package golangtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsolute(t *testing.T) {
	//sub test
	t.Run("negative test case", func(t *testing.T) {
		if testing.Short() {
			t.Skip()
		}

		// <-time.After(5 * time.Second)

		c := Absolute(-5)
		assert.Equal(t, 5, c, "expect 5 , you got %d", c)
	})
	t.Run("positif test case", func(t *testing.T) {
		c := Absolute(5)
		assert.Equal(t, 5, c)
	})

}

func TestAdd_Table(t *testing.T) {
	//table test

	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "negative and negatif",
			a:        -1,
			b:        -1,
			expected: -2,
		},
		{
			name:     "negative and positif",
			a:        -1,
			b:        1,
			expected: 0,
		},
		{
			name:     "positif and positif",
			a:        1,
			b:        1,
			expected: 2,
		},
		{
			name:     "positif and negatif",
			a:        1,
			b:        -1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, c)
		})

	}

}

func TestAdd(t *testing.T) {
	//sub test
	t.Run("negative test case", func(t *testing.T) {
		c := Add(-5, -1)
		assert.Equal(t, -6, c, "expect 6 , you got %d", c)
	})
	t.Run("positif test case", func(t *testing.T) {
		c := Add(5, 1)
		assert.Equal(t, 6, c)
	})

}
