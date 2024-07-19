package puppy

import "github.com/lucas-o-ferreira/dog"

func Bark() string {
	return "Bark!"
}

func Barks() string {
	return "Bark! Bark! Bark!"
}

func BigBark() {
	dog.WhenGrownUp(Bark())
}

func BigBarks() {
	dog.WhenGrownUp(Barks())
}
