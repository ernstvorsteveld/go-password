package password

type Password interface {
	hashPassword() error
	checkPasswordHash() bool
}
