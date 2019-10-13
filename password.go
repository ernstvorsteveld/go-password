package password

type Password interface {
	Hash(password string) error
	Check(password string) bool
}
