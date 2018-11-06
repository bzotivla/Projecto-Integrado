package bob

import "strings"

const versaoTeste = 2

func Ola(s string) string {
	s = strings.TrimSpace(s)
	if Silencio(s) {
		return "Fine. Be that way!"
	} else if Gritar(s) {
		return "Whoa, chill out!"
	} else if Pergunta(s) {
		return "Sure."
	}
	return "Whatever."
}

func Silencio(s string) bool {
	return len(s) == 0
}

func Gritar(s string) bool {
	return s == strings.ToUpper(s) && s != strings.ToLower(s)
}

func Pergunta(s string) bool {
	return s[len(s)-1:] == "?"
}
