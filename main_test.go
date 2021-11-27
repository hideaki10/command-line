package main

import (
	"errors"
	"testing"
	"time"
)

var errorMessage = "width and height must be positive"

func CalcArea(width, height int) (int, error) {
	if width < 0 || height < 0 {
		return 0, errors.New(errorMessage)
	}

	return width * height, nil
}

func TestCalcAreaSuccess(t *testing.T) {
	result, err := CalcArea(3, 5)
	if err != nil {
		t.Error(err)
	} else if result != 15 {
		t.Errorf("CalcArea(3,5) return %d . Expected 15", result)
	}
}

func TestCalcAreaFail(t *testing.T) {
	_, err := CalcArea(-3, 6)
	if err == nil {
		t.Error("Except CalcArea(-3,6) to return an error")
	}
	if err.Error() != errorMessage {
		t.Errorf("Expected error to be: " + errorMessage)
	}
}

func TestCalcAreaViaTable(t *testing.T) {
	var tests = []struct {
		width    int
		height   int
		expected int
	}{
		{1, 1, 1},
		{5, 6, 30},
		{1, 99, 99},
		{7, 6, 42},
	}

	for _, test := range tests {
		t.Run("", func(tt *testing.T) {
			tt.Parallel()
			time.Sleep(time.Second)
			w := test.width
			h := test.height
			r, err := CalcArea(w, h)
			if err != nil {
				tt.Errorf("CalcArea(%d, %d) returned an error", w, h)
			} else if r != test.expected {
				tt.Errorf("CalcArea(%d, %d) returned %d Expected %d", w, h, r, test.expected)
			}
		})

	}
}
