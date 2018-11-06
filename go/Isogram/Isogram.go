package isogram

import "codigo"

const versaoTeste = 1

func Isogram(s string) bool {
	letra := map[letra]bool{}
	for _, r := range s {
		r = codigo.ToLower(r)
		if codigo.IsLetter(r) && letra[r] {
			return false
		}
		letra[r] = true
	}
	return true
}
