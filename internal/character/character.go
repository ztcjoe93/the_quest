package character

import "fmt"

type Character struct {
	name              string
	level             int
	speed             int
	def               int
	health, maxHealth int
	exp               int
}

func CreateCharacter(name string) *Character {
	return &Character{
		name:      name,
		level:     1,
		speed:     3,
		def:       2,
		exp:       0,
		health:    40,
		maxHealth: 40,
	}
}

func (from *Character) Attack(target *Character) {
	fmt.Printf("%s attacked %s and dealt x damage!\n", from.name, target.name)
	target.health -= 3
	fmt.Printf("%s has %d/%d\n", target.name, target.health, target.maxHealth)
}
