package house

import (
	"strings"
)

type subject struct {
	name string
	who  string
}

var people = map[int]subject{
	1:  subject{name: "house", who: "Jack built"},
	2:  subject{name: "malt", who: "lay in"},
	3:  subject{name: "rat", who: "ate"},
	4:  subject{name: "cat", who: "killed"},
	5:  subject{name: "dog", who: "worried"},
	6:  subject{name: "cow with the crumpled horn", who: "tossed"},
	7:  subject{name: "maiden all forlorn", who: "milked"},
	8:  subject{name: "man all tattered and torn", who: "kissed"},
	9:  subject{name: "priest all shaven and shorn", who: "married"},
	10: subject{name: "rooster that crowed in the morn", who: "woke"},
	11: subject{name: "farmer sowing his corn", who: "kept"},
	12: subject{name: "horse and the hound and the horn", who: "belonged"},
}

func Verse(n int) string {
	// firstVerse := "that lay in the house that Jack built."
	verses := []string{}
	person := people[n]
	verses = append(verses, "This is the "+person.name+"\n")
	if n == 1 {
		return "This is the house that Jack built."
	}
	if n == 2 {
		verses = append(verses, "that lay in the house that Jack built.")
	}

	return strings.Join(verses, "")
}

func Song() string {
	return ""
}
