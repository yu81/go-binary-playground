package service

import (
	"math/rand"
	"sync"

	character "github.com/yu81/go-binary-playground/domain/rpg/values"
)

type CharacterLevelService struct {
	mu sync.RWMutex
}

func (c *CharacterLevelService) LevelUp(character *character.Character) *character.Character {
	c.mu.Lock()
	defer c.mu.Unlock()
	character.AddStrength(rand.Intn(10))
	vitality := rand.Intn(10)
	character.AddVitality(vitality)
	character.AddAgility(rand.Intn(10))
	character.AddDexterity(rand.Intn(10))
	intelligence := rand.Intn(10)
	character.AddIntelligence(intelligence)
	character.AddLuck(rand.Intn(10))
	character.AddHp(int(float64(vitality) * 1.2))
	character.AddMp(int(float64(intelligence) * 1.2))

	return character
}
