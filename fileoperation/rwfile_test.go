package fileoperation

import (
	"testing"
)

func TestGetCurrentPath(t *testing.T) {
	t.Log(GetCurrentPath())
}

func TestReadFileBase(t *testing.T) {
	ReadFileBase("a.txt")
}

func TestWriteFileBase(t *testing.T) {
	WriteFileBase("a.txt", "demo")
}

func TestReadFileBufio(t *testing.T) {
	ReadFileBufio("a.txt")
}

func TestWriteFileBufio(t *testing.T) {
	WriteFileBufio("a.txt")
}
