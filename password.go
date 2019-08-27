package password

import "golang.org/x/crypto/bcrypt"

type PasswordHash []byte

type BcryptPassword struct {
  password string
  hash PasswordHash
  cost int
} 

type Password interface {
  hashPassword() (error)
  checkPasswordHash() bool
}

func (h *BcryptPassword) hashPassword() (error) {
  initCost(h)
  bytes, err := bcrypt.GenerateFromPassword([]byte(h.password), h.cost)
  h.hash = PasswordHash(bytes)
  return err
}

func (h *BcryptPassword) checkPasswordHash() bool {
  initCost(h)
  err := bcrypt.CompareHashAndPassword([]byte(h.hash), []byte(h.password))
  return err == nil
}

func initCost(h *BcryptPassword) {
  if h.cost == 0 {
    h.cost = 12
  }
}

func Hash(password Password) (error) {
  return password.hashPassword()
}

func Validate(password Password) bool {
  return password.checkPasswordHash()
}