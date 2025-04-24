package commonutils

import (
	"testing"
)

func TestGetValuePractice(t *testing.T) {
	ret, _ := GetValuesPractice()
	t.Log(ret)
}

func TestGetPractice(t *testing.T) {
	ret, _ := GetPractice()
	t.Log(ret)
}

func TestPostPractice(t *testing.T) {
	ret, _ := PostPractice()
	t.Log(ret)
}