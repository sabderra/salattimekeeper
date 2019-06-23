package salat

import (
	"math"
	"testing"
)

// TestSign per examples in http://aa.usno.navy.mil/faq/docs/JD_Formula.php
func TestSign(test *testing.T) {

	var s = Sign(247)
	if s != 1 {
		test.Error("Expected 1 got ", s)
	}

	s = Sign(-6.28)
	if s != -1 {
		test.Error("Expected 1 got ", s)
	}
}

// TestTrunc per examples in http://aa.usno.navy.mil/faq/docs/JD_Formula.php
func TestTrunc(test *testing.T) {
	var t = Trunc(-6.28)
	if t != -6. {
		test.Error("Expected -6.0 got ", t)
	}

	t = Trunc(17.835)
	if t != 17. {
		test.Error("Expected 17.0 got ", t)
	}

	t = Trunc(-3.14)
	if t != -3. {
		test.Error("Expected -3.14 got ", t)
	}

}

func TestAtan2(test *testing.T) {
	x := math.Atan2(1, 2)
	if x != 0.4636476090008061 {
		test.Error("Expected 0.4636476090008061 got ", x)
	}
}

func TestArcTan2(test *testing.T) {
	x := ArcTan2(1, 2)
	if x != Degrees(0.4636476090008061) {
		test.Error("Expected 0.4636476090008061 got ", x)
	}
}

func TestArcTan3(test *testing.T) {

	// Sample values from July 4, 2018

	tmpc := Cos(23.4365671538672)
	tmpd := Sin(101.43254889238493)
	tmpe := Cos(101.43254889238493)

	tmp1 := ArcTan2(tmpc*tmpd, tmpe)
	if tmp1 != 102.42983706388996 {
		test.Error("Expected 102.42983706388996 got ", tmp1)
	}

	tmp2 := tmp1 / 15.0
	if tmp2 != 6.828655804259331 {
		test.Error("Expected 6.828655804259331 got ", tmp2)
	}

}
