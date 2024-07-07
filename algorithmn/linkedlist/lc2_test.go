package linkedlist

// import (
// 	"os"
// 	"testing"
// )

// type testExample struct {
// 	l1 *Lc2ListNode
// 	l2 *Lc2ListNode
// 	expect *Lc2ListNode
// }

// var (
// 	l1 = [][]int{
// 		{2, 4, 3},
// 		{0},
// 		{9, 9, 9, 9, 9, 9, 9, 9, 9},
// 	}
	
// 	l2 = [][]int{
// 		{5, 6, 4},
// 		{0},
// 		{9, 9, 9, 9},
// 	}

// 	tests map[string]testExample
// )



// func TestAddTwoNumbers(t *testing.T) {
// 	for name, te := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			resList := addTwoNumbers(te.l1, te.l2)
// 			p1 := resList
// 			p2 := te.expect

// 			for p2 != nil {
// 				if p1 == nil {
// 					t.Error("res failed")
// 					return 
// 				}

// 				if p2.Val != p1.Val{
// 					t.Errorf("expected: %v, got: %v\n", p2.Val, p1.Val)
// 					return 
// 				}

// 				p2 = p2.Next
// 				p1 = p1.Next
// 			}
// 		})
// 	}
// }

// func TestMain(m *testing.M) {
	
// 	retTest := m.Run()
// 	os.Exit(retTest)
// }