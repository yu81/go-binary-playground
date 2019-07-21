package flate

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testByteSlice(l int) []byte {
	result := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		result = append(result, byte(i%256))
	}
	return result
}

func TestDeflateAndInflate(t *testing.T) {
	testCases := [][]byte{
		{0, 255, 254, 222, 143, 234, 200},
	}
	compressed, err := Deflate(testCases[0], 7)
	assert.NoError(t, err)

	decompressed, err := Inflate(compressed)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(testCases[0], decompressed))
}

func BenchmarkDeflate(b *testing.B) {
	data := testByteSlice(102400)
	benchmarks := []struct {
		name  string
		level int
	}{
		{"deflate level 0", 0},
		{"deflate level 1", 1},
		{"deflate level 2", 2},
		{"deflate level 3", 3},
		{"deflate level 4", 4},
		{"deflate level 5", 5},
		{"deflate level 6", 6},
		{"deflate level 7", 7},
		{"deflate level 8", 8},
		{"deflate level 9", 9},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Deflate(data, 7)
			}
		})
	}
}

func BenchmarkInflate(b *testing.B) {
	rawData := testByteSlice(102400)
	dataList := make([][]byte, 0, 10)
	for i := 0; i < 10; i++ {
		d, _ := Deflate(rawData, i)
		dataList = append(dataList, d)
	}
	benchmarks := []struct {
		name string
		data []byte
	}{
		{"deflate level 0", dataList[0]},
		{"deflate level 1", dataList[1]},
		{"deflate level 2", dataList[2]},
		{"deflate level 3", dataList[3]},
		{"deflate level 4", dataList[4]},
		{"deflate level 5", dataList[5]},
		{"deflate level 6", dataList[6]},
		{"deflate level 7", dataList[7]},
		{"deflate level 8", dataList[8]},
		{"deflate level 9", dataList[9]},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Inflate(bm.data)
			}
		})
	}
}
