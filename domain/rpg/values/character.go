package character

import (
	"bytes"
	"encoding/gob"

	"github.com/golang/protobuf/proto"

	libgzip "github.com/yu81/go-binary-playground/utility/compress/gzip"
)

func (m *Character) AddStrength(v int) *Character {
	m.Strength = m.GetStrength() + int64(v)
	return m
}

func (m *Character) AddAgility(v int) *Character {
	m.Agility = m.GetAgility() + int64(v)
	return m
}

func (m *Character) AddDexterity(v int) *Character {
	m.Dexterity = m.GetDexterity() + int64(v)
	return m
}

func (m *Character) AddIntelligence(v int) *Character {
	m.Intelligence = m.GetIntelligence() + int64(v)
	return m
}

func (m *Character) AddLuck(v int) *Character {
	m.Luck = m.GetLuck() + int64(v)
	return m
}

func (m *Character) AddVitality(v int) *Character {
	m.Vitality = m.GetVitality() + int64(v)
	return m
}

func (m *Character) AddHp(v int) *Character {
	m.Hp = m.GetHp() + int64(v)
	return m
}

func (m *Character) AddMp(v int) *Character {
	m.Mp = m.GetMp() + int64(v)
	return m
}

func (m *Character) CompressWithGob(gzip bool) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	buf.Grow(128 + 32*len(m.KilledMonsters))
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(*m); err != nil {
		return nil, err
	}
	if !gzip {
		return buf.Bytes(), nil
	}
	return libgzip.Deflate(buf.Bytes(), 7)
}

func DecompressWithGob(v []byte, gzip bool) (*Character, error) {
	if !gzip {
		cc := Character{}
		if err := gob.NewDecoder(bytes.NewReader(v)).Decode(&cc); err != nil {
			return nil, err
		}
		return &cc, nil
	}
	decompressed, err := libgzip.Inflate(v)
	if err != nil {
		return nil, err
	}
	cc := Character{}
	if err := gob.NewDecoder(bytes.NewReader(decompressed)).Decode(&cc); err != nil {
		return nil, err
	}
	return &cc, nil
}

func (m *Character) CompressWithProtoBuf(gzip bool) ([]byte, error) {
	encoded, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}
	if !gzip {
		return encoded, nil
	}
	return libgzip.Deflate(encoded, 7)
}

func DecompressWithProtoBuf(v []byte, gzip bool) (*Character, error) {
	if !gzip {
		cc := Character{}
		if err := proto.Unmarshal(v, &cc); err != nil {
			return nil, err
		}
		return &cc, nil
	}
	inflated, err := libgzip.Inflate(v)
	if err != nil {
		return nil, err
	}
	cc := Character{}
	if err := proto.Unmarshal(inflated, &cc); err != nil {
		return nil, err
	}
	return &cc, nil
}
