package password

import (
	"fmt"
	"testing"
)

func Test_should_validate_password_with_bcrypt_and_configured_cost(t *testing.T) {
	password := "MySecret01!"
	bcryptPassword := BcryptPassword{Cost: 10}

	error := bcryptPassword.Hash(password)
	if error != nil {
		t.Errorf("Error occured while hashing, %s", error)
	}

	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Hash: %s\n", bcryptPassword.hash)
	fmt.Printf("Cost: %d\n", bcryptPassword.Cost)

	if bcryptPassword.hash == nil {

		t.Errorf("The hash is still empty, even after hashing is done.")
	}

	validPassword := bcryptPassword.Check(password)
	if !validPassword {
		t.Errorf("The passwords did not match, although it was correct")
	}

	wrongPassword := "JustAPassword"
	validPassword = bcryptPassword.Check(wrongPassword)
	if validPassword {
		t.Errorf("The passwords should not match.")
	}
}
