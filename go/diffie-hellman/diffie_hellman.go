package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

const testVersion = 1

// PrivateKey generates a random private key.
func PrivateKey(p *big.Int) *big.Int {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := big.NewInt(2)
	limit := big.NewInt(0).Sub(p, min)
	key := big.NewInt(0).Rand(source, limit)
	return key.Add(key, min)
}

// PublicKey generates a public key for a private key.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

// NewPair generates a new key pair.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey generates a secret key.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}
