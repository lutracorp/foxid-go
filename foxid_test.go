package foxid

import (
	"errors"
	"testing"
)

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func equalErr(t *testing.T, want, got error) {
	if !errors.Is(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func equal[T uint16 | uint32 | uint64](t *testing.T, want, got T) {
	if want != got {
		t.Errorf("want %x, got %x", want, got)
	}
}

// Should create empty FOxID
func TestEmpty(t *testing.T) {
	id := Empty()

	equal(t, id.Timestamp(), 0)
	equal(t, id.Datacenter(), 0)
	equal(t, id.Worker(), 0)
	equal(t, id.Counter(), 0)
	equal(t, id.Random(), 0)
}

// Should decode valid FOxID from string
func TestParse(t *testing.T) {
	id, err := Parse("068A2H8PWEBA8G3H00000QTKKM")

	checkErr(t, err)
	equal(t, id.Timestamp(), 1720692578019)
	equal(t, id.Datacenter(), 38564)
	equal(t, id.Worker(), 16497)
	equal(t, id.Counter(), 0)
	equal(t, id.Random(), 6247325)
}

// Should not decode invalid FOxID from string
func TestInvalidParse(t *testing.T) {
	_, err := Parse("")

	equalErr(t, err, ErrInvalidFOxID)
}

// Should generate FOxID with predefined variables
func TestGenerate(t *testing.T) {
	id := Generate(Config{
		Datacenter: 10,
		Worker:     20,
		Counter:    30,
		Random:     40,
	})

	equal(t, id.Datacenter(), 10)
	equal(t, id.Worker(), 20)
	equal(t, id.Counter(), 30)
	equal(t, id.Random(), 40)
}
