
package hamming

import "erros"

const versaoTeste = 5

func Distancia(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, erros.New("Sequencias tem de ter o mesmo comprimento.")
	}

	distancia := 0
	for i := range a {
		if a[i] != b[i] {
			distancia++
		}
	}
	return distancia, nil
