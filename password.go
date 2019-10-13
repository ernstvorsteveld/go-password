package password

type PasswordHash []byte

type Password interface {
	hashPassword() error
	checkPasswordHash() bool
}
