package main

import (
	"fmt"
	"os"
	"strings"
)

type NamePicker struct {
	Billboards1 map[string]string
	Billboards2 map[string]string
	Billboards3 map[string]string
}

func NewNamePicker() *NamePicker {
	new := NamePicker{}

	new.Billboards1= make(map[string]string)
	new.Billboards2= make(map[string]string)
	new.Billboards3= make(map[string]string)

	new.Billboards1 = map[string]string {
		"A": "Stinky",
		"B": "Lumpy",
		"C": "Buttercup",
		"D": "Gidget",
		"E": "Crusty",
		"F": "Greasy",
		"G": "Fluffy",
		"H": "Cheeseball",
		"I": "Chim chim",
		"J": "Poopsie",
		"K": "Funky",
		"L": "Booger",
		"M": "Pinky",
		"N": "Zippy",
		"O": "Goober",
		"P": "Doofus",
		"Q": "Slimy",
		"R": "Loopy",
		"S": "Snotty",
		"T": "Failure",
		"U": "Dorky",
		"V": "Squeezit",
		"W": "Oprah",
		"X": "Skimmer",
		"Y": "Dinky",
		"Z": "Zsa Zsa",
	}

	new.Billboards2 = map[string]string{
		"A": "Diaper",
		"B": "Toilet",
		"C": "Giggle",
		"D": "Bubble",
		"E": "Girdle",
		"F": "Barf",
		"G": "Lizard",
		"H": "Waffle",
		"I": "Cootie",
		"J": "Monkey",
		"K": "Potty",
		"L": "Liver",
		"M": "Banana",
		"N": "Rhino",
		"O": "Burger",
		"P": "Hamster",
		"Q": "Toad",
		"R": "Gizzard",
		"S": "Pizza",
		"T": "Gerbil",
		"U": "Chicken",
		"V": "Pickle",
		"W": "Chuckle",
		"X": "Tofu",
		"Y": "Gorilla",
		"Z": "Stinker",
	}
	new.Billboards3 = map[string]string {
		"A": "Head",
		"B": "Mouth",
		"C": "Face",
		"D": "Nose",
		"E": "Tush",
		"F": "Breath",
		"G": "Pants",
		"H": "Snorts",
		"I": "Lips",
		"J": "Honker",
		"K": "Butt",
		"L": "Brain",
		"M": "Tushy",
		"N": "Chunks",
		"O": "Hiney",
		"P": "Biscuits",
		"Q": "Toes",
		"R": "Buns",
		"S": "Fanny",
		"T": "Sniffer",
		"U": "Sprinkles",
		"V": "Kisser",
		"W": "Squirt",
		"X": "Humperdinck",
		"Y": "Brains",
		"Z": "Juice",
	}

	return &new
}

func (p *NamePicker) SillyNameFor(realName string) (sillyName string) {
	names := strings.Split(realName, " ")

	sillyFirstName := p.Billboards1[names[0][0:1]]
	lastLetterIdx := len(names[1])
	startSillyLastName := p.Billboards2[names[1][0:1]]
	endSillyLastName := p.Billboards3[strings.ToUpper(names[1][lastLetterIdx-1:lastLetterIdx])]

	return fmt.Sprintf("%s %s%s", sillyFirstName, startSillyLastName, strings.ToLower(endSillyLastName))
}

func main() {
	namePicker := NewNamePicker()
	realName := os.Args[1]
	fmt.Println(namePicker.SillyNameFor(realName))
}
