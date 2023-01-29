package challenger

import "testing"

func TestHoopCount(t *testing.T) {
	result := HoopCount(5)
	expect := "Great, now move on to tricks"

	if result != expect {
		t.Errorf("expected %s result %s", expect, result)
	}

}
