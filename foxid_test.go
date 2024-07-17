package foxid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Should create empty FOxID
func TestEmpty(t *testing.T) {
	id := Empty()

	assert.Equal(t, id.Timestamp(), uint64(0))
	assert.Equal(t, id.Datacenter(), uint16(0))
	assert.Equal(t, id.Worker(), uint16(0))
	assert.Equal(t, id.Counter(), uint32(0))
	assert.Equal(t, id.Random(), uint32(0))
}

// Should decode valid FOxID from string
func TestParse(t *testing.T) {
	id, err := Parse("068A2H8PWEBA8G3H00000QTKKM")

	assert.Nil(t, err)
	assert.Equal(t, id.Timestamp(), uint64(1720692578019))
	assert.Equal(t, id.Datacenter(), uint16(38564))
	assert.Equal(t, id.Worker(), uint16(16497))
	assert.Equal(t, id.Counter(), uint32(0))
	assert.Equal(t, id.Random(), uint32(6247325))
}

// Should not decode invalid FOxID from string
func TestInvalidParse(t *testing.T) {
	_, err := Parse("")

	assert.Error(t, err, ErrInvalidFOxID)
}

// Should generate FOxID with predefined variables
func TestGenerate(t *testing.T) {
	id := Generate(Config{
		Datacenter: 10,
		Worker:     20,
		Counter:    30,
		Random:     40,
	})

	assert.Equal(t, id.Datacenter(), uint16(10))
	assert.Equal(t, id.Worker(), uint16(20))
	assert.Equal(t, id.Counter(), uint32(30))
	assert.Equal(t, id.Random(), uint32(40))
}
