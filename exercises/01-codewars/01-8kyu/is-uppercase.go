package kata

import "strings"

type MyString string

func (s MyString) IsUpperCase() bool {
	isUpper := true

	for _, letter := range s {
		if string(letter) != strings.ToUpper(string(letter)) {
			return !isUpper
		}
	}

	return isUpper
}
