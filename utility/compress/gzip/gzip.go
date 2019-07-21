package gzip

import (
	"bytes"
	gz "compress/gzip"
)

func Deflate(v []byte, level int) ([]byte, error) {
	gzipBuffer := bytes.NewBuffer([]byte{})
	gzipBuffer.Grow(len(v) / 10)
	gzipWriter, _ := gz.NewWriterLevel(gzipBuffer, level)
	defer gzipWriter.Close()
	if _, err := gzipWriter.Write(v); err != nil {
		return nil, err
	}
	gzipWriter.Flush()
	return gzipBuffer.Bytes(), nil
}

func Inflate(v []byte) ([]byte, error) {
	gzipDecoder, err := gz.NewReader(bytes.NewReader(v))
	if err != nil {
		return nil, err
	}
	var decompressed []byte
	if _, err := gzipDecoder.Read(decompressed); err != nil {
		return nil, err
	}
	return decompressed, nil
}
