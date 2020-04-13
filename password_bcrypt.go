package password

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptPassword struct {
	hash []byte
	Cost int
}

func (h *BcryptPassword) Hash(password string) error {
	initCost(h)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.Cost)
	h.hash = bytes
	return err
}

func (h *BcryptPassword) Check(password string) bool {
	initCost(h)
	err := bcrypt.CompareHashAndPassword([]byte(h.hash), []byte(password))
	return err == nil
}

func initCost(h *BcryptPassword) {
	if h.Cost == 0 {
		h.Cost = 12
	}
}
