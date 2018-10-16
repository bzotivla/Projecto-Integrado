package clock

import "fmt"

const versaoTeste = 2

type Relogio struct {
	min int
}

func Time(hora, minuto int) Relogio {
	c := Relogio{(60*hora + minuto) % 1440}
	if c.min < 0 {
		c.min += 1440
	}
	return c
}

func (c Relogio) String() string {
	return fmt.Sprintf("%02d:%02d", c.min/60, c.min%60)
}

func (c Relogio) Add(min int) Relogio {
	c.min = (c.min + min) % 1440
	if c.min < 0 {
		c.min += 1440
	}
	return c
}