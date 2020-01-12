package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte,error){
	return bcrypt.GenerateFromPassword([]byte(password), 12)
}
func ArePasswordsSame(hashedPassword string,rawPassword string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err != bcrypt.ErrMismatchedHashAndPassword
}
