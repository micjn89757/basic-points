package commonutils

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	t.Log(ctx.Deadline())
} 

func TestWithValuePractice(t *testing.T) {
	withValuePractice()
}

func TestWithTimeoutPractice(t *testing.T) {
	withTimeoutPractice()
}

func TestWithDeadlinePractice(t *testing.T) {
	withDeadlinePractice()
}

func TestWithCancelPractice(t *testing.T) {
	withCancelPractice()
}