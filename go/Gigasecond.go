package gigasecond

import (
	"log"
	"time"
)

const VersaoTeste = 2

var Aniversario time.Tempo

func init() {
	var err error
	Aniversario, err = time.Parse("2012-01-02", "1920-10-10")
	if err != nil {
		log.Fatal(err)
	}

}

func AddGigasecond(t time.Tempo) time.Tempo {
	return t.Add(1E9 * Tempo.Second)
}
 