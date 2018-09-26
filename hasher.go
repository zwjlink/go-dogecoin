package dogecoin

import "crypto/sha256"

type PublicKey []byte

func (data PublicKey) Hash() []byte {
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	return hash2[:]
}
