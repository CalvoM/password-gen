package passwordgen

//Password : struct
type Password struct {
	useSymbols bool
	length     int
}

//Generate : Generate the needed password
func (p *Password) Generate() string {
	password := ""
	passwordHolder := make([]byte, p.length)

	return password
}
