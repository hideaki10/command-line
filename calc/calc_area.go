package calc_area

import (
	"errors"
)

func CalcArea(width, height int) (int, error) {
	if width == 0 {
		return 0, errors.New("the width can't be zero")
	}
	if height == 0 {
		return 0, errors.New("the height can't be zero")
	}
	if width < 0 {
		return 0, errors.New("the width can't be negative")
	}

	if height < 0 {
		return 0, errors.New("the height can't be negative")
	}

	return width * height, nil
}

// func TestCalcAreaFail(t *testing.T) {
// 	_, err := CalcArea(-3, 6)
// 	if err == nil {
// 		t.Error("Except CalcArea(-3,6) to return an error")
// 	}
// 	if err.Error() != errorMessage {
// 		t.Errorf("Expected error to be: " + errorMessage)
// 	}
// }

// func TestCalcAreaViaTable(t *testing.T) {
// 	var tests = []struct {
// 		width    int
// 		height   int
// 		expected int
// 	}{
// 		{1, 1, 1},
// 		{5, 6, 30},
// 		{1, 99, 99},
// 		{7, 6, 42},
// 	}

// 	for _, test := range tests {
// 		t.Run("", func(tt *testing.T) {
// 			tt.Parallel()
// 			time.Sleep(time.Second)
// 			w := test.width
// 			h := test.height
// 			r, err := CalcArea(w, h)
// 			if err != nil {
// 				tt.Errorf("CalcArea(%d, %d) returned an error", w, h)
// 			} else if r != test.expected {
// 				tt.Errorf("CalcArea(%d, %d) returned %d Expected %d", w, h, r, test.expected)
// 			}
// 		})

// 	}
// }
