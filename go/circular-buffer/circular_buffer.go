package circular

import (
	"errors"
)

const testVersion = 4

// Errors returned by Buffer.
var (
	ErrBufferEmpty = errors.New("circular: buffer is empty")
	ErrBufferFull  = errors.New("circular: buffer is full")
)

// Buffer implements a circular buffer.
type Buffer struct {
	count      int
	writeIndex int
	readIndex  int
	bytes      []byte
}

// NewBuffer creates a new circular buffer.
func NewBuffer(size int) *Buffer {
	return &Buffer{bytes: make([]byte, size)}
}

// ReadByte returns the next available byte in the buffer.
func (buffer *Buffer) ReadByte() (byte, error) {
	if buffer.count == 0 {
		return 0, ErrBufferEmpty
	}
	c := buffer.bytes[buffer.readIndex]
	buffer.readIndex++
	buffer.readIndex %= len(buffer.bytes)
	buffer.count--
	return c, nil
}

// WriteByte stores a byte in the buffer.
func (buffer *Buffer) WriteByte(c byte) error {
	if buffer.count == len(buffer.bytes) {
		return ErrBufferFull
	}
	buffer.bytes[buffer.writeIndex] = c
	buffer.writeIndex++
	buffer.writeIndex %= len(buffer.bytes)
	buffer.count++
	return nil
}

// Overwrite stores a byte in the buffer, overwriting the next position if the buffer is full.
func (buffer *Buffer) Overwrite(c byte) {
	if buffer.count == len(buffer.bytes) {
		buffer.ReadByte()
	}
	buffer.WriteByte(c)
}

// Reset puts the buffer in an empty state.
func (buffer *Buffer) Reset() {
	buffer.readIndex = buffer.writeIndex
	buffer.count = 0
}
