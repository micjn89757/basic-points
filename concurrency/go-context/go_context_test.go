package context

import (
	"context"
	"testing"
)

func TestStep(t *testing.T) {
	base := context.TODO()
	s1 := step1(base)
	s2 := step2(s1)
	step3(s2)
}