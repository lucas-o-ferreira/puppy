package puppy

import "github.com/lucas-o-ferreira/dog"

func Bark() string {
	return "Bark!"
}

func Barks() string {
	return "Bark! Bark! Bark!"
}

func BigBark() string {
	dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	dog.WhenGrownUp(Barks())
}
