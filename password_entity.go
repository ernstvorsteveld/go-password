package password

type Algorithm int

const (
	BCRYPT Algorithm = iota
	C1
	C2
)

type PasswordEntity struct {
	password map[Algorithm][]byte
}
