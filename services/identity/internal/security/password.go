package security

import "golang.org/x/crypto/bcrypt"

const defaultCost = bcrypt.DefaultCost

type PasswordService struct {
	cost int
}

func NewPasswordService() *PasswordService {
	return &PasswordService{cost: defaultCost}
}

func (s *PasswordService) Hash(plain string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), s.cost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (s *PasswordService) Compare(hash, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}
