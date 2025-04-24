package collections

import "testing"


func Test_SliceSlice(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]  // s := a[low:high]
	t.Logf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	t.Logf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}