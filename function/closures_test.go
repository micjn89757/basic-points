package function

import "testing"

func TestClosures(t *testing.T) {
	nextInt := intSeq()
	for range 5 {
		t.Log(nextInt())
	}


	nextInt2 := intSeq()

	for range 5 {
		t.Log(nextInt2())
	}
}