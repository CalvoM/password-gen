package passwordgen

import (
	"math/rand"
	"time"
)

//Password : struct
type Password struct {
	UseSymbols bool
	Length     int
}

//Generate : Generate the needed password
func (p *Password) Generate() string {
	rand.Seed(time.Now().UnixNano())
	ucLetter := rand.Intn(25) + 65
	lwLetter := rand.Intn(25) + 97
	symbols := rand.Intn(31) + 33
	password := ""
	det := rand.Intn(50)
	passwordHolder := make([]byte, p.Length)

	for i := 0; i < p.Length; i++ {
		if det%3 == 0 && p.UseSymbols {
			symbols = rand.Intn(31) + 33
			passwordHolder[i] = byte(symbols)
		} else {
			if det%2 == 0 {
				ucLetter = rand.Intn(25) + 65
				passwordHolder[i] = byte(ucLetter)
			} else {
				lwLetter = rand.Intn(25) + 97
				passwordHolder[i] = byte(lwLetter)
			}
		}
		det = rand.Intn(1000)
	}
	password = string(passwordHolder)
	return password
}
