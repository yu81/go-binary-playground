package character

import (
	"bytes"
	gz "compress/gzip"
	"encoding/gob"

	"github.com/golang/protobuf/proto"
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
	gzipBuffer := bytes.NewBuffer([]byte{})
	gzipBuffer.Grow(buf.Len() / 10)
	gzipWriter, _ := gz.NewWriterLevel(gzipBuffer, 7)
	defer gzipWriter.Close()
	if _, err := gzipWriter.Write(buf.Bytes()); err != nil {
		return nil, err
	}
	gzipWriter.Flush()
	return gzipBuffer.Bytes(), nil
}

func DecompressWithGob(v []byte, gzip bool) (*Character, error) {
	if !gzip {
		cc := Character{}
		if err := gob.NewDecoder(bytes.NewReader(v)).Decode(&cc); err != nil {
			return nil, err
		}
		return &cc, nil
	}
	gzDecoder, err := gz.NewReader(bytes.NewReader(v))
	if err != nil {
		return nil, err
	}
	var decompressed []byte
	if _, err := gzDecoder.Read(decompressed); err != nil {
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
	gzipBuffer := bytes.NewBuffer([]byte{})
	gzipBuffer.Grow(len(encoded) / 10)
	gzipWriter := gz.NewWriter(gzipBuffer)
	if _, err := gzipWriter.Write(encoded); err != nil {
		return nil, err
	}
	defer gzipWriter.Close()
	if err := gzipWriter.Flush();err != nil {
		return nil, err
	}
	return gzipBuffer.Bytes(), nil
}

func DecompressWithProtoBuf(v []byte, gzip bool) (*Character, error) {
	if !gzip {
		cc := Character{}
		if err := proto.Unmarshal(v, &cc); err != nil {
			return nil, err
		}
		return &cc, nil
	}
	gzDecoder, err := gz.NewReader(bytes.NewReader(v))
	if err != nil {
		return nil, err
	}
	defer gzDecoder.Close()
	var decompressed []byte
	if _, err := gzDecoder.Read(decompressed); err != nil {
		return nil, err
	}
	cc := Character{}
	if err := proto.Unmarshal(v, &cc); err != nil {
		return nil, err
	}
	return &cc, nil
}
