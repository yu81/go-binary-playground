package character

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func generateKilledMonsters(n int) []*KilledMonsters {
	result := make([]*KilledMonsters, 0, n)
	for i := 0; i < n; i++ {
		count := rand.Int63n(500)
		wanted := false
		if count >= 100 {
			wanted = true
		}
		m := KilledMonsters{
			Id:     time.Now().Unix(),
			Count:  count,
			Wanted: wanted,
		}
		result = append(result, &m)
	}
	return result
}

func testCharacterDataStruct() Character {
	return Character{
		Name:         "yoshihiko",
		Id:           1,
		Strength:     45,
		Vitality:     50,
		Agility:      120,
		Intelligence: 13,
		Dexterity:    20,
		Luck:         75,
		Hp:           134,
		Mp:           4,
		KilledMonsters: []*KilledMonsters{
			{Id: 1, Count: 80, Wanted: false},
			{Id: 4, Count: 180, Wanted: true},
			{Id: 9, Count: 20, Wanted: false},
			{Id: 16, Count: 99, Wanted: false},
			{Id: 25, Count: 370, Wanted: true},
			{Id: 36, Count: 175, Wanted: true},
		},
	}
}

func testCharacterDataStructRandomKilledMonsters(n int) Character {
	c := testCharacterDataStruct()
	c.KilledMonsters = generateKilledMonsters(n)
	return c
}

func TestCharacter_CompressWithProtoBuf(t *testing.T) {
	c := testCharacterDataStructRandomKilledMonsters(1000)
	serialized, err := c.CompressWithProtoBuf(false)
	assert.NoError(t, err)
	assert.True(t, len(serialized) > 100)
	gzipSerialized, err := c.CompressWithProtoBuf(true)
	assert.NoError(t, err)
	assert.True(t, len(gzipSerialized) > 100)
	assert.True(t, len(serialized) > len(gzipSerialized), "gzipped bytes should be smaller")

	const protoVer = "3"
	fileName := "./testdata/protobuf_test_data_protover" + protoVer + "_" + runtime.Version()
	err = ioutil.WriteFile(fileName, serialized, 0666)
	assert.NoError(t, err)
	fmt.Printf("sizeof protobuf binary: %d\n", len(serialized))

	gzipFileName := fileName + ".gz"
	fmt.Printf("sizeof protobuf binary gzip: %d\n", len(gzipSerialized))
	err = ioutil.WriteFile(gzipFileName, gzipSerialized, 0666)
	assert.NoError(t, err)
}

func TestDecompressWithProtoBuf(t *testing.T) {
	c := testCharacterDataStructRandomKilledMonsters(10)
	serialized, _ := proto.Marshal(&c)
	assert.Equal(t, 154, len(serialized))
	var cc Character
	err := proto.Unmarshal(serialized, &cc)
	assert.NoError(t, err)
	serialized2, err := c.CompressWithProtoBuf(true)
	assert.NoError(t, err)
	decompressedCharacter, err := DecompressWithProtoBuf(serialized2, true)
	assert.NoError(t, err)
	assert.Len(t, decompressedCharacter.KilledMonsters, 10)
}

func TestCharacter_CompressWithGob(t *testing.T) {
	c := testCharacterDataStructRandomKilledMonsters(1000)
	gobSerialized, err := c.CompressWithGob(false)
	assert.NoError(t, err)
	fileName := "./testdata/gob_test_data_" + runtime.Version()
	err = ioutil.WriteFile(fileName, gobSerialized, 0666)
	assert.NoError(t, err)
	fmt.Printf("sizeof gob binary: %d\n", len(gobSerialized))

	gobGzSerialized, err := c.CompressWithGob(true)
	fileName = "./testdata/gob_test_data_" + runtime.Version() + ".gz"
	err = ioutil.WriteFile(fileName, gobGzSerialized, 0666)
	assert.NoError(t, err)
	fmt.Printf("sizeof gob binary gzipped: %d\n", len(gobGzSerialized))
}

func TestDecompressWithGob(t *testing.T) {
	const fileNameCommon = "./testdata/gob_test_data_go"
	for _, v := range []string{"1.7.6", "1.11.1", "1.12.6"} {
		buf, err := os.OpenFile(fileNameCommon+v, os.O_RDONLY, 0666)
		data, err := ioutil.ReadAll(buf)
		assert.NoError(t, err)
		c, err := DecompressWithGob(data, false)
		assert.NoError(t, err)
		assert.True(t, c.Id > 0)
		assert.True(t, len(c.KilledMonsters) > 1)
	}
}

type benchmarkCase struct {
	name string
	data Character
	gzip bool
}

func BenchmarkCharacter_proto_Marshal(b *testing.B) {
	c := testCharacterDataStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(&c)
	}
}

func BenchmarkCharacter_proto_Unmarshal(b *testing.B) {
	c := testCharacterDataStruct()
	serialized, _ := proto.Marshal(&c)
	var cc Character
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(serialized, &cc)
	}
}

func BenchmarkCharacter_CompressWithProtoBuf(b *testing.B) {
	c := testCharacterDataStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.CompressWithProtoBuf(false)
	}
}

func BenchmarkDecompressWithProtoBuf(b *testing.B) {
	c := testCharacterDataStruct()
	data, _ := c.CompressWithProtoBuf(false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DecompressWithProtoBuf(data, false)
	}
}

func BenchmarkCharacter_GobEncode(b *testing.B) {
	c := testCharacterDataStruct()
	buf := bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(&c)
	}
}

func BenchmarkCharacter_GobDecode(b *testing.B) {
	c := testCharacterDataStruct()
	buf := bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buf)
	encoder.Encode(&c)
	decoder := gob.NewDecoder(buf)
	cc := Character{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(&cc)
	}
}

func BenchmarkCharacter_CompressWithGob(b *testing.B) {
	benchmarks := []benchmarkCase{
		{"gzip_on", testCharacterDataStruct(), true},
		{"gzip_off", testCharacterDataStruct(), false},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.data.CompressWithGob(bm.gzip)
			}
		})
	}
	b.ResetTimer()
}

func BenchmarkDecompressWithGob(b *testing.B) {
	benchmarks := []benchmarkCase{
		{"gzip_on", testCharacterDataStruct(), true},
		{"gzip_off", testCharacterDataStruct(), false},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			d, _ := bm.data.CompressWithGob(bm.gzip)
			for i := 0; i < b.N; i++ {
				DecompressWithGob(d, bm.gzip)
			}
		})
	}
}
