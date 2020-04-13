package password

type Password interface {
	Hash(plain string) error
	Check(plain string) bool
}
