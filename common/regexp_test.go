package commonutils

import (
	"regexp"
	"testing"
)

func TestFindSubMatch(t *testing.T) {
	t.Log(findSubMatch())
	for i, v := range findSubMatch() {
		t.Logf("%d, %v", i, v)
	}
}

func TestMatchStringBytes(t *testing.T) {
	res, err := matchStringBytes()
	if err != nil {
		t.Logf("%v", err)
	}

	t.Log(res)
}


func TestCreateRegExpObj(t *testing.T) {
	match, err := createRegExpObj()
	if err != nil {
		t.Logf("%v", err)
	}

	t.Log(match)
}

func TestRegExpDemo(t *testing.T) {
	str := "Hello, world!"
	re := regexp.MustCompile(`H.llo`)

	t.Logf("%v",re.MatchString(str)) 
}

func TestStrReplace(t *testing.T) {
	t.Log(strReplace())
}