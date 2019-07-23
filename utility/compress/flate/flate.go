package flate

import (
	"bytes"
	"compress/flate"
	"io"
)

func Deflate(v []byte, level int) ([]byte, error) {
	flateBuffer := bytes.NewBuffer([]byte{})
	flateBuffer.Grow(len(v))
	flateWriter, _ := flate.NewWriter(flateBuffer, level)
	if _, err := flateWriter.Write(v); err != nil {
		return nil, err
	}
	if err := flateWriter.Flush(); err != nil {
		return nil, err
	}
	if err := flateWriter.Close(); err != nil {
		return nil, err
	}
	return flateBuffer.Bytes(), nil
}

func Inflate(v []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(v)
	flateDecoder := flate.NewReader(buffer)

	decompressed := bytes.NewBuffer([]byte{})
	decompressed.Grow(len(v))
	if _, err := io.Copy(decompressed, flateDecoder); err != nil {
		return nil, err
	}
	if err := flateDecoder.Close(); err != nil {
		return nil, err
	}
	return decompressed.Bytes(), nil
}
