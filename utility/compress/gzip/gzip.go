package gzip

import (
	"bytes"
	gz "compress/gzip"
	"io"
)

func Deflate(v []byte, level int) ([]byte, error) {
	gzipBuffer := bytes.NewBuffer([]byte{})
	gzipBuffer.Grow(len(v))
	gzipWriter, _ := gz.NewWriterLevel(gzipBuffer, level)
	if _, err := gzipWriter.Write(v); err != nil {
		return nil, err
	}
	if err := gzipWriter.Flush(); err != nil {
		return nil, err
	}
	if err := gzipWriter.Close(); err != nil {
		return nil, err
	}
	return gzipBuffer.Bytes(), nil
}

func Inflate(v []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(v)
	gzipDecoder, err := gz.NewReader(buffer)
	if err != nil {
		return nil, err
	}
	decompressed := bytes.NewBuffer([]byte{})
	decompressed.Grow(len(v))
	if _, err := io.Copy(decompressed, gzipDecoder); err != nil {
		return nil, err
	}
	if err := gzipDecoder.Close(); err != nil {
		return nil, err
	}
	return decompressed.Bytes(), nil
}
