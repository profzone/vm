package vm

import (
	"testing"
	"bytes"
)

var m *Memory

const (
	FIRST = 0x01 + iota
	SECOND
	THIRD
	FOUTH
	FIFTH
)

func init() {
	m = NewMemory()
}

func TestResize(t *testing.T) {
	m.Resize(1024)
	if m.Len() != 1024 {
		t.Errorf("expected 1024, got %d", m.Len())
	}
}

func TestSet(t *testing.T) {
	source := []byte{
		FIRST, FIRST, FIRST, FIRST, FIRST, FIRST, FIRST, FIRST,
		FIRST, FIRST, FIRST, FIRST, FIRST, FIRST, FIRST, FIRST,
		FIRST, FIRST, FIRST, FIRST, FIRST, FIRST, FIRST, FIRST,
		FIRST, FIRST, FIRST, FIRST, SECOND, THIRD, FOUTH, FIFTH,
	}
	m.Set(0, uint64(len(source)), source)
	m.Print()
}

func TestGet(t *testing.T) {
	result := m.Get(27, 5)
	if !bytes.Equal([]byte{FIRST, SECOND, THIRD, FOUTH, FIFTH}, result) {
		t.Errorf("expected []byte{FIRST, SECOND, THIRD, FOUTH, FIFTH}, got %v", result)
	}
}

func TestGetPtr(t *testing.T) {
	result := m.GetPtr(27, 5)
	if !bytes.Equal([]byte{FIRST, SECOND, THIRD, FOUTH, FIFTH}, result) {
		t.Errorf("expected []byte{FIRST, SECOND, THIRD, FOUTH, FIFTH}, got %v", result)
	}

	result[0] = FIFTH

	cpy := m.Get(27, 5)
	if !bytes.Equal([]byte{FIFTH, SECOND, THIRD, FOUTH, FIFTH}, result) {
		t.Errorf("expected []byte{FIFTH, SECOND, THIRD, FOUTH, FIFTH}, got %v", cpy)
	}
}
