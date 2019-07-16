package character

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"encoding/gob"

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

func TestCharacter_Marshal(t *testing.T) {
	c := testCharacterDataStruct()
	fmt.Println(c)
	serialized, err := proto.Marshal(&c)
	assert.NoError(t, err)
	cc := new(Character)
	err = proto.Unmarshal(serialized, cc)
	assert.NoError(t, err)
}

func TestCharacter_XXX_Unmarshal(t *testing.T) {
	c := testCharacterDataStructRandomKilledMonsters(10)
	serialized, _ := proto.Marshal(&c)
	var cc Character
	err := proto.Unmarshal(serialized, &cc)
	assert.NoError(t, err)
}

func BenchmarkCharacter_XXX_Marshal(b *testing.B) {
	c := testCharacterDataStruct()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(&c)
	}
}

func BenchmarkCharacter_XXX_Unmarshal(b *testing.B) {
	c := testCharacterDataStructRandomKilledMonsters(1000)
	serialized, _ := proto.Marshal(&c)
	var cc Character
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(serialized, &cc)
	}
}

func BenchmarkCharacter_GobEncode(b *testing.B) {
	c := testCharacterDataStructRandomKilledMonsters(1000)
	buf := bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(&c)
	}
}

func BenchmarkCharacter_GobDecode(b *testing.B) {
	c := testCharacterDataStructRandomKilledMonsters(1000)
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
