package nbs

import (
	"bytes"
)

// readShort reads a short from a buffer.
func readShort(buf *bytes.Buffer) (int16, error) {
	byte1, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	byte2, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	return int16(byte1) + (int16(byte2) << 8), nil
}

// readInt reads an int from a buffer.
func readInt(buf *bytes.Buffer) (int32, error) {
	byte1, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	byte2, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	byte3, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	byte4, err := buf.ReadByte()
	if err != nil {
		return 0, err
	}
	return int32(byte1) + (int32(byte2) << 8) + (int32(byte3) << 16) + (int32(byte4) << 24), nil
}

// readString reads a string from a buffer.
func readString(buf *bytes.Buffer) (string, error) {
	length, err := readInt(buf)
	if err != nil {
		return "", err
	}
	return string(buf.Next(int(length))), nil
}