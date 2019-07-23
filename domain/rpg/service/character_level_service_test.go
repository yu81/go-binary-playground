package service

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	character "github.com/yu81/go-binary-playground/domain/rpg/values"

	"github.com/stretchr/testify/assert"
)

func TestCharacterLevelService_LevelUp(t *testing.T) {
	type fields struct {
		mu sync.RWMutex
	}
	type args struct {
		character *character.Character
	}
	testNoviceCharacter := character.Character{
		Name:           "onion",
		Hp:             14,
		Mp:             0,
		Strength:       6,
		Vitality:       5,
		Agility:        3,
		Dexterity:      3,
		Intelligence:   3,
		Luck:           3,
		Id:             1,
		KilledMonsters: []*character.KilledMonsters{},
	}
	origin := testNoviceCharacter
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *character.Character
	}{
		{name: "novice", args: args{&testNoviceCharacter}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CharacterLevelService{mu: tt.fields.mu}
			if got := c.LevelUp(tt.args.character); !reflect.DeepEqual(got, tt.want) {
				assert.True(t, got.Hp >= origin.Hp, fmt.Sprintf("CharacterLevelService.LevelUp() = %v, original is %v", got.Hp, origin.Hp))
				assert.True(t, got.Mp >= origin.Mp, fmt.Sprintf("CharacterLevelService.LevelUp() = %v, original is %v", got.Mp, origin.Mp))
				assert.True(t, got.Strength >= origin.Strength, fmt.Sprintf("CharacterLevelService.LevelUp() = %v, original is %v", got.Strength, origin.Strength))
				assert.True(t, got.Vitality >= origin.Vitality, fmt.Sprintf("CharacterLevelService.LevelUp() = %v, original is %v", got.Vitality, origin.Vitality))
				assert.True(t, got.Agility >= origin.Agility, fmt.Sprintf("CharacterLevelService.LevelUp() = %v, original is %v", got.Agility, origin.Agility))
				assert.True(t, got.Dexterity >= origin.Dexterity, fmt.Sprintf("CharacterLevelService.LevelUp() = %v, original is %v", got.Dexterity, origin.Dexterity))
				assert.True(t, got.Intelligence >= origin.Intelligence, fmt.Sprintf("CharacterLevelService.LevelUp() = %v, original is %v", got.Intelligence, origin.Intelligence))
				assert.True(t, got.Luck >= origin.Luck, fmt.Sprintf("CharacterLevelService.LevelUp() = %v, original is %v", got.Luck, origin.Luck))
			}
		})
	}
}

func BenchmarkCharacterLevelService_LevelUp(b *testing.B) {
	testNoviceCharacter := character.Character{
		Name:           "onion",
		Hp:             14,
		Mp:             0,
		Strength:       6,
		Vitality:       5,
		Agility:        3,
		Dexterity:      3,
		Intelligence:   3,
		Luck:           3,
		Id:             1,
		KilledMonsters: []*character.KilledMonsters{},
	}
	s := CharacterLevelService{sync.RWMutex{}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.LevelUp(&testNoviceCharacter)
	}
	fmt.Println(testNoviceCharacter)
}
