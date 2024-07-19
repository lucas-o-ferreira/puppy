package puppy

import "github.com/lucas-o-ferreira/dog"

// doggy

func Bark() string {
	return "Bark!"
}

func Barks() string {
	return "Bark! Bark! Bark!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
