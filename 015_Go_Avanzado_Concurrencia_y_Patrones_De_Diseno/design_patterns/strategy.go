package design_patterns

import "fmt"

type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

func NewPasswordProtector(user, passwordName string, algorithm HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: algorithm,
	}
}

func (p *PasswordProtector) SetHashAlgorithm(algorithm HashAlgorithm) {
	p.hashAlgorithm = algorithm
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}

func (SHA) Hash(p *PasswordProtector) {
	fmt.Println("Hashing", p.passwordName, "with SHA")
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProtector) {
	fmt.Println("Hashing", p.passwordName, "with MD5")
}

func StrategyExample() {
	sha := &SHA{}
	md5 := &MD5{}

	pp := NewPasswordProtector("Carlos", "Email", md5)
	pp.Hash()

	pp.SetHashAlgorithm(sha)
	pp.Hash()
}
