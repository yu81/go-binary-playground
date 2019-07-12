package service

import character "github.com/yu81/go-binary-playground/domain/rpg/values"

func AddStrength(c character.Character, v int) character.Character {
	c.Strength = c.GetStrength() + int64(v)
	return c
}

func AddAgility(c character.Character, v int) character.Character {
	c.Agility = c.GetAgility() + int64(v)
	return c
}

func AddDexterity(c character.Character, v int) character.Character {
	c.Dexterity = c.GetDexterity() + int64(v)
	return c
}

func AddIntelligence(c character.Character, v int) character.Character {
	c.Intelligence = c.GetIntelligence() + int64(v)
	return c
}

func AddLuck(c character.Character, v int) character.Character {
	c.Luck = c.GetLuck() + int64(v)
	return c
}

func AddVitality(c character.Character, v int) character.Character {
	c.Vitality = c.GetVitality() + int64(v)
	return c
}
