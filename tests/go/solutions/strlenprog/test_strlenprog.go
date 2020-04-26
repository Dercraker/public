package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
)

func main() {
	randomStringCharset := "a b c d e f g h ijklmnopqrstuvwxyz A B C D E FGHIJKLMNOPRSTUVWXYZ"

	table := []string{}
	for i := 0; i < 10; i++ {
		randomLenghtOfWord := z01.RandIntBetween(1, 20)
		randomStrRandomLenght := z01.RandStr(randomLenghtOfWord, randomStringCharset)
		table = append(table, randomStrRandomLenght)
	}
	table = append(table, "Héllo!")
	table = append(table, randomStringCharset)

	for _, s := range table {
		z01.Challenge("StrLenProg", StrLen, correct.StrLen, s)
	}
}