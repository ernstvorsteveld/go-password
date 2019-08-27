package password

import (
  "testing"
  "fmt"
)

func Test_should_validate_password_with_bcrypt(t *testing.T) {
  got := "MySecret01!"
  bcryptPassword := BcryptPassword{password: got, cost: 10}
  
  error := Hash(&bcryptPassword)
  if error != nil {
    t.Errorf("Got an error while hashing, %s", error)
  }
  
  fmt.Printf("Password: %s\n", bcryptPassword.password)
  fmt.Printf("Hash: %s\n", bcryptPassword.hash)
  fmt.Printf("Cost: %d\n", bcryptPassword.cost)
  
  if bcryptPassword.hash == nil {
    t.Errorf("The hash is still empty, even after hashing is done.")
  }
  
  bcryptPassword.password = got
  validPassword := Validate(&bcryptPassword)
  if !validPassword {
    t.Errorf("The passwords did not match, although it was correct")
  }
  
  wrongPassword := "JustAPassword"
  bcryptPassword.password = wrongPassword
  validPassword = Validate(&bcryptPassword)
  if(validPassword) {
    t.Errorf("The passwords should not match.")
  }
 }