package aa

import (
	"math"
	"testing"
)

func Test_eq(t *testing.T) {
	if !eq(math.NaN(), math.NaN()) {
		t.Error()
	}
	if !eq(math.Copysign(0, +1), math.Copysign(0, -1)) {
		t.Error()
	}
}
