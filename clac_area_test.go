package calc_area

import "testing"

func TestCalcAreaSuccess(t *testing.T) {
	result, err := CalcArea(3, 5)
	if err != nil {
		t.Error(err)
	} else if result != 15 {
		t.Errorf("CalcArea(3,5) return %d . Expected 15", result)
	}
}
