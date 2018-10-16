package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Abreviado(in string) string {
	palavras := strings.FieldsFunc(in, Delimitador)
	var out string
	for _, palavra := range palavras {
		out += string(palavra[0])
	}

	return strings.ToUpper(out)
}

func Delimiter(c rune) bool {
	return unicode.IsSpace(c) || c == '-'
}
