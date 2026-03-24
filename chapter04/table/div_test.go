package table

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableDivie(t *testing.T) {
	// Arange

	type testCase struct {
		x         int8
		y         int8
		wantError error
		want      string
	}

	tests := map[string]testCase{
		"pos x, pos y":   {x: 8, y: 4, want: "2.00"},
		"neg x, neg y":   {x: -4, y: -8, want: "0.50"},
		"equal x, y":     {x: 4, y: 4, want: "1.00"},
		"max x, pos y":   {x: 127, y: 2, want: "63.50"},
		"min x, pos y":   {x: -128, y: 2, want: "-64.00"},
		"zero x, pos y":  {x: 0, y: 2, want: "0.00"},
		"pos x, zero y":  {x: 10, y: 0, wantError: errors.New("cannot divide by 0")},
		"zero x, zero y": {x: 0, y: 0, wantError: errors.New("cannot divide by 0")},
		"max x, max y":   {x: 127, y: 127, want: "1.00"},
		"min x, min y":   {x: -128, y: -128, want: "1.00"},
	}

	for name, testCase := range tests {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			// Act
			t.Parallel()
			r, err := Divide(tc.x, tc.y)

			// Assert
			if testCase.wantError != nil {
				assert.Equal(t, tc.wantError, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tc.want, *r)
		})
	}
}

func TestDivide(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		x, y := int8(8), int8(4)
		r, err := Divide(x, y)
		assert.Nil(t, err)
		assert.Equal(t, "2.00", *r)
	})
}
