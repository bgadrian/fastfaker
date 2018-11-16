/*
Package subtext is part of the FastFaker package but it can be used as a standalone library.
It generates random strings and text like words, phrases and passwords.
*/
package subtext

import "github.com/bgadrian/fastfaker/randomizer"

type FakerText struct {
	textRand randomizer.Randomizer
}

func New(randomizer randomizer.Randomizer) *FakerText {
	text := &FakerText{}
	text.textRand = randomizer
	return text
}
