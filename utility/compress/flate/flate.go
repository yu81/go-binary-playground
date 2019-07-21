package flate

import (
	"bytes"
	"compress/flate"
)

func Deflate(v []byte, level int) ([]byte, error) {
	flateBuffer := bytes.NewBuffer([]byte{})
	flateBuffer.Grow(len(v) / 10)
	flateWriter, _ := flate.NewWriter(flateBuffer, level)
	defer flateWriter.Close()
	if _, err := flateWriter.Write(v); err != nil {
		return nil, err
	}
	flateWriter.Flush()
	return flateBuffer.Bytes(), nil
}

func Inflate(v []byte) ([]byte, error) {
	flateDecoder := flate.NewReader(bytes.NewReader(v))
	var decompressed []byte
	if _, err := flateDecoder.Read(decompressed); err != nil {
		return nil, err
	}
	return decompressed, nil
}
