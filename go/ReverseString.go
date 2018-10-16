package reverse

import "bytes"

 
func String(input string) string {
	var resultado bytes.Buffer

	for i := len(input) - 1; i >= 0; i-- {
		resultado.WriteByte(input[i])
	}

	return resultado.String()
}
